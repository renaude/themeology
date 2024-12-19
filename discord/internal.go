package discord

import (
	"log"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createBackground(css string) string {
	s := []string{"--background: #", strings.ReplaceAll(css, "background = #", ""), ";"}
	return strings.Join(s, "")
}

func createForeground(css string) string {
	s := []string{"--foreground: #", strings.ReplaceAll(css, "foreground = #", ""), ";"}
	return strings.Join(s, "")
}

func createCursor(css string) string {
	s := []string{"--cursor: #", strings.ReplaceAll(css, "cursor = #", ""), ";"}
	return strings.Join(s, "")
}

func createBorder(css string) string {
	s := []string{"--border: #", strings.ReplaceAll(css, "border = #", ""), ";"}
	return strings.Join(s, "")
}

func createColor(css string) string {
	if len(css) < 2 {
		return ""
	}
	cleanedcss := strings.ReplaceAll(css, css[0:strings.Index(css, "#")], "")
	s := []string{"--", strings.Trim(css[0:7], " "), ": ", cleanedcss, ";"}
	return strings.Join(s, "")
}
