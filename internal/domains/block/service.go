package block

import "clean_arhitecture_golang/internal/adapters/api/block"

type service struct {
	storage Storage
}

func NewService(storage Storage) block.Service {
	return  &service{storage: storage}
}

func (s service) GetBlockById(id int) *Block {
	return s.storage.GetOne(id)
}

func (s service) GetAllBlocksById() []*Block {
	return s.storage.GetAll()
}

