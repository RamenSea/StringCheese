package main

import (
)

func kotlinMapTransformKeyTokotlinMapVarName(key string) string {
	return key
	//reg, _ := regexp.Compile(`(_).`)
	//workingKey := reg.ReplaceAllStringFunc(key, func(rKey string) string {
	//	return strings.ToUpper(rKey)
	//})
	//return strings.Replace(workingKey, "_","",-1)
}

func getKotlinMapLanguageMethodNames(value *StringKeys,  config *StringCheeseConfig) string {
	if (value.languageId == LANGUAGE_ID_NONE) {
		return config.rootLanguageIdToUse + "_" + config.className
	}
	return value.languageId + "_" + config.className
}

func getKotlinMapFileName(config *StringCheeseConfig) string {
	return config.className + ".kt"
}