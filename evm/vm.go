/*
 * Copyright (c) 2017-2020 The qitmeer developers
 */

package evm

import (
	"encoding/hex"
	"fmt"
	"github.com/Qitmeer/meerevm/chain"
	"github.com/Qitmeer/meerevm/evm/util"
	"github.com/Qitmeer/qng-core/common/hash"
	"github.com/Qitmeer/qng-core/consensus"
	"github.com/Qitmeer/qng-core/core/address"
	"github.com/Qitmeer/qng-core/core/blockchain/opreturn"
	qtypes "github.com/Qitmeer/qng-core/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"runtime"
	"sync"
)

// meerevm ID of the platform
const (
	MeerEVMID = "meerevm"

	txSlotSize = 32 * 1024
	txMaxSize = 4 * txSlotSize
)

type VM struct {
	ctx          consensus.Context
	shutdownChan chan struct{}
	shutdownWg   sync.WaitGroup

	chain  *chain.ETHChain

	txsCh  chan core.NewTxsEvent
	txsSub event.Subscription
}

func (vm *VM) GetID() string {
	return MeerEVMID
}

func (vm *VM) Initialize(ctx consensus.Context) error {
	util.InitLog(ctx.GetConfig().DebugLevel,ctx.GetConfig().DebugPrintOrigins)

	log.Info("System info", "ETH VM Version", util.Version, "Go version", runtime.Version())

	log.Info(fmt.Sprintf("Initialize:%s", ctx.GetConfig().DataDir))

	vm.ctx = ctx

	//

	ethchain,err:=chain.NewETHChain(vm.ctx.GetConfig().DataDir)
	if err != nil {
		return err
	}
	vm.chain = ethchain

	vm.txsSub = ethchain.Ether().TxPool().SubscribeNewTxsEvent(vm.txsCh)

	vm.shutdownWg.Add(1)
	go vm.handler()

	return nil
}

func (vm *VM) Bootstrapping() error {
	log.Debug("Bootstrapping")
	err:=vm.chain.Start()
	if err != nil {
		return err
	}
	//
	rpcClient, err := vm.chain.Node().Attach()
	if err != nil {
		log.Error(fmt.Sprintf("Failed to attach to self: %v", err))
	}
	client := ethclient.NewClient(rpcClient)

	blockNum, err := client.BlockNumber(vm.ctx)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info(fmt.Sprintf("MeerETH block chain current block number:%d", blockNum))
	}

	cbh:=vm.chain.Ether().BlockChain().CurrentBlock().Header()
	if cbh != nil {
		log.Info(fmt.Sprintf("MeerETH block chain current block:number=%d hash=%s",cbh.Number.Uint64(),cbh.Hash().String()))
	}

	//
	state, err := vm.chain.Ether().BlockChain().State()
	if err != nil {
		return nil
	}

	log.Info(fmt.Sprintf("miner account,addr:%v balance:%v", vm.chain.Config().Eth.Miner.Etherbase, state.GetBalance(vm.chain.Config().Eth.Miner.Etherbase)))

	//
	pending, err := vm.chain.Ether().TxPool().Pending(true)
	if err != nil {
		log.Error("Failed to fetch pending transactions", "err", err)
	}else{
		if len(pending) > 0 {
			for _,txs :=range pending {
				for _,tx:=range txs {
					err:=vm.addTx(tx)
					if err != nil {
						log.Error(err.Error())
					}
				}
			}
		}
	}
	
	return nil
}

func (vm *VM) Bootstrapped() error {
	log.Debug("Bootstrapped")
	return nil
}

func (vm *VM) Shutdown() error {
	log.Debug("Shutdown")
	if vm.ctx == nil {
		return nil
	}

	close(vm.shutdownChan)
	vm.chain.Stop()

	vm.chain.Wait()
	vm.shutdownWg.Wait()
	return nil
}

func (vm *VM) Version() string {
	return util.Version + " " + vm.chain.Config().Node.Version
}

func (vm *VM) GetBlock(bh *hash.Hash) (consensus.Block, error) {
	block := vm.chain.Ether().BlockChain().CurrentBlock()
	h := hash.MustBytesToHash(block.Hash().Bytes())
	return &Block{id: &h, ethBlock: block, vm: vm, status: consensus.Accepted}, nil
}

func (vm *VM) BuildBlock(txs []consensus.Tx) (consensus.Block, error) {
	blocks, _ := core.GenerateChain(vm.chain.Config().Eth.Genesis.Config, vm.chain.Ether().BlockChain().CurrentBlock(), vm.chain.Ether().Engine(), vm.chain.Ether().ChainDb(), 1, func(i int, block *core.BlockGen) {

		for _, tx := range txs {
			if tx.GetTxType() == qtypes.TxTypeCrossChainExport {
				pubkBytes, err := hex.DecodeString(tx.GetTo())
				if err != nil {
					log.Warn(err.Error())
					continue
				}
				publicKey, err := crypto.UnmarshalPubkey(pubkBytes)
				if err != nil {
					log.Warn(err.Error())
					continue
				}

				toAddr := crypto.PubkeyToAddress(*publicKey)
				txData := &types.AccessListTx{
					To:    &toAddr,
					Value: big.NewInt(int64(tx.GetValue())),
					Nonce: uint64(tx.GetTxType()),
				}
				etx := types.NewTx(txData)
				txmb, err := etx.MarshalBinary()
				if err != nil {
					log.Warn("could not create transaction: %v", err)
					return
				}
				block.SetExtra(txmb)
				log.Info(hex.EncodeToString(txmb))
			} else if tx.GetTxType() == qtypes.TxTypeCrossChainImport {
				pubkBytes, err := hex.DecodeString(tx.GetFrom())
				if err != nil {
					log.Warn(err.Error())
					continue
				}
				publicKey, err := crypto.UnmarshalPubkey(pubkBytes)
				if err != nil {
					log.Warn(err.Error())
					continue
				}

				toAddr := crypto.PubkeyToAddress(*publicKey)
				txData := &types.AccessListTx{
					To:    &toAddr,
					Value: big.NewInt(int64(tx.GetValue())),
					Nonce: uint64(tx.GetTxType()),
				}
				etx := types.NewTx(txData)
				txmb, err := etx.MarshalBinary()
				if err != nil {
					log.Warn("could not create transaction: %v", err)
					return
				}
				block.SetExtra(txmb)
				log.Info(hex.EncodeToString(txmb))
			} else if tx.GetTxType() == qtypes.TxTypeCrossChainVM {
				txb := common.FromHex(string(tx.GetData()))
				var txmb = &types.Transaction{}
				if err := txmb.UnmarshalBinary(txb); err != nil {
					log.Error(fmt.Sprintf("rlp decoding failed: %v", err))
					continue
				}
				pubkBytes, err := hex.DecodeString(tx.GetTo())
				if err != nil {
					log.Warn(err.Error())
					continue
				}
				publicKey, err := crypto.UnmarshalPubkey(pubkBytes)
				if err != nil {
					log.Warn(err.Error())
					continue
				}

				toAddr := crypto.PubkeyToAddress(*publicKey)
				block.SetCoinbase(toAddr)
				block.AddTx(txmb)
			}
		}

	})
	if len(blocks) != 1 {
		return nil, fmt.Errorf("BuildBlock error")
	}
	num, err := vm.chain.Ether().BlockChain().InsertChainWithoutSealVerification(blocks[0])
	if err != nil {
		return nil, err
	}
	if num != 1 {
		return nil, fmt.Errorf("BuildBlock error")
	}

	log.Info(fmt.Sprintf("BuildBlock:number=%d hash=%s txs=%d,%d", blocks[0].Number().Uint64(), blocks[0].Hash().String(), len(blocks[0].Transactions()), len(txs)))

	h := hash.MustBytesToHash(blocks[0].Hash().Bytes())
	return &Block{id: &h, ethBlock: blocks[0], vm: vm, status: consensus.Accepted}, nil
}

func (vm *VM) ParseBlock([]byte) (consensus.Block, error) {
	return nil, nil
}

func (vm *VM) LastAccepted() (*hash.Hash, error) {
	block := vm.chain.Ether().BlockChain().CurrentBlock()
	h := hash.MustBytesToHash(block.Hash().Bytes())
	return &h, nil
}

func (vm *VM) GetBalance(addre string) (int64, error) {
	addr, err := address.DecodeAddress(addre)
	if err != nil {
		return 0, err
	}
	secpPksAddr, ok := addr.(*address.SecpPubKeyAddress)
	if !ok {
		return 0, fmt.Errorf("Not SecpPubKeyAddress:%s", addr.String())
	}
	publicKey, err := crypto.UnmarshalPubkey(secpPksAddr.PubKey().SerializeUncompressed())
	if err != nil {
		return 0, err
	}
	eAddr := crypto.PubkeyToAddress(*publicKey)
	state, err := vm.chain.Ether().BlockChain().State()
	if err != nil {
		return 0, err
	}
	return state.GetBalance(eAddr).Int64(), nil
}

func (vm *VM) VerifyTx(tx consensus.Tx) (int64, error) {
	if tx.GetTxType() == qtypes.TxTypeCrossChainVM {
		txb := common.FromHex(string(tx.GetData()))
		var txe = &types.Transaction{}
		if err := txe.UnmarshalBinary(txb); err != nil {
			return 0,fmt.Errorf("rlp decoding failed: %v", err)
		}
		err:=vm.validateTx(txe)
		if err != nil {
			return 0,err
		}
		return txe.Cost().Int64(),nil
	}
	return 0,fmt.Errorf("Not support")
}

func (vm *VM)  validateTx(tx *types.Transaction) error {
	if uint64(tx.Size()) > txMaxSize {
		return core.ErrOversizedData
	}
	if tx.Value().Sign() < 0 {
		return core.ErrNegativeValue
	}
	if tx.GasFeeCap().BitLen() > 256 {
		return core.ErrFeeCapVeryHigh
	}
	if tx.GasTipCap().BitLen() > 256 {
		return core.ErrTipVeryHigh
	}
	if tx.GasFeeCapIntCmp(tx.GasTipCap()) < 0 {
		return core.ErrTipAboveFeeCap
	}
	from, err := types.Sender(types.LatestSigner(vm.chain.Ether().BlockChain().Config()), tx)
	if err != nil {
		return core.ErrInvalidSender
	}
	currentState, err := vm.chain.Ether().BlockChain().State()
	if err != nil {
		return err
	}
	if currentState.GetNonce(from) > tx.Nonce() {
		return core.ErrNonceTooLow
	}
	if currentState.GetBalance(from).Cmp(tx.Cost()) < 0 {
		return core.ErrInsufficientFunds
	}
	intrGas, err := core.IntrinsicGas(tx.Data(), tx.AccessList(), tx.To() == nil, true, true)
	if err != nil {
		return err
	}
	if tx.Gas() < intrGas {
		return core.ErrIntrinsicGas
	}
	return nil
}

func (vm *VM) addTx(tx *types.Transaction) error {
	mtx := qtypes.NewTransaction()
	mtx.AddTxIn(&qtypes.TxInput{
		PreviousOut: *qtypes.NewOutPoint(&hash.ZeroHash, qtypes.SupperPrevOutIndex),
		Sequence:    uint32(qtypes.TxTypeCrossChainVM),
		AmountIn: qtypes.Amount{Id:qtypes.ETHID,Value:0},
		SignScript: []byte{},
	})
	txmb, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	txmbHex:=hexutil.Encode(txmb)
	mtx.AddTxOut(&qtypes.TxOutput{
		Amount:   qtypes.Amount{Value: 0, Id: qtypes.ETHID},
		PkScript: opreturn.NewEVMTx(txmbHex).PKScript(),
	})

	_,err=vm.ctx.GetTxPool().MaybeAcceptTransaction(qtypes.NewTx(mtx),true,false)
	return err
}

func (vm *VM) handler() {
	log.Debug("Meerevm handler start")
	defer vm.txsSub.Unsubscribe()

out:
	for {
		select {

		case ev := <-vm.txsCh:
			for _, tx := range ev.Txs {
				err:=vm.addTx(tx)
				if err != nil {
					log.Error(err.Error())
				}
			}

		case <-vm.shutdownChan:
			break out
		}
	}

cleanup:
	for {
		select {
		case <-vm.txsCh:
		default:
			break cleanup
		}
	}


	vm.shutdownWg.Done()
	log.Debug("Meerevm handler done")
}

func New() *VM {
	return &VM{
		shutdownChan: make(chan struct{}),
		txsCh:        make(chan core.NewTxsEvent, 256),
	}
}
