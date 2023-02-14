package handlers

/*
import (
	"net/http"
	"proj1/data"
)

// Cap nhat handle put
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Dang cap nhat san pham", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}*/
