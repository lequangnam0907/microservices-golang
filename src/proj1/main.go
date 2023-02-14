package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"proj1/handlers"
	"time"

	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

func main() {
	env.Parse()
	l := log.New(os.Stdout, "Nam's demo Microservices ", log.LstdFlags)
	// tao mot handlers
	ph := handlers.NewProducts(l)

	// tao mot serve mux  va cac xu li
	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", ph.ListAll)
	getR.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products/{id:[0-9]+}", ph.UpdateProducts)
	putR.Use(ph.MiddlewareValidateProductForPP)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/products", ph.AddProduct)
	postR.Use(ph.MiddlewareValidateProductForPP)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", ph.Delete)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	//sm.Handle("/products", ph)
	// tao mot server moi
	s := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      ch(sm),            // set the default handler
		ErrorLog:     l,                 // dat log cho serve : de xem nhat ky hoat dong
		ReadTimeout:  5 * time.Second,   // Thoi gian doc toi da cho cac request
		WriteTimeout: 10 * time.Second,  // Thoi gian toi da de viet phan hoi toi nguoi dung
		IdleTimeout:  120 * time.Second, // Thoi gian toi da cho ket noi giu hoat dong
	}

	// bat dau server
	go func() {
		l.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// tao mot bay ki hieu
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	//Chan cho den khi tin hieu duoc nhan
	sig := <-c
	log.Println("Got signal:", sig)

	// tat may chu voi context, doi toi da 30s
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
