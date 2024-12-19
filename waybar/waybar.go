package waybar

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func CreateWaybarBackground(css string) string {
	s := []string{"@define-color background ", strings.Replace(css, "background = ", "", -1), ";"}
	return strings.Join(s, "")
}

func CreateWaybarForeground(css string) string {
	s := []string{"@define-color foreground ", strings.Replace(css, "foreground = ", "", -1), ";"}
	return strings.Join(s, "")
}

func CreateWaybarCursor(css string) string {
	s := []string{"@define-color cursor ", strings.Replace(css, "cursor = ", "", -1), ";"}
	return strings.Join(s, "")
}

func CreateWaybarBorder(css string) string {
	s := []string{"@define-color border ", strings.Replace(css, "border = ", "", -1), ";"}
	return strings.Join(s, "")
}

func CreateWaybarColor(css string) string {
	if len(css) < 2 {
		return ""
	}
	cleanedcss := strings.Replace(css, css[0:strings.Index(css, "#")], "", -1)
	s := []string{"@define-color ", strings.Trim(css[0:7], " "), " ", cleanedcss, ";"}
	return strings.Join(s, "")
}

func ParseWaybar(filepath string) []string {
	hellwalfile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var templatevalues []string = make([]string, 0)
	stringvalues := strings.Split(string(hellwalfile), "%%")
	for _, str := range stringvalues {
		if strings.Contains(str, "wallpaper") {
			continue
		}
		if strings.Contains(str, "background") {
			background := CreateWaybarBackground(str)
			templatevalues = append(templatevalues, background)
		} else if strings.Contains(str, "foreground") {
			foreground := CreateWaybarForeground(str)
			templatevalues = append(templatevalues, foreground)
		} else if strings.Contains(str, "cursor") {
			cursor := CreateWaybarCursor(str)
			templatevalues = append(templatevalues, cursor)
		} else if strings.Contains(str, "border") {
			border := CreateWaybarBorder(str)
			templatevalues = append(templatevalues, border)
		} else {
			color := CreateWaybarColor(str)
			templatevalues = append(templatevalues, color)
		}
	}
	return templatevalues
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
