package utils

import (
	"os"
	"log"
	"bufio"
	"errors"
)

func ReadLineFromFile(filepath string) string {
	if filepath != "" {
		file, err := os.Open(filepath)
		if err != nil {
			log.Fatal("could not find version file at:"+filepath, err)
		}
		defer file.Close()

		var scanner = bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		scanner.Scan()
		var version = scanner.Text()
		if version == "" {
			log.Fatal("reading version from version file", errors.New("your version file seems to be empty"))
		}
		// probably validate further
		return version
	}
	// else
	return ""
}
