package main

import (
	"flag"
	"fmt"
	"github.com/kyleconroy/go-tmx/tmx"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func exportTileset(filename string, tmpl *template.Template) {
	fmt.Printf("Processing %s\n", filename)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	tileset, err := tmx.Read(file)

	if err != nil {
		log.Fatal(err)
	}

	luaFilename := strings.Replace(filepath.Base(filename), filepath.Ext(filename), ".lua", 1)
	outputPath := filepath.Join(filepath.Dir(filename), luaFilename)

	output, err := os.Create(outputPath)

	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(output, tileset)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	luaTable := `return {
  width = {{.Width}},
  height = {{.Height}},
  tilewidth = {{.TileWidth}},
  tileheight = {{.TileHeight}},
  orientation = "{{.Orientation}}",
  properties = { {{range .Properties}}
    ["{{.Name}}"] = "{{.Value}}",{{end}}
  },
  tilesets = { {{range .Tilesets}}
    {
      name = "{{.Name}}",
      tilewidth = {{.TileWidth}},
      tileheight = {{.TileHeight}},
      spacing = {{.Spacing}},
      margin = {{.Margin}},
      image = {
        source = "{{.Image.Source}}",
        width = "{{.Image.Width}}",
        height = "{{.Image.Height}}",
      },
      properties = { {{range .Properties}}
        ["{{.Name}}"] = "{{.Value}}",{{end}}
      },
    },{{end}}
  },
  tilelayers = { {{range .Layers }}
    {
      name = "{{.Name}}",
      properties = { {{range .Properties}}
        ["{{.Name}}"] = "{{.Value}}",{{end}}
      },
      tiles = { {{range .DecodedTiles}}{id = {{.ID}},{{if .HorizontalFlip}} flipHorizontal = true,{{end}}{{if .VerticalFlip}} flipVertical = true,{{end}}{{if .DiagonalFlip}} flipDiagonal = true,{{end}}},{{end}} },
    },{{end}}
  },
  objectgroups = { {{range .ObjectGroups }}
    {
      name = "{{.Name}}",
      properties = { {{range .Properties}}
        ["{{.Name}}"] = "{{.Value}}",{{end}}
      },
      objects = { {{range .Objects}}
        {
          name = "{{.Name}}",
          x = {{.X}},
          y = {{.Y}},
          width = {{.Width}},
          height = {{.Height}},
          type = "{{.Type}}",
          {{if .PolyLines}}polyline = { {{range .PolyLines }}{{range .Decode}}
            { x = {{.X}}, y = {{.Y}} },{{end}}{{end}}
          },{{end}}
          {{if .Polygons}}polygon = { {{range .Polygons }}{{range .Decode}}
            { x = {{.X}}, y = {{.Y}} },{{end}}{{end}}
          },{{end}}
          properties = { {{range .Properties}}
            ["{{.Name}}"] = "{{.Value}}",{{end}}
          },
        },{{end}}
      },
    },{{end}}
  }
}`
	tmpl, err := template.New("lua").Parse(luaTable)

	if err != nil {
		log.Fatal(err)
	}

	for _, handle := range flag.Args() {
		exportTileset(handle, tmpl)
	}
}
