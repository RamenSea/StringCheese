package main

import (
	"os"
	"sort"
)

/*
	Writes the KotlinMap StringCheese api for a given StringKey struct
 */
func writeKotlinMapFiles(rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) error {
	pathToKotlinFile := config.pathToKotlinFolder + getKotlinMapFileName(config)

	_ = os.Remove(pathToKotlinFile) //ignore error
	//if err != nil {
	//	return err
	//}
	file, err := os.Create(pathToKotlinFile)
	if err != nil {
		return err
	}
	writeKotlinMapFileHeader(file, rootValue, values, config)

	for _, value := range values {
		writeKotlinMap(file,value,config)
	}

	file.WriteString("}")
	file.Close()
	return nil
}

func writeKotlinMapFileHeader(file *os.File, rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) {
	file.WriteString(config.timeStampString +
		"\n" +
		"//Generated String Cheese Kotlin Map file\n" +
		"package org.stringcheese\n" +
		"\n" +
		"class StringCheeseKotlinMapProvider() {\n")
	for _, value := range values {
		file.WriteString("    val " + getKotlinMapLanguageMethodNames(value, config) + ": Map<String, String> = create_" + getKotlinMapLanguageMethodNames(value, config)+ "()\n")
	}

	file.WriteString("    fun getLanguageMap(languageShortName: String): Map<String, String> {\n" +
		"        return when(languageShortName) {\n")
	for _, value := range values {
		file.WriteString("            \"" + value.languageId + "\" -> " + getKotlinMapLanguageMethodNames(value, config)+ "\n")
	}
	file.WriteString("            else -> " + getKotlinMapLanguageMethodNames(rootValue, config)+ "\n" +
		"    }\n"+
		"}\n")
}
func writeKotlinMap(file *os.File, value *StringKeys, config *StringCheeseConfig) {
	file.WriteString("\n" +
		"//Generating " + value.languageId + " String Cheese Kotlin map creator\n" +
		"    fun create_" + getKotlinMapLanguageMethodNames(value, config) +"(): Map<String, String> { \n" +
			"        val hashMap = HashMap<String, String>()\n")

	valueMap := value.strings
	useArgs := config.shouldCreateArguments

	keys := make([]string,0)
	for key, _ := range valueMap {
		keys = append(keys,key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := valueMap[key]
		if useArgs && value.numberOfArguments > 0 {

			//file.WriteString("	get raw_"+ jsTransformKeyToJSVarName(value.originalKey) +"() {" +
			//	" return '" + value.formatString + "';" +
			//	" }\n" )
			//file.WriteString("	format_"+ jsTransformKeyToJSVarName(value.originalKey) +"("+value.argumentString+") {" +
			//	" return `" + value.formatString + "`;" +
			//	" }\n" )

		} else {
			file.WriteString("        hashMap[\""+ kotlinMapTransformKeyTokotlinMapVarName(value.originalKey) +"\"] = \"" + value.value +"\"\n")
		}

	}
	file.WriteString("        return hashMap\n" +
		"    }\n\n")
}
