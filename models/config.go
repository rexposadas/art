package models

import (
	"art/util"
	"encoding/json"
	"fmt"
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
    "dir": ""
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

	log.Printf("%+v", cfg)

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
		url = fmt.Sprintf("%s-%s.png", cfg.Out.Prefix, util.RnID())
	} else {
		url = fmt.Sprintf("%s/%s-%s.png", cfg.Out.Dir, cfg.Out.Prefix, util.RnID())
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
