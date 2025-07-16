package main

import (
	"fmt"
	"os"
)

func main() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Error: can't load config directory. Make sure your $HOME or $XDG_CONFIG_DIR is defined.\n%v\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Printf("USAGE: theme-switch light|dark.")
		os.Exit(1)
	}

	themeSelected := os.Args[1]

	switch themeSelected {
	case "light":
	case "dark":
	default:
		fmt.Printf("USAGE: theme-switch light|dark.")
		os.Exit(1)
	}
}
