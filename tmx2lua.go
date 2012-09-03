package main

import (
	"flag"
	"github.com/salviati/go-tmx/tmx"
	"fmt"
	"os"
	"log"
)

func main() {
	flag.Parse()

	handle := flag.Arg(0)

	file, err := os.Open(handle)

	if err != nil {
		log.Fatal(err)
	}

	tileset, err := tmx.Read(file)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("return { ")
	fmt.Printf("width = \"%d\", ", tileset.Width)
	fmt.Printf("height = \"%d\", ", tileset.Height)
	fmt.Printf("tilewidth = \"%d\", ", tileset.TileWidth)
	fmt.Printf("tileheight = \"%d\", ", tileset.TileHeight)
	fmt.Printf("orientation = \"%s\", ", tileset.Orientation)
	fmt.Print(" properties = { ")

	for _, p := range tileset.Properties.Properties {
		fmt.Printf("[\"%s\"] = \"%s\", ", p.Name, p.Value)
	}

	fmt.Print(" }, tileset = { ")

	for _, t := range tileset.Tilesets {
		fmt.Printf("[\"%s\"] = \"%s\", ", p.Name, p.Value)
	}

	fmt.Print(" }, layers = { ")

	for _, t := range tileset.Layers {
		fmt.Printf("[\"%s\"] = \"%s\", ", p.Name, p.Value)
	}



	fmt.Print(" }")
}
