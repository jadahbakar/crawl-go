package main

import (
	"fmt"

	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

func main() {

	geziyor.NewGeziyor(&geziyor.Options{
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			g.GetRendered("https://csr-assessment.miqdadyyy.vercel.app/", g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(string(r.Body))
		},
		// BrowserEndpoint: "ws://localhost:3000",
	}).Start()
}
