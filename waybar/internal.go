package waybar

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

func createBackground(css string) string {
	s := []string{"@define-color background ", strings.Replace(css, "background = ", "", -1), ";"}
	return strings.Join(s, "")
}

func createForeground(css string) string {
	s := []string{"@define-color foreground ", strings.Replace(css, "foreground = ", "", -1), ";"}
	return strings.Join(s, "")
}

func createCursor(css string) string {
	s := []string{"@define-color cursor ", strings.Replace(css, "cursor = ", "", -1), ";"}
	return strings.Join(s, "")
}

func createBorder(css string) string {
	s := []string{"@define-color border ", strings.Replace(css, "border = ", "", -1), ";"}
	return strings.Join(s, "")
}

func createRule(key string, value string) string {
	s := []string{fmt.Sprintf("%s ", key), value, ";"}
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
				background := createBackground(str)
				templateValues = append(templateValues, background)
			} else if strings.Contains(str, "foreground") {
				foreground := createForeground(str)
				templateValues = append(templateValues, foreground)
			} else if strings.Contains(str, "cursor") {
				cursor := createCursor(str)
				templateValues = append(templateValues, cursor)
			} else if strings.Contains(str, "border") {
				border := createBorder(str)
				templateValues = append(templateValues, border)
			} else {
				if len(str) < 2 {
					continue
				}
				cleanedcss := strings.ReplaceAll(str, str[0:strings.Index(str, "#")], "")
				color := createRule(fmt.Sprintf("@define-color color%d", counter), cleanedcss)
				templateValues = append(templateValues, color)
				counter++
			}
		} else {
			color := createRule(fmt.Sprintf("@define-color color%d", counter), str)
			templateValues = append(templateValues, color)
			counter++
		}
	}
	return templateValues
}
