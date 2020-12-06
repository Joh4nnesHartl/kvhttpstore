package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Joh4nnesHartl/kvhttpstore/storage"
	"github.com/julienschmidt/httprouter"
)

var (
	kvstorage storage.KVStorage
)

func handlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")

	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	kvstorage.Store(key, value)

	w.WriteHeader(http.StatusOK)
}

func handleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")

	value, ok := kvstorage.Receive(key)
	if !ok {
		log.Printf("No value found at key %s", key)
		w.WriteHeader(http.StatusNoContent)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(value)
}

func main() {
	router := httprouter.New()
	router.GET("/:key", handleGet)
	router.POST("/:key", handlePost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
