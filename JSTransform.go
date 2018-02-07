package main

import (
	"regexp"
	"strings"
)

func jsTransformKeyToJSVarName(key string) string {
	reg, _ := regexp.Compile(`(_).`)
	workingKey := reg.ReplaceAllStringFunc(key, func(rKey string) string {
		return strings.ToUpper(rKey)
	})
	return strings.Replace(workingKey, "_","",-1)
}

func getJSClassName(value *StringKeys,  config *StringCheeseConfig) string {
	if (value.languageId == LANGUAGE_ID_NONE) {
		return config.rootLanguageIdToUse + "_" + config.className
	}
	return value.languageId + "_" + config.className
}

func getJSFileName(config *StringCheeseConfig) string {
	return config.className + ".js"
}