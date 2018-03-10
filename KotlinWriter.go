package main

import (
	"os"
	"sort"
)

/*
	Writes the KotlinMap StringCheese api for a given StringKey struct
 */
func writeKotlinFiles(rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) error {
	pathToKotlinFile := config.pathToKotlinFolder + getKotlinFileName(config)

	_ = os.Remove(pathToKotlinFile) //ignore error
	//if err != nil {
	//	return err
	//}
	file, err := os.Create(pathToKotlinFile)
	if err != nil {
		return err
	}
	writeKotlinFileHeader(file, rootValue, values, config)
	writeKotlinInterface(file, rootValue, values, config)

	for _, value := range values {
		writeKotlin(file,value,config)
	}

	file.Close()
	return nil
}

func writeKotlinFileHeader(file *os.File, rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) {
	file.WriteString(config.timeStampString +
		"\n" +
		"//Generated String Cheese Kotlin Map file\n" +
		"package org.stringcheese\n" +
		"\n" +
		"class StringCheeseProvider {\n")
	for _, value := range values {
		file.WriteString("    val " + getKotlinMemberName(value, config) + ": "+config.className+" = " + getKotlinClassName(value, config)+ "()\n")
	}

	file.WriteString("    fun get(languageShortName: String): " + config.className + " {\n" +
		"        return when(languageShortName) {\n")
	for _, value := range values {
		file.WriteString("            \"" + value.languageId + "\" -> " + getKotlinMemberName(value, config)+ "\n")
	}
	file.WriteString("            else -> " + getKotlinMemberName(rootValue, config)+ "\n" +
		"    }\n"+
		"}\n" +
			"}\n")
}
func writeKotlinInterface(file *os.File, rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) {
	file.WriteString("\n" +
		"interface "+config.className+" {\n")
	valueMap := rootValue.strings
	useArgs := config.shouldCreateArguments

	keys := make([]string,0)
	for key, _ := range valueMap {
		keys = append(keys,key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := valueMap[key]
		if useArgs && value.numberOfArguments > 0 {
			file.WriteString("	val raw_"+ kotlinTransformKeyToKotlinVarName(value.originalKey) + ": String\n")
			file.WriteString("	fun format_"+ kotlinTransformKeyToKotlinVarName(value.originalKey) + "("+value.argumentString+"): String\n")
		} else {
			file.WriteString("	val "+ kotlinTransformKeyToKotlinVarName(value.originalKey) + ": String\n")
		}
	}
	file.WriteString("}\n")
}
func writeKotlin(file *os.File, value *StringKeys, config *StringCheeseConfig) {
	file.WriteString("\n" +
		"//Generating " + value.languageId + " String Cheese\n" +
		"    class " + getKotlinClassName(value, config) +": StringCheese { \n")

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
			file.WriteString("	override val raw_"+ kotlinTransformKeyToKotlinVarName(value.originalKey) + ": String = \""+value.value+"\"\n")
			file.WriteString("	override fun format_"+ kotlinTransformKeyToKotlinVarName(value.originalKey) + "("+value.argumentString+"): String" +
				" = \""+value.formatString+"\"\n")
		} else {
			file.WriteString("	override val "+ kotlinTransformKeyToKotlinVarName(value.originalKey) + ": String = \""+value.value+"\"\n")
		}

	}
	file.WriteString("}\n")
}
