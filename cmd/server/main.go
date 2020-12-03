package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Joh4nnesHartl/kvhttpstore/storage"
	"github.com/julienschmidt/httprouter"
)

var (
	kvstorage = make(storage.KVStorage)
)

func handlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	value, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Body read error: %s", err.Error())
		w.WriteHeader(500)

		return
	}

	key := ps.ByName("key")
	kvstorage.Store(key, value)

	w.WriteHeader(200)
}

func handleGet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")

	value, ok := kvstorage.Receive(key)
	if !ok {
		log.Printf("No value found at key %s", key)
		w.WriteHeader(404)

		return
	}

	w.WriteHeader(200)
	w.Write(value)
}

func main() {
	router := httprouter.New()
	router.GET("/:key", handleGet)
	router.POST("/:key", handlePost)

	log.Fatal(http.ListenAndServe(":8080", router))
}
