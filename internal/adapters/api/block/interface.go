package block

import (
	"clean_arhitecture_golang/internal/domains/block"
	"context"
)

type Service interface {
	GetBlockById(id int) *block.Block
	GetAll(ctx context.Context)  ([]block.Block, error)
	AddNewBlock(ctx context.Context, b *block.DtoBlockRequest) (map[string]string, error)
}
