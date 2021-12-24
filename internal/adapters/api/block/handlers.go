package block

import (
	"clean_arhitecture_golang/internal/adapters/api"
	block2 "clean_arhitecture_golang/internal/domains/block"
	"clean_arhitecture_golang/internal/middleware"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
)





const (
	blocksURL = "/blocks"
	AddBlock  = "/addBlock"
)

type handler struct {
	blockService Service
}


// use a single instance of Validate, it caches struct info
var validate *validator.Validate


func NewHandler(service Service) api.Handler {
	return &handler{blockService: service}
}
func (h *handler) Register(router *httprouter.Router) {
	router.GET(blocksURL, h.GetAllBlocks)
	//router.POST(addBlock, h.NewBlock)
	//router.GET("testerror", h.TestError)
	router.HandlerFunc(http.MethodGet, AddBlock,  middleware.Middleware(h.TestError, AddBlock))
	router.HandlerFunc(http.MethodPost, AddBlock, middleware.Middleware(h.NewBlock,  AddBlock))
}

func (h *handler) GetAllBlocks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//id := h.blockService.GetAllBlocksById()
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}

//TODO логи со всех слоев. Контекст проверить
func (h *handler) NewBlock(w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var newBlock block2.DtoBlockRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil { return err }
	json.Unmarshal(reqBody, &newBlock)

	resp, err := h.blockService.AddNewBlock(ctx, &newBlock)
	if err != nil { return err }

	st := api.NewStatusResponse(0, "", resp, nil, "")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(st)

	return nil
}

func (h *handler) TestError(w http.ResponseWriter, r *http.Request) error {
	//return exceptions.ErrNotFound
	return api.NewStatusResponse(-1,"test",nil,nil,"")
}



