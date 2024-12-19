package config

import (
	"bufio"
	"os"
	"strings"
)

func ParseConfig() []string {
	var result []string

	homedir, err := os.UserHomeDir()
	checkError(err)

	configDirPath := getConfigDirPath(homedir)
	ensureDir(configDirPath)

	configFilePath := getConfigFilePath(configDirPath)
	configFile := ensureFile(configFilePath)

	// Now that we're sure the configuration exists...
	// Let's read the contents of the file
	scanner := bufio.NewScanner(bufio.NewReader(configFile))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		} else if strings.HasPrefix(line, "#") {
			continue
		} else {
			result = append(result, line)
		}
	}

	defer configFile.Close()

	return result
}
