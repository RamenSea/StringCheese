package main

import (
	"regexp"
	"strings"
	"os"
)

func DartTransformKeyToDartVarName(key string) string {
	reg, _ := regexp.Compile(`(_).`)
	workingKey := reg.ReplaceAllStringFunc(key, func(rKey string) string {
		return strings.ToUpper(rKey)
	})
	return strings.Replace(workingKey, "_","",-1)
}
func getDartClassName(value *StringKeys,  config *StringValueConfig) string {
	if (value.languageId == LANGUAGE_ID_NONE) {
		return config.idOfNoLanguage + "_" + config.swiftClassName
	}
	return value.languageId + "_" + config.swiftClassName
}
func getDartClassFileName(value *StringKeys,  config *StringValueConfig) string {
	return getDartClassName(value, config) + ".dart"
}
func getAbstractDartClassFileName(config *StringValueConfig) string {
	return config.swiftClassName + ".dart"
}
/*
	Writes the Swift StringKey api for a given StringKey struct
 */
func writeDartFile(rootValue *StringKeys, values []*StringKeys, config *StringValueConfig) {
	pathToDartFolder := config.pathToDartFile + "/" + strings.ToLower(config.swiftClassName)

	err := os.RemoveAll(pathToDartFolder) //skipped err check
	if err != nil {
		print(err.Error())
		return
	}
	err = os.Mkdir(pathToDartFolder, os.ModePerm)
	if err != nil {
		print(err.Error())
		return
	}
	writeAbstractDartClass(rootValue, values, config)
	writeOutSpecificDartFile(rootValue, config)
	for _, value := range values {
		writeOutSpecificDartFile(value, config)
	}
}
func writeAbstractDartClass(rootValue *StringKeys, values []*StringKeys, config *StringValueConfig) {
	pathToDartFile := config.pathToDartFile + "/" + strings.ToLower(config.swiftClassName) + "/" + getAbstractDartClassFileName(config)
	_ = os.Remove(pathToDartFile) //skipped err check
	file, err := os.Create(pathToDartFile)
	if err != nil {
		//todo log
		return
	}

	file.WriteString(config.timeStampString +
		"\n" +
		"library " + strings.ToLower(config.swiftClassName) + ";\n\n" +
		"import \"dart:html\";\n\n" +
		"part \"" + getDartClassFileName(rootValue, config) + "\";\n")

	for _, value := range values {
		file.WriteString("part \"" + getDartClassFileName(value, config) + "\";\n")
	}

	//if config.useStaticForSwiftAPI {
	//	file.WriteString("	private static var _shared: " + config.swiftClassName + "? = nil\n" +
	//		"	static var shared: " + config.swiftClassName + " {\n" +
	//		"		if let s = _shared {\n" +
	//		"			return s\n" +
	//		"		}\n" +
	//		"		let s = " + config.swiftClassName + "()\n" +
	//		"		_shared = s\n" +
	//		"		return s\n" +
	//		"	}\n\n")
	//}

	//file.WriteString("	private func localize(_ key: String) -> String { \n" +
	//	"		return NSLocalizedString(key, comment: \"\")\n" +
	//	"	}\n")
	file.WriteString("\n" +
		"\n" +
		"abstract class " + config.swiftClassName + " { \n" +
		"	static StringCheese CreateFromLanguageId(String langId) { \n" +
		"		switch(langId) {\n" +
		"		case \"" + config.idOfNoLanguage + "\": return new " + getDartClassName(rootValue, config) +"();\n")
	for _, value := range values {
		file.WriteString("		case \"" + value.languageId + "\": return new " + getDartClassName(value, config) +"();\n")
	}
	file.WriteString("		default: return new " + getDartClassName(rootValue, config) +"();\n" +
		"		}\n" +
		"	}\n\n" +
		"	static StringCheese CreateFromNavigator() { \n" +
		"		if (window.navigator.language.length >= 2) { \n" +
		"			final lang = window.navigator.language.substring(0,2).toLowerCase(); \n" +
		"			return CreateFromLanguageId(lang); \n" +
		"		} \n" +
		"		return CreateFromLanguageId(\"\"); \n" +
		"	}\n\n"	)

	valueMap := rootValue.strings

	writeArgSwiftFuncs := config.shouldCreateArgumentsInSwiftAPI

	for _, value := range valueMap {
		if value.translatable == false {
			file.WriteString("	final String " + DartTransformKeyToDartVarName(value.originalKey) +" = " +
				"\"" + value.value + "\";\n")
		} else if writeArgSwiftFuncs && value.numberOfArguments > 0 {
			//I added the raw string just incase
			//file.WriteString("	final String raw_" + DartTransformKeyToDartVarName(value.originalKey) +";\n")

			file.WriteString("	String " + DartTransformKeyToDartVarName(value.originalKey) +"(" +
				value.argumentString + ");\n")

		} else {
			file.WriteString("	String get " + DartTransformKeyToDartVarName(value.originalKey) +";\n")
		}
	}
	//for _, value := range valueMap {

	//}
	file.WriteString("}")
	file.Close()
}

func writeOutSpecificDartFile(value *StringKeys, config *StringValueConfig) {
	pathToDartFile := config.pathToDartFile + "/" + strings.ToLower(config.swiftClassName) + "/" + getDartClassFileName(value, config)

	_ = os.Remove(pathToDartFile) //skipped err check
	file, err := os.Create(pathToDartFile)
	if err != nil {
		//todo log
		return
	}

	//
	file.WriteString(config.timeStampString +
		"\n" +
		"part of " + strings.ToLower(config.swiftClassName) + ";\n" +
		"\n" +
		"class " + getDartClassName(value, config) + " extends "+ config.swiftClassName + " { \n")

	valueMap := value.strings

	writeArgSwiftFuncs := config.shouldCreateArgumentsInSwiftAPI

	for _, value := range valueMap {
		if value.translatable == false {
			continue
		} else if writeArgSwiftFuncs && value.numberOfArguments > 0 {

			file.WriteString("	@override\n" +
				"	String " + DartTransformKeyToDartVarName(value.originalKey) +"(" +
				value.argumentString + ") {\n" +
				"		return \"" + value.formatString + "\";\n" +
				"	}\n" )

		} else {
			file.WriteString("	@override\n" +
			 "	String get " + DartTransformKeyToDartVarName(value.originalKey) +" => \"" + value.value + "\";\n")
		}

	}
	//for _, value := range valueMap {

	//}
	file.WriteString("}")
	file.Close()
}