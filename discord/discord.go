package discord

import (
	"fmt"
	"os"
	"strings"
)

func ParseHellwalDiscord(filepath string) []string {
	f, err := os.ReadFile(filepath)
	checkError(err)

	stringValues := strings.Split(string(f), "%%")
	return generateCSS(stringValues, "hellwal")
}

func ParseGowallDiscord(backendResponse *string) []string {
	stringValues := strings.Split(*backendResponse, "\n")
	return generateCSS(stringValues, "gowall")
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
