package block

import "clean_arhitecture_golang/internal/domains/block"

type Service interface {
	GetBlockById(id int) *block.Block
	GetAllBlocksById()   []*block.Block
}
