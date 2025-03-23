package initializers

import (
	"bufio"
	"os"
	"strings"
)

//-----------------------------
// Config loader for .properties file
//-----------------------------
func LoadProperties(filename string) (map[string]string, error) {
	props := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines or comments starting with #
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Split on the first '=' character
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // or return error if desired
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		props[key] = value
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return props, nil
}