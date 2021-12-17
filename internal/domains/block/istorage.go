package block

import "context"

type Storage interface {
	GetOne(id int) *Block
	GetAll(ctx context.Context)  ([]Block, error)
	Create(ctx context.Context, block *Block) (string, error)
	Delete(block *Block)
}
