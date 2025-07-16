package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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
			fmt.Printf("Warning: source file %s not found, skipping...\n", sourcePath)
			continue
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

	if theme == "dark" {
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", "adw-gtk3-dark")
		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("GTK theme changed to dark.\n")
		}

		cmd = exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", "Papirus-Dark")
		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("Icon theme changed to dark.\n")
		}

		cmd = exec.Command("kvantummanager", "--set", "KvLibadwaitaDark")
		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("QT theme changed to dark.\n")
		}
	} else {
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", "adw-gtk3")
		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("GTK theme changed to light.\n")
		}

		cmd = exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", "Papirus-Light")
		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("Icon theme changed to light.\n")
		}

		cmd = exec.Command("kvantummanager", "--set", "KvLibadwaita")
		if err := cmd.Run(); err != nil {
			return err
		} else {
			fmt.Printf("QT theme changed to light.\n")
		}
	}

	return nil
}
