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
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)
	h := simplehttp.New(wd, true)

	simplehttp.DefaultStaticData["Title"] = "MirBFT Visualizer"

	simplehttp.DefaultStaticData["CSSFiles"] = []string{
		"https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css",
	}

	simplehttp.DefaultStaticData["MetaTags"] = map[string]string{
		"viewport": "width=device-width, initial-scale=1, shrink-to-fit=no",
	}

	log.Fatal(http.ListenAndServe(l, h))
}
