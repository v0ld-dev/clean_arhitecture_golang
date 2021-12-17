package simple_map

import (
	"clean_arhitecture_golang/internal/adapters/db"
	"clean_arhitecture_golang/internal/domains/block"
)

type bookStorage struct {
	clientDb db.IDataBaseClient
}

func NewStorage() block.Storage {
	return &bookStorage{}
}

func (b bookStorage) GetOne(id int) *block.Block {
	panic("implement me")
}

func (b bookStorage) GetAll() []*block.Block {
	//b.clientDb.getConnection().query("select * from block")
	panic("implement me")
}

func (b bookStorage) Create(block *block.Block) {
	panic("implement me")
}

func (b bookStorage) Delete(block *block.Block) {
	panic("implement me")
}



