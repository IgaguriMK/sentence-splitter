package main

import (
	"github.com/IgaguriMK/sentence-splitter/subcmd"
	_ "github.com/IgaguriMK/sentence-splitter/subcmd/split"
)

func main() {
	subcmd.RunApp("splitter", "Split & Extract text for translation.")
}
