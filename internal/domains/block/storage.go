package block

type Storage interface {
	GetOne(id int) *Block
	GetAll() []*Block
	Create(block *Block)
	Delete(block *Block)
}
