package block

type Service interface {
	GetBlockById(id int) *Block
	GetAllBlocksById() []*Block
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return  &service{storage: storage}
}

func (s service) GetBlockById(id int) *Block {
	return s.storage.GetOne(id)
}

func (s service) GetAllBlocksById() []*Block {
	return s.storage.GetAll()
}

