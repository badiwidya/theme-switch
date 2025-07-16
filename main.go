package main

import (
	"fmt"
	"io"
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
		if err := switchTheme("light", configDir); err != nil {
			fmt.Printf("Error: failed to change theme.\n%v\n", err)
			os.Exit(1)
		}
	case "dark":
		if err := switchTheme("dark", configDir); err != nil {
			fmt.Printf("Error: failed to change theme.\n%v\n", err)
			os.Exit(1)
		}
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

	themeDir := filepath.Join(configDir, "themes", theme)

	for source, targetPath := range configurationMap {
		sourcePath := filepath.Join(themeDir, source)

		sourceFile, err := os.Open(sourcePath)
		if err != nil {
			return err
		}
		defer sourceFile.Close()

		targetFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer targetFile.Close()

		_, err = io.Copy(targetFile, sourceFile)
		if err != nil {
			return err
		}

		targetFile.Sync()

		fmt.Printf("Changing %s...\n", targetPath)
	}

	return nil
}
