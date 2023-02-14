package handlers

import (
	"context"
	"fmt"
	"net/http"
	"proj1/data"
)

/*

func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the product
		errs := p.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}
		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})

}*/
func (p Products) MiddlewareValidateProductForPP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] Loi doc product", err)
			http.Error(rw, "Loi doc san pham", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] Loi validating", err)
			http.Error(
				rw,
				fmt.Sprintf("Loi validating san pham: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// them san pham vao context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// goi xu ly tiep theo,co the 1 middleware khac hoac cuoi cung.
		next.ServeHTTP(rw, r)
	})
}
