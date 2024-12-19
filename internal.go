package main

import (
	"fmt"
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

func executeHellwal(wallpaperPath string) {
	cmd := exec.Command("hellwal", "-i", wallpaperPath)
	stdout, err := cmd.Output()
	checkError(err)
	fmt.Println(string(stdout))
}

func executeBackend(backend string, wallpaperPath string) {
	if backend == "hellwal" {
		executeHellwal(wallpaperPath)
	}
}

func executeSWWW(wallpaperPath string) {
	cmd := exec.Command("swww", "img", wallpaperPath)
	stdout, err := cmd.Output()
	checkError(err)
	fmt.Println(string(stdout))
}

func executeWallpaperDaemon(daemon string, wallpaperPath string) {
	if daemon == "swww" {
		executeSWWW(wallpaperPath)
	}
}

func executeThemecord() {
	cmd := exec.Command("themecord", "-g")
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
	fmt.Printf("Booting Backend: %s\n\n", backend)
	executeBackend(backend, wallpaperPath)

	wallpaperDaemon := extractConfigValue(configuration, "wallpaper_daemon")
	fmt.Printf("Booting Wallpaper Daemon: %s\n", wallpaperDaemon)
	executeWallpaperDaemon(wallpaperDaemon, wallpaperPath)

	fmt.Println("Generating CSS")
	var discordCSS []string
	var waybarCSS []string
	if backend == "hellwal" {
		cacheDir := extractConfigValue(configuration, "cache")
		hellwalFile := filepath.Join(cacheDir, fmt.Sprintf("%s.hellwal", randomWallaper))

		discordCSS = discord.ParseDiscord(hellwalFile)
		waybarCSS = waybar.ParseWaybar(hellwalFile)
	}
	discordOutput := extractConfigValue(configuration, "discord_output")
	discord.WriteDiscord(discordOutput, discordCSS)

	waybarConfig := extractConfigValue(configuration, "waybar_config")
	waybarOutput := extractConfigValue(configuration, "waybar_output")
	waybar.WriteWaybar(waybarOutput, waybarCSS)
	waybar.CopyWaybar(waybarOutput, fmt.Sprintf("%s/colors.css", waybarConfig))

	executeThemecord()
}
