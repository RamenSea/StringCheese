package main

import (
	"regexp"
	"strings"
)

func kotlinTransformKeyToKotlinVarName(key string) string {
	reg, _ := regexp.Compile(`(_).`)
	workingKey := reg.ReplaceAllStringFunc(key, func(rKey string) string {
		return strings.ToUpper(rKey)
	})
	return strings.Replace(workingKey, "_","",-1)
}

func getKotlinClassName(value *StringKeys,  config *StringCheeseConfig) string {
	if (value.languageId == LANGUAGE_ID_NONE) {
		return config.rootLanguageIdToUse + "_" + config.className
	}
	return value.languageId + "_" + config.className
}
func getKotlinMemberName(value *StringKeys,  config *StringCheeseConfig) string {
	if (value.languageId == LANGUAGE_ID_NONE) {
		return config.rootLanguageIdToUse + "_" + config.className
	}
	return value.languageId + "_" + config.className
}

func getKotlinFileName(config *StringCheeseConfig) string {
	return config.className + ".kt"
}