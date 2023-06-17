package z01

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("🚨 " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("🚨 Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}
	return scanner.Err()
}
