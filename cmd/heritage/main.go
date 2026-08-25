package main

import (
	"flag"
	"heritage/api"
	"heritage/content"
	"heritage/service"
	"heritage/storage"
	"log"
	"net/http"
)

func main() {
	path := flag.String("db", "heritage.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	st, e := storage.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	sv := service.New(st)
	for _, c := range content.SeedCategories() {
		if e = st.SaveCategory(c); e != nil {
			log.Fatal(e)
		}
	}
	for _, a := range content.SeedArticles() {
		if e = sv.PublishArticle(a); e != nil {
			log.Fatal(e)
		}
	}
	for _, c := range content.SeedCollections() {
		if e = sv.PublishCollection(c); e != nil {
			log.Fatal(e)
		}
	}
	log.Printf("heritage listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, api.New(sv).Handler()))
}
