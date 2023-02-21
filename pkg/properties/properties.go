package properties

import (
	"bufio"
	"bytes"
	"errors"
	"strings"
)

func readFromBytes(data []byte) (map[string]string, error) {
	var result = map[string]string{}

	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// if line is empty or commented, skip it
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) < 2 {
			return nil, errors.New("invalid file format")
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(strings.Join(parts[1:], "="))

		result[key] = value
	}

	return result, nil
}
