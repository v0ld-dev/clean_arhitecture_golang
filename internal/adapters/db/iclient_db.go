package db

type IDataBaseClient interface {
	GetConnection() interface{}
}
