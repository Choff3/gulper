package main

import (
	"github.com/Choff3/gulper/gulpers"
	"github.com/Choff3/gulper/utils"
)

func main() {
	porterBeers := gulpers.GetPorterBeers()
	utils.StoreBeers(porterBeers)
}
