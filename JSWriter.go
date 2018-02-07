package main

import (
	"os"
	"sort"
)

/*
	Writes the JS StringCheese api for a given StringKey struct
 */
func writeJSFiles(rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) error {
	pathToJSFile := config.pathToJSFolder + getJSFileName(config)

	_ = os.Remove(pathToJSFile) //ignore error
	//if err != nil {
	//	return err
	//}
	file, err := os.Create(pathToJSFile)
	if err != nil {
		return err
	}
	writeJSFileHeader(file, rootValue, values, config)

	for _, value := range values {
		writeJSClass(file,value,config)
	}

	file.Close()
	return nil
}

func writeJSFileHeader(file *os.File, rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) {
	file.WriteString(config.timeStampString +
		"\n" +
		"//Generated String Cheese JS file" +
		"\n" +
		"function CreateStringCheeseWithLanguageShort(languageShortName) {\n" +
		"    switch(languageShortName) {\n")
	for _, value := range values {
		file.WriteString("        case '" + value.languageId + "': return new " + getJSClassName(value, config)+ "();\n")
	}
	file.WriteString("        default: return new " + getJSClassName(rootValue, config)+ "();\n" +
		"    }\n"+
		"}\n")
}
func writeJSClass(file *os.File, value *StringKeys, config *StringCheeseConfig) {
	file.WriteString("\n" +
		"//Generating " + value.languageId + " String Cheese class\n" +
		"class " + getJSClassName(value, config) +" { \n")

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

			file.WriteString("	get raw_"+ jsTransformKeyToJSVarName(value.originalKey) +"() {" +
				" return '" + value.formatString + "';" +
				" }\n" )
			file.WriteString("	format_"+ jsTransformKeyToJSVarName(value.originalKey) +"("+value.argumentString+") {" +
				" return `" + value.formatString + "`;" +
				" }\n" )

		} else {
			file.WriteString("	get "+ jsTransformKeyToJSVarName(value.originalKey) +"() {" +
				" return '" + value.value + "';" +
				" }\n" )
		}

	}
	file.WriteString("}")
}
