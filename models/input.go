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
	url := fmt.Sprintf("%s/%s-%s.png", cfg.Out.Dir, cfg.Out.Prefix, util.RnID())
	return url
}

type Out struct {
	Prefix string // for naming files
	Dir    string // where the output of the process will go.
}

type Canvas struct {
	Width, Height int
}
