package models

import (
	"encoding/json"
	"fmt"
	"github.com/rexposadas/art/util"
	"image/color"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Type   string
	Out    Out
	Canvas Canvas
	Colors Colors
}

type Colors struct {
	Scheme []color.RGBA
}

func DefaultConfig() *Config {
	defaultCfg := []byte(`
{
  "out": {
    "prefix": "samples",
    "dir": "output"
  },
  "canvas": {
    "width": 300,
    "height": 300
  },
  "colors": {
    "scheme": [
      {"r":111, "g": 84, "b": 140, "a":  255},
      {"r":92, "g": 204, "b": 206, "a":  255},
      {"r":178, "g": 162, "b": 150, "a":  255}
    ]
  }
}`)

	var cfg Config

	if err := json.Unmarshal(defaultCfg, &cfg); err != nil {
		return nil
	}

	return &cfg
}

func NewConfig(filename string) *Config {

	cfgFile, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to read file %q, %s", filename, err)
	}
	defer cfgFile.Close()

	b, err := ioutil.ReadAll(cfgFile)
	if err != nil {
		log.Fatalf("failed to read config file on %s - %s", filename, err)
	}

	var cfg Config
	if err := json.Unmarshal(b, &cfg); err != nil {
		log.Fatalf("failed to parse config %s", err)
	}

	return &cfg
}

// OutURL returns the default file location
func (cfg Config) OutURL() string {

	// We default to PNG , for now.
	var url string

	if cfg.Out.Dir == "" {
		url = fmt.Sprintf("%s", cfg.Out.Prefix)
	} else {
		url = fmt.Sprintf("%s/%s", cfg.Out.Dir, cfg.Out.Prefix)
	}

	url = fmt.Sprintf("%s-%d_x_%d_%s.png", url, cfg.Canvas.Width, cfg.Canvas.Height, util.RnID())

	return url
}

type Out struct {
	Prefix string // for naming files
	Dir    string // where the output of the process will go.
}

type Canvas struct {
	Width, Height int
}
