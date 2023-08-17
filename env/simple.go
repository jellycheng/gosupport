package env

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type EnvKv struct {
	Key   string
	Value string
}

func GetEnvConfig(f string) (error, []EnvKv, map[string]string) {
	var envKvObj []EnvKv
	var envMap = make(map[string]string)
	fh, err := os.Open(f)
	if err != nil {
		return err, envKvObj, envMap
	}
	defer func(fh *os.File) {
		_ = fh.Close()
	}(fh)

	scanner := bufio.NewScanner(fh)
	reg := regexp.MustCompile(`^\s*([\w.-]+)\s*=\s*(.*)\s*$`)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// 空行
		if len(line) == 0 {
			continue
		}
		// 注释行
		if line[0] == '#' {
			continue
		}

		res := reg.FindStringSubmatch(line)
		if len(res) != 3 {
			return fmt.Errorf("config error in line %d", lineNo), envKvObj, envMap
		}
		envKvObj = append(envKvObj, EnvKv{
			Key:   res[1],
			Value: res[2],
		})
		envMap[res[1]] = res[2]
	}
	return nil, envKvObj, envMap
}
