// A simple html page that presents a grid of squares where you can populate
// the current board along with a row where you place your tiles. The
// suggested words ranked by score are listed alongside.

package main

import (
	"github.com/dbtedman/kata-scrabbled/web"
	"log"
)

func main() {
	// TODO: Need to move app core setup here, including the repository data
	// TODO: Need to seed data currently populated in RenderSquares

	aWeb := web.Web{
		Addr: ":8080",
	}

	log.Fatal(aWeb.Listen())
}
