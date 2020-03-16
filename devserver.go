// +build ignore

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/vugu/vugu/simplehttp"
)

func main() {
	wd, _ := os.Getwd()
	l := "127.0.0.1:8845"
	log.Printf("Starting HTTP Server at %q", l)
	h := simplehttp.New(wd, true)

	simplehttp.DefaultStaticData["MetaTags"] = map[string]string{
		"viewport": "width=device-width, initial-scale=1, shrink-to-fit=no",
	}

	log.Fatal(http.ListenAndServe(l, h))
}
