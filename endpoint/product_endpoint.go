package endpoint

import (
	"ecom-product/dto/in"
	"ecom-product/middleware"
	"ecom-product/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	serviceFunc := func(r *http.Request) (interface{}, error, interface{}) {
		var req in.ProductDTOIn
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return nil, err, nil
		}

		resp, err := service.CreateProduct(&req)
		return resp, err, req
	}

	middleware.HandleRequest(w, r, serviceFunc)
}

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	serviceFunc := func(r *http.Request) (interface{}, error, interface{}) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid id"), nil
		}

		resp, err := service.GetProductByID(id)
		return resp, err, id
	}

	middleware.HandleRequest(w, r, serviceFunc)
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	serviceFunc := func(r *http.Request) (interface{}, error, interface{}) {
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		search := r.URL.Query().Get("search")

		req := in.GetListDTO{
			Limit:  limit,
			Offset: offset,
			Search: search,
		}

		resp, err := service.GetProducts(&req)
		return resp, err, req
	}

	middleware.HandleRequest(w, r, serviceFunc)
}
