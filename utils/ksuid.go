package utils

import "github.com/segmentio/ksuid"

func GenerateKSUID() string {
	id := ksuid.New()
	return id.String()
}
