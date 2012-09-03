package main

import (
	"flag"
	"github.com/kyleconroy/go-tmx/tmx"
	"log"
	"os"
	"text/template"
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

	err = tmpl.Execute(os.Stdout, tileset)

	if err != nil {
		log.Fatal(err)
	}
}
