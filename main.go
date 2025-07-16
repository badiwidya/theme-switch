package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("Error: can't get config path. Make sure your $HOME or $XDG_CONFIG_HOME is defined.\n%v\n", err)
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

func switchTheme(theme, configDir string) error {
	configurationMap := map[string]string{
		"kitty-theme.conf":     filepath.Join(configDir, "kitty", "theme.conf"),
		"waybar-colors.css":    filepath.Join(configDir, "waybar", "colors.css"),
		"tofi-colors":          filepath.Join(configDir, "tofi", "colors"),
		"hyprland-colors.conf": filepath.Join(configDir, "hypr", "colors.conf"),
		"dunstrc":              filepath.Join(configDir, "dunst", "dunstrc"),
	}

	return nil
}
