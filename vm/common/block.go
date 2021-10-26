package common

import (
	"github.com/Qitmeer/qng/common/hash"
	"time"
)

type Block interface {
	Decidable
	Parent() *hash.Hash
	Verify() error
	Bytes() []byte
	Height() uint64
	Timestamp() time.Time
}
