package utils

import (
	"github.com/go-kratos/kratos/v2/log"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"regexp"
)

const (
	UUIDPattern = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	NanoIDChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func IsValidUUID(input string) bool {
	regex := regexp.MustCompile(UUIDPattern)
	return regex.MatchString(input)
}

func NewNanoID() string {
	id, err := gonanoid.Generate(NanoIDChars, 21)
	if err != nil {
		log.Errorf("error occurred while generating nanoId err: %v ", err)
		return ""
	}
	return id
}
