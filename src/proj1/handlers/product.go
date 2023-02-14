// Package classification of Product API
//
// Documentation for Product API222
//
// Schemes:http
// @host localhost:8080
// BasePath:/api
// Version:1.0.0
//
// Consumes:
// -application/json
//
// Produces
// -application/json
//swagger:meta

package handlers

import (
	"fmt"
	"log"
	"net/http"
	"proj1/data"

	//"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

// Products handler for lấy và cập nhật sp
type Products struct {
	l *log.Logger
}

// NewProducts  tra ve mot xu ly dc hien boi logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

/*func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "khong the ma hoa json", http.StatusInternalServerError)
	}
}*/
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Khong the convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Xu ly PUT product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProductI(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Khong tim thay san pham", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Khong tim thay san pham", http.StatusInternalServerError)
		return
	}
}

// ErrInvalidProductPath :thong bao loi khi duong dan san pham k dung
var ErrInvalidProductPath = fmt.Errorf("Duong dan sai, duong dan nen la /products/[id]")

// GenericError :tap hop thong bao loi chung dc may chu tra ve
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError : tap hop thong bao cac loi xac thuc
type ValidationError struct {
	Messages []string `json:"messages"`
}

//getProductID tra ve ID san pham tu URL
// loi panic neu ko chuyen dc id thanh so nguyen
func getProductID(r *http.Request) int {
	// chuyen id san pham tu url
	vars := mux.Vars(r)
	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//
		panic(err)
	}

	return id
}

/*func (p Products) MiddlewareValidateProductForPP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate(prod)
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})
}*/
