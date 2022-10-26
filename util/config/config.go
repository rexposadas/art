package config

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

// Default to a small image for the sake of speed.
func Default() Config {
	defaultCfg := []byte(`
{
  "out": {
    "prefix": "sample",
    "dir": "output"
  },
  "canvas": {
    "width": 300,
    "height": 300
  }
}`)

	var cfg Config
	if err := json.Unmarshal(defaultCfg, &cfg); err != nil {
		return Config{}
	}

	return cfg
}

func New(filename string) Config {

	// Return a default config if no filename is provided.
	if filename == "" {
		return Default()
	}

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

	return cfg
}

// OutURL returns the default file location
func (cfg Config) OutURL() string {

	// We default to PNG , for now.
	var url string

	if cfg.Out.Dir == "" {
		cfg.Out.Dir = "output" // default to output dir
	}

	url = fmt.Sprintf("%s/%s", cfg.Out.Dir, cfg.Out.Prefix)

	url = fmt.Sprintf("%s-%dx%d_%s.png", url, cfg.Canvas.Width, cfg.Canvas.Height, util.RnID())

	if err := util.EnsureDir(cfg.Out.Dir); err != nil {
		log.Fatalf("failed to create dir %s - %s", cfg.Out.Dir, err)
	}

	return url
}

type Out struct {
	Prefix string // for naming files
	Dir    string // where the output of the process will go.
}

type Canvas struct {
	Width, Height int
}
