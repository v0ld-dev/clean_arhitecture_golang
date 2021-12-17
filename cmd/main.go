package main

import (
	block3 "clean_arhitecture_golang/internal/adapters/api/block"
	"clean_arhitecture_golang/internal/adapters/db/block/mongo_db"
	block2 "clean_arhitecture_golang/internal/domains/block"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	_router := httprouter.New()
	_storageBook := mongo_db.NewStorage()
	_serviceBook := block2.NewService(_storageBook)

	handler := block3.NewHandler(_serviceBook)
	handler.Register(_router)

	http.ListenAndServe(":8085", _router)

}
