package block

import (
	"clean_arhitecture_golang/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)





const (
	blocksURL = "/blocks"
)

type handler struct {
	blockService Service
}


func NewHandler(service Service) api.Handler {
	return &handler{blockService: service}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET(blocksURL, h.GetAllBlocks)
}

func (h *handler) GetAllBlocks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//id := h.blockService.GetAllBlocksById()
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}



