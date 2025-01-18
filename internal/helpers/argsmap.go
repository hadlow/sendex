package helpers

import (
	"fmt"
	"strings"
)

func CreateArgsmap(args []string) (map[string]string, error) {
	argsMap := make(map[string]string)

	for _, _arg := range args {
		arg := strings.Split(_arg, "=")

		if len(arg) != 2 {
			return nil, fmt.Errorf("incorrect arg format")
		}

		argsMap[arg[0]] = arg[1]
	}

	return argsMap, nil
}
