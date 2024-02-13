package controllers

import (
	"strings"
)
type DeleteMode int

const (
	SoftDelete DeleteMode = iota
	HardDelete
)

func Str2Enum(s string) DeleteMode {
	switch strings.ToLower(s) {
	case "soft":
		return SoftDelete
	case "hard":
		return HardDelete
	default: // TODO: error handling
		return SoftDelete
	}
}