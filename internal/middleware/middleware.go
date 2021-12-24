package middleware

import (
	"clean_arhitecture_golang/internal/adapters/api"
	block2 "clean_arhitecture_golang/internal/domains/block"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler, route string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//validate inbound data
		if r.URL.Path == route {
			var newBlock block2.DtoBlockRequest
			err := validate(w, r, &newBlock)
			if err != nil {return}
		}

		//main func
		err := h(w, r)


		//catch all errors from all app layers
		if err != nil {
			responseWithError(w, err)
		}
	}

	return fn
}

func validate(w http.ResponseWriter, r *http.Request, v interface{}) error {
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, v)
	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		w.WriteHeader(http.StatusTeapot)
		w.Write(api.SystemError(err).Marshal())
		return err
	}
	return nil
}

func responseWithError(w http.ResponseWriter, err error) {

	var appErr *api.StatusResponse
	w.Header().Set("Content-Type", "application/json")
	if errors.As(err, &appErr) {
		if errors.Is(err, api.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			w.Write(api.ErrNotFound.Marshal())
			return
		}

		err = err.(*api.StatusResponse)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(appErr.Marshal())
		return
	}

	w.WriteHeader(http.StatusTeapot)
	w.Write(api.SystemError(err).Marshal())
}
