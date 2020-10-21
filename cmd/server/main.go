package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryan-hancock/restaurant-api/pkg/handler"
	"github.com/ryan-hancock/restaurant-api/pkg/storage/memory"
)

func main() {
	itemStor := memory.NewItemRepository()
	orderStor := memory.NewOrderRepository()

	router := mux.NewRouter()
	handler.Setup(router, itemStor, orderStor)

	addr := "127.0.0.1:8000"
	srv := &http.Server{
		Handler: router,
		Addr:    addr,
	}

	log.Printf("server listening on: %s", addr)
	log.Fatal(srv.ListenAndServe())
}
