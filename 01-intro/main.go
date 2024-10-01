package main

import (
	"os"

	"github.com/mdp/qrterminal/v3"
)

func main() {
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig("hello world", config)
}