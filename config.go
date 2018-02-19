package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var cfg = map[string]string{}

const (
	s_comment = "#"
	s_split   = "="
)

func CfgParse(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, s_comment) == false && strings.Contains(line, s_split) == true {
			kv := strings.SplitN(line, s_split, 2)
			cfg[strings.Trim(kv[0], " ")] = strings.Trim(kv[1], " ")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func CfgString(key string) string {
	return cfg[key]
}

func CfgHasKey(key string) bool {
	_, ok := cfg[key]
	return ok
}

func CfgBool(key string) bool {
	b, _ := strconv.ParseBool(cfg[key])
	return b
}

func CfgInt(key string) int {
	i, _ := strconv.Atoi(cfg[key])
	return i
}
