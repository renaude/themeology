package discord

import (
	"fmt"
	"log"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createBackground(css string, backend string) string {
	var s []string
	if backend == "hellwal" {
		s = []string{"--background: #", strings.ReplaceAll(css, "background = #", ""), ";"}
	} else if backend == "gowall" {
		s = []string{"--background: ", css, ";"}
	}
	return strings.Join(s, "")
}

func createForeground(css string, backend string) string {
	var s []string
	if backend == "hellwal" {
		s = []string{"--foreground: #", strings.ReplaceAll(css, "foreground = #", ""), ";"}
	} else if backend == "gowall" {
		s = []string{"--foreground: ", css, ";"}
	}
	return strings.Join(s, "")
}

func createCursor(css string, backend string) string {
	var s []string
	if backend == "hellwal" {
		s = []string{"--cursor: #", strings.ReplaceAll(css, "cursor = #", ""), ";"}
	} else if backend == "gowall" {
		s = []string{"--cursor: ", css, ";"}
	}
	return strings.Join(s, "")
}

func createBorder(css string, backend string) string {
	var s []string
	if backend == "hellwal" {
		s = []string{"--border: #", strings.ReplaceAll(css, "border = #", ""), ";"}
	} else if backend == "gowall" {
		s = []string{"--border: ", css, ";"}
	}
	return strings.Join(s, "")
}

func createRule(key string, value string) string {
	s := []string{"--", key, ": ", value, ";"}
	return strings.Join(s, "")
}

func generateCSS(stringValues []string, backend string) []string {
	var templateValues []string = make([]string, 0)
	counter := 0
	for _, str := range stringValues {
		if strings.Contains(str, "wallpaper") {
			continue
		}
		if str == "" {
			continue
		}
		if backend == "hellwal" {
			if strings.Contains(str, "background") {
				background := createBackground(str, backend)
				templateValues = append(templateValues, background)
			} else if strings.Contains(str, "foreground") {
				foreground := createForeground(str, backend)
				templateValues = append(templateValues, foreground)
			} else if strings.Contains(str, "cursor") {
				cursor := createCursor(str, backend)
				templateValues = append(templateValues, cursor)
			} else if strings.Contains(str, "border") {
				border := createBorder(str, backend)
				templateValues = append(templateValues, border)
			} else {
				if len(str) < 2 {
					continue
				}
				cleanedcss := strings.ReplaceAll(str, str[0:strings.Index(str, "#")], "")
				color := createRule(fmt.Sprintf("color%d", counter), cleanedcss)
				templateValues = append(templateValues, color)
				counter++
			}
		} else {
			color := createRule(fmt.Sprintf("color%d", counter), str)
			templateValues = append(templateValues, color)
			if counter == 0 {
				background := createRule("background", str)
				templateValues = append(templateValues, background)
			} else if counter == 10 {
				foreground := createRule("foreground", str)
				templateValues = append(templateValues, foreground)
				cursor := createRule("cursor", str)
				templateValues = append(templateValues, cursor)
				border := createRule("border", str)
				templateValues = append(templateValues, border)
			}
			counter++
		}
	}
	return templateValues
}
