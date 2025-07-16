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
	case "light", "dark":
		if err := switchTheme(themeSelected, configDir); err != nil {
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

	gtkThemes := map[string]string{"light": "adw-gtk3", "dark": "adw-gtk3-dark"}
	iconThemes := map[string]string{"light": "Papirus-Light", "dark": "Papirus-Light"}
	qtThemes := map[string]string{"light": "KvLibadwaita", "dark": "KvLibadwaitaDark"}

	commands := []*exec.Cmd{
		exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", gtkThemes[theme]),
		exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", iconThemes[theme]),
		exec.Command("kvantummanager", "--set", qtThemes[theme]),
		exec.Command("systemctl", "--user", "reload", "waybar"),
		exec.Command("systemctl", "--user", "restart", "dunst"),
		exec.Command("hyprctl", "reload"),
		exec.Command("kitty", "@", "set-colors", "-a", "-c", configurationMap["kitty-theme.conf"]),
	}

	for _, cmd := range commands {
		if err := cmd.Run(); err != nil {
			fmt.Printf("Warning: failed to run '%s'.\n%v\n", cmd.String(), err)
		} else {
			fmt.Printf("Command '%s' executed successfully\n", cmd.String())
		}
	}

	return nil
}
