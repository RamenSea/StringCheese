package main

import (
	"regexp"
	"strings"
	"os"
)

/*
	Used to transform a key of a strings.xml to an idiomatic Swift var name
 */
func SwiftTransformKeyToSwiftVarName(key string) string {
	reg, _ := regexp.Compile(`(_).`)
	workingKey := reg.ReplaceAllStringFunc(key, func(rKey string) string {
		return strings.ToUpper(rKey)
	})
	return strings.Replace(workingKey, "_","",-1)
}
/*
	Writes the Swift StringKey api for a given StringKey struct
 */
func writeSwiftKeyFile(value *StringKeys, config *StringValueConfig) {
	pathToSwiftKey := config.pathToSwiftKey + config.swiftClassName + ".swift"
	_ = os.Remove(pathToSwiftKey) //skipped err check
	file, err := os.Create(pathToSwiftKey)
	if err != nil {
		//todo log
		return
	}

	file.WriteString(config.timeStampString +
		"import Foundation\n" +
		"\n" +
		"class " + config.swiftClassName + " { \n")

	if config.useStaticForSwiftAPI {
		file.WriteString("	private static var _shared: StringKeys? = nil\n" +
			"	static var shared: StringKeys {\n" +
			"		if let s = _shared {\n" +
			"			return s\n" +
			"		}\n" +
			"		let s = StringKeys()\n" +
			"		_shared = s\n" +
			"		return s\n" +
			"	}\n\n")
	}

	file.WriteString("	private func localize(_ key: String) -> String { \n" +
		"		return NSLocalizedString(key, comment: \"\")\n" +
		"	}\n")
	valueMap := value.strings

	writeArgSwiftFuncs := config.shouldCreateArgumentsInSwiftAPI

	for _, value := range valueMap {
		if value.translatable == false {
			file.WriteString("	var " + SwiftTransformKeyToSwiftVarName(value.originalKey) +": String {\n" +
				"		return \"" + value.value + "\"\n" +
				"	}\n")
		} else if writeArgSwiftFuncs && value.numberOfArguments > 0 {
			//I added the raw string just incase
			file.WriteString("	var raw_" + SwiftTransformKeyToSwiftVarName(value.originalKey) +": String {\n" +
				"		return localize(\"" + value.key + "\")\n" +
				"	}\n")

			file.WriteString("	func " + SwiftTransformKeyToSwiftVarName(value.originalKey) +"(" +
				value.argumentString + ") -> String {\n" +
				"		let s = localize(\"" + value.key +  "\")\n" +
				"		return String(format: s, " + value.formatString + ")\n" +
				"	}\n" )

		} else {
			file.WriteString("	var " + SwiftTransformKeyToSwiftVarName(value.originalKey) +": String {\n" +
				"		return localize(\"" + value.key + "\")\n" +
				"	}\n")
		}
	}
	file.WriteString("}")
	file.Close()
}
