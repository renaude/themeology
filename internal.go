package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"themeology.local/config"
	"themeology.local/discord"
	"themeology.local/waybar"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ensureDir(dirpath string) {
	err := os.MkdirAll(dirpath, os.ModePerm)
	checkError(err)
}

func ensureFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	checkError(err)

	_, err = f.Seek(0, io.SeekStart)
	checkError(err)

	return f
}

func getRandomImage(wallpapersDir string) string {
	files, err := os.ReadDir(wallpapersDir)
	checkError(err)

	min := 0
	max := len(files) - 1

	randomIndex := rand.Intn(max-min+1) + min
	randomImage := files[randomIndex].Name()

	return randomImage
}

func extractConfigValue(configArray []string, key string) string {
	var result string
	for i := 0; i < len(configArray); i++ {
		value := configArray[i]
		if strings.Contains(value, key) {
			result = strings.ReplaceAll(value, fmt.Sprintf("%s=", key), "")
		}
	}
	return result
}

func executeHellwal(wallpaperPath string) (response *string, command string) {
	cmd := exec.Command("hellwal", "-i", wallpaperPath)
	stdout, err := cmd.Output()
	checkError(err)
	result := string(stdout)
	return &result, cmd.String()
}

func executeGowall(wallpaperPath string) (response *string, command string) {
	cmd := exec.Command("gowall", "extract", wallpaperPath, "-c", "16")
	stdout, err := cmd.Output()
	checkError(err)
	result := string(stdout)
	return &result, cmd.String()
}

func executeBackend(backend string, wallpaperPath string) (response *string, command string) {
	if backend == "hellwal" {
		return executeHellwal(wallpaperPath)
	} else if backend == "gowall" {
		return executeGowall(wallpaperPath)
	}
	return nil, ""
}

func executeSWWW(wallpaperPath string) (response *string, command string) {
	cmd := exec.Command("swww", "img", wallpaperPath)
	stdout, err := cmd.Output()
	checkError(err)
	result := string(stdout)
	return &result, cmd.String()
}

func executeWallpaperDaemon(daemon string, wallpaperPath string) (response *string, command string) {
	if daemon == "swww" {
		return executeSWWW(wallpaperPath)
	}
	return nil, ""
}

func executeThemecord(discordColors string) {
	cmd := exec.Command("themecord", "-f", discordColors)
	stdout, err := cmd.Output()
	checkError(err)
	fmt.Println(string(stdout))
}

func applyThemeology() {
	configuration := config.ParseConfig()

	wallpapers := extractConfigValue(configuration, "wallpapers")
	randomWallaper := getRandomImage(wallpapers)
	wallpaperPath := fmt.Sprintf("%s/%s", wallpapers, randomWallaper)

	backend := extractConfigValue(configuration, "backend")
	backendResponse, command := executeBackend(backend, wallpaperPath)
	fmt.Printf("Booting Backend: %s\n\n", command)
	fmt.Println(string(*backendResponse))

	wallpaperDaemon := extractConfigValue(configuration, "wallpaper_daemon")
	wallpaperDaemonResponse, command := executeWallpaperDaemon(wallpaperDaemon, wallpaperPath)
	fmt.Printf("Booting Wallpaper Daemon: %s\n", command)
	fmt.Println(string(*wallpaperDaemonResponse))

	fmt.Println("Generating CSS")
	var discordCSS []string
	var waybarCSS []string
	if backend == "hellwal" {
		cacheDir := extractConfigValue(configuration, "cache")
		hellwalFile := filepath.Join(cacheDir, "cache", fmt.Sprintf("%s.hellwal", randomWallaper))

		discordCSS = discord.ParseHellwalDiscord(hellwalFile)
		waybarCSS = waybar.ParseHellwalWaybar(hellwalFile)
	} else if backend == "gowall" {
		discordCSS = discord.ParseGowallDiscord(backendResponse)
		waybarCSS = waybar.ParseGowallWaybar(backendResponse)
	}
	cacheDir := extractConfigValue(configuration, "cache")
	ensureDir(cacheDir)

	discordOutput := filepath.Join(cacheDir, "discord-colors.css")
	ensureFile(discordOutput)
	discord.WriteDiscord(discordOutput, discordCSS)

	waybarConfig := extractConfigValue(configuration, "waybar_config")
	waybarOutput := filepath.Join(cacheDir, "waybar-colors.css")
	ensureFile(waybarOutput)
	waybar.WriteWaybar(waybarOutput, waybarCSS)
	waybar.CopyWaybar(waybarOutput, fmt.Sprintf("%s/colors.css", waybarConfig))

	executeThemecord(discordOutput)
}
