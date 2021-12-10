package block

import "clean_arhitecture_golang/internal/domains/block"

type bookStorage struct {
	//client from pkg/client/mongodb/*.go
}

func (b bookStorage) GetOne(id int) *block.Block {
	panic("implement me")
}

func (b bookStorage) GetAll() []*block.Block {
	panic("implement me")
}

func (b bookStorage) Create(block *block.Block) {
	panic("implement me")
}

func (b bookStorage) Delete(block *block.Block) {
	panic("implement me")
}



