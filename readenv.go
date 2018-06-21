package envs

import (
	"os"
	"strings"
)

func Read() map[string]string {
	lines := os.Environ()
	key := make(map[string]string)
	for _, line := range lines {
		keyValue := strings.Split(line, "=")
		key[string(keyValue[0])] = keyValue[1]
	}
	return key
}
