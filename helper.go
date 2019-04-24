package main

import (
	"strings"

	"github.com/agungdwiprasetyo/go-utils/stringprocessing"
)

func searchByValue(val string, m map[string]string) (string, string) {
	val = strings.ToLower(val)
	for k, v := range m {
		score := stringprocessing.Jaro(val, strings.ToLower(v))
		if score > Tolerance {
			return k, v
		}
	}
	return "", ""
}
