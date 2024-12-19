package waybar

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ParseHellwalWaybar(filepath string) []string {
	f, err := os.ReadFile(filepath)
	checkError(err)

	stringValues := strings.Split(string(f), "%%")
	return generateCSS(stringValues, "hellwal")
}

func ParseGowallWaybar(backendResponse *string) []string {
	stringValues := strings.Split(*backendResponse, "\n")
	return generateCSS(stringValues, "gowall")
}

func WriteWaybar(filepath string, waybarcss []string) {
	f, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range waybarcss {
		if element == "" {
			continue
		}
		_, err := f.WriteString(fmt.Sprintf("%s\n", element))
		if err != nil {
			log.Fatal(err)
		}
	}
	f.Sync()

	defer f.Close()
}

func CopyWaybar(frompath string, topath string) {
	f, err := os.Open(frompath)
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.Create(topath)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(out, f)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	defer out.Close()
}
