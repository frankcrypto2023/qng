package rawdb

import (
	"bytes"
	"github.com/Qitmeer/qng/common/hash"
	"github.com/Qitmeer/qng/meerdag"
	"github.com/ethereum/go-ethereum/ethdb"

	"github.com/Qitmeer/qng/core/types"
)

func ReadBlockRaw(db ethdb.Reader, hash *hash.Hash) []byte {
	var data []byte
	data, _ = db.Get(blockKey(hash))
	if len(data) > 0 {
		return data
	}
	blockID := ReadBlockID(db, hash)
	if blockID == nil {
		return nil
	}

	db.ReadAncients(func(reader ethdb.AncientReaderOp) error {
		data, _ = reader.Ancient(ChainFreezerBlockTable, *blockID)
		return nil
	})
	return data
}

func ReadBlock(db ethdb.Reader, hash *hash.Hash) *types.SerializedBlock {
	data := ReadBlockRaw(db, hash)
	if len(data) == 0 {
		return nil
	}
	block, err := types.NewBlockFromBytes(data)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return block
}

func WriteBlock(db ethdb.KeyValueWriter, block *types.SerializedBlock) error {
	data, err := block.Bytes()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	key := blockKey(block.Hash())
	err = db.Put(key, data)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func DeleteBlock(db ethdb.KeyValueWriter, hash *hash.Hash) {
	if err := db.Delete(blockKey(hash)); err != nil {
		log.Crit("Failed to delete hash to block mapping", "err", err)
	}
}

func HasBlock(db ethdb.Reader, hash *hash.Hash) bool {
	if has, err := db.Has(blockKey(hash)); !has || err != nil {
		return false
	}
	blockID := ReadBlockID(db, hash)
	return blockID != nil
}

func WriteAncientBlocks(db ethdb.AncientWriter, blocks []*types.SerializedBlock, dagblocks []meerdag.IBlock) (int64, error) {
	return db.ModifyAncients(func(op ethdb.AncientWriteOp) error {
		for i, block := range blocks {
			if err := writeAncientBlock(op, block, dagblocks[i]); err != nil {
				return err
			}
		}
		return nil
	})
}

func writeAncientBlock(op ethdb.AncientWriteOp, block *types.SerializedBlock, dagblock meerdag.IBlock) error {
	data, err := block.Bytes()
	if err != nil {
		return err
	}
	err = op.AppendRaw(ChainFreezerBlockTable, uint64(dagblock.GetID()), data)
	if err != nil {
		return err
	}
	err = op.AppendRaw(ChainFreezerDAGBlockTable, uint64(dagblock.GetID()), dagblock.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// header
func ReadHeaderRaw(db ethdb.Reader, hash *hash.Hash) []byte {
	var data []byte
	data, _ = db.Get(headerKey(hash))
	if len(data) > 0 {
		return data
	}
	blockID := ReadBlockID(db, hash)
	if blockID == nil {
		return nil
	}
	db.ReadAncients(func(reader ethdb.AncientReaderOp) error {
		data, _ = reader.Ancient(ChainFreezerHeaderTable, *blockID)
		return nil
	})
	return data
}

func HasHeader(db ethdb.Reader, hash *hash.Hash) bool {
	if has, err := db.Has(headerKey(hash)); !has || err != nil {
		return false
	}
	blockID := ReadBlockID(db, hash)
	return blockID != nil
}

func ReadHeader(db ethdb.Reader, hash *hash.Hash) *types.BlockHeader {
	data := ReadHeaderRaw(db, hash)
	if len(data) == 0 {
		return nil
	}
	var header types.BlockHeader
	err := header.Deserialize(bytes.NewReader(data))
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return &header
}

func WriteHeader(db ethdb.KeyValueWriter, header *types.BlockHeader) error {
	var headerBuf bytes.Buffer
	err := header.Serialize(&headerBuf)
	if err != nil {
		return err
	}
	data := headerBuf.Bytes()
	h := header.BlockHash()
	key := headerKey(&h)
	err = db.Put(key, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteHeader(db ethdb.KeyValueWriter, hash *hash.Hash) {
	if err := db.Delete(headerKey(hash)); err != nil {
		log.Crit("Failed to delete hash to header mapping", "err", err)
	}
}
