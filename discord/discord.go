package discord

import (
	"fmt"
	"os"
	"strings"
)

func ParseDiscord(filepath string) []string {
	f, err := os.ReadFile(filepath)
	checkError(err)

	var templatevalues []string = make([]string, 0)
	stringvalues := strings.Split(string(f), "%%")
	for _, str := range stringvalues {
		if strings.Contains(str, "wallpaper") {
			continue
		}
		if strings.Contains(str, "background") {
			background := createBackground(str)
			templatevalues = append(templatevalues, background)
		} else if strings.Contains(str, "foreground") {
			foreground := createForeground(str)
			templatevalues = append(templatevalues, foreground)
		} else if strings.Contains(str, "cursor") {
			cursor := createCursor(str)
			templatevalues = append(templatevalues, cursor)
		} else if strings.Contains(str, "border") {
			border := createBorder(str)
			templatevalues = append(templatevalues, border)
		} else {
			color := createColor(str)
			templatevalues = append(templatevalues, color)
		}
	}
	return templatevalues
}

func WriteDiscord(filepath string, discordcss []string) {
	f, err := os.Create(filepath)
	checkError(err)

	for _, element := range discordcss {
		_, err := f.WriteString(fmt.Sprintf("%s\n", element))
		checkError(err)
	}
	f.Sync()

	defer f.Close()
}
