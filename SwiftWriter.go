package main

import (
	"os"
	"regexp"
	"strings"
)

/*
	Used to transform a key of a strings.xml to an idiomatic Swift var name
*/
func swiftTransformKeyToSwiftVarName(key string) string {
	reg, _ := regexp.Compile(`(_).`)
	workingKey := reg.ReplaceAllStringFunc(key, func(rKey string) string {
		return strings.ToUpper(rKey)
	})
	workingKey = strings.Replace(workingKey, "_", "", -1)
	workingKey = strings.Replace(workingKey, ".", "_", -1)
	return workingKey
}

/*
	Writes the Swift StringKey api for a given StringKey struct
*/
func writeSwiftKeyFile(value *StringKeys, config *StringCheeseConfig) error {
	pathToSwiftKey := config.pathToIOSProject + config.pathToSwiftKey + "/" + config.className + ".swift"
	_ = os.Remove(pathToSwiftKey) //skipped err check
	file, err := os.Create(pathToSwiftKey)
	if err != nil {
		return err
	}

	file.WriteString(config.timeStampString +
		"//This will be deleted and generated each time you run StringCheese.\n" +
		"import Foundation\n" +
		"\n" +
		"class " + config.className)

	if config.objCSupport {
		file.WriteString(": NSObject")
	}
	file.WriteString(" { \n")
	if config.createStaticKeyClass {
		file.WriteString("	private static var _shared: " + config.className + "? = nil\n" +
			"	static var shared: " + config.className + " {\n" +
			"		if let s = _shared {\n" +
			"			return s\n" +
			"		}\n" +
			"		let s = " + config.className + "()\n" +
			"		_shared = s\n" +
			"		return s\n" +
			"	}\n\n")
	}

	file.WriteString("	private func localize(_ key: String) -> String { \n" +
		"		return NSLocalizedString(key, comment: \"\")\n" +
		"	}\n")
	valueMap := value.strings

	writeArgSwiftFuncs := config.shouldCreateArguments

	var objCStringTag = ""
	if config.objCSupport {
		objCStringTag = "@objc "
	}
	for _, value := range valueMap {
		if value.translatable == false {
			file.WriteString("	" + objCStringTag + "let " + swiftTransformKeyToSwiftVarName(value.originalKey) + ": String = " +
				"\"" + value.value + "\"\n")
		} else if writeArgSwiftFuncs && value.numberOfArguments > 0 {
			//I added the raw string just incase
			file.WriteString("	" + objCStringTag + "var raw_" + swiftTransformKeyToSwiftVarName(value.originalKey) + ": String {\n" +
				"		return localize(\"" + value.key + "\")\n" +
				"	}\n")

			file.WriteString("	" + objCStringTag + "func " + swiftTransformKeyToSwiftVarName(value.originalKey) + "(" +
				value.argumentString + ") -> String {\n" +
				"		let s = localize(\"" + value.key + "\")\n" +
				"		return String(format: s, " + value.formatString + ")\n" +
				"	}\n")

		} else {
			file.WriteString("	" + objCStringTag + "var " + swiftTransformKeyToSwiftVarName(value.originalKey) + ": String {\n" +
				"		return localize(\"" + value.key + "\")\n" +
				"	}\n")
		}
	}
	file.WriteString("}")
	file.Close()

	return nil
}
