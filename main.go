package main

import (
	"fmt"
	"os"

	"github.com/dustinbowers/ascii-openssh-drunken-bishop/drunkenbishop"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 2 {
		fmt.Printf("Usage: %s <hash>\n", os.Args[0])
		os.Exit(1)
		return
	}
	fingerprint := os.Args[1]

	db := drunkenbishop.NewDrunkenBishop()
	db.SetTopLabel("")
	db.SetBottomLabel("")
	ascii, err := db.ToAscii(fingerprint)
	if err != nil {
		fmt.Printf("Error converting hex fingerprint to ascii: %v\n", err)
		os.Exit(1)
	}
	fmt.Print(ascii)
}
