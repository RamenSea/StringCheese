package main

import (
	"strings"
	"os"
)

/*
	Writes the Dart StringCheese api for a given StringKey struct
 */
func writeDartFiles(rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) error {
	pathToDartFolder := config.pathToDartFile + "/" + strings.ToLower(config.className)

	err := os.RemoveAll(pathToDartFolder) //skipped err check
	if err != nil {
		return err
	}
	err = os.Mkdir(pathToDartFolder, os.ModePerm)
	if err != nil {
		return err
	}
	err = writeAbstractDartClass(rootValue, values, config)
	if err != nil {
		return err
	}
	err = writeOutSpecificDartFile(rootValue, config)
	if err != nil {
		return err
	}
	for _, value := range values {
		err = writeOutSpecificDartFile(value, config)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeAbstractDartClass(rootValue *StringKeys, values []*StringKeys, config *StringCheeseConfig) error {
	pathToDartFile := config.pathToDartFile + "/" + strings.ToLower(config.className) + "/" + getAbstractDartClassFileName(config)
	_ = os.Remove(pathToDartFile) //skipped err check
	file, err := os.Create(pathToDartFile)
	if err != nil {
		return err
	}

	file.WriteString(config.timeStampString +
		"//This will be deleted and generated each time you run StringCheese.\n" +
		"\n" +
		"library " + strings.ToLower(config.className) + ";\n\n" +
		"import \"dart:html\";\n\n" +
		"part \"" + getDartClassFileName(rootValue, config) + "\";\n")

	for _, value := range values {
		file.WriteString("part \"" + getDartClassFileName(value, config) + "\";\n")
	}

	file.WriteString("\n" +
		"\n" +
		"abstract class " + config.className + " { \n" +
		"	static StringCheese CreateFromLanguageId(String langId) { \n" +
		"		switch(langId) {\n" +
		"		case \"" + config.rootLanguageIdToUse + "\": return new " + getDartClassName(rootValue, config) +"();\n")
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
	useArgs := config.shouldCreateArguments

	for _, value := range valueMap {
		if value.translatable == false {
			file.WriteString("	final String " + dartTransformKeyToDartVarName(value.originalKey) +" = " +
				"\"" + value.value + "\";\n")
		} else if useArgs && value.numberOfArguments > 0 {
			file.WriteString("	String " + dartTransformKeyToDartVarName(value.originalKey) +"(" +
				value.argumentString + ");\n")

		} else {
			file.WriteString("	String get " + dartTransformKeyToDartVarName(value.originalKey) +";\n")
		}
	}

	file.WriteString("}")
	file.Close()

	return nil
}

func writeOutSpecificDartFile(value *StringKeys, config *StringCheeseConfig) error {
	pathToDartFile := config.pathToDartFile + "/" + strings.ToLower(config.className) + "/" + getDartClassFileName(value, config)

	_ = os.Remove(pathToDartFile) //skipped err check
	file, err := os.Create(pathToDartFile)
	if err != nil {
		return err
	}

	//
	file.WriteString(config.timeStampString +
		"\n" +
		"part of " + strings.ToLower(config.className) + ";\n" +
		"\n" +
		"class " + getDartClassName(value, config) + " extends "+ config.className + " { \n")

	valueMap := value.strings
	useArgs := config.shouldCreateArguments

	for _, value := range valueMap {
		if value.translatable == false {
			continue
		} else if useArgs && value.numberOfArguments > 0 {

			file.WriteString("	@override\n" +
				"	String " + dartTransformKeyToDartVarName(value.originalKey) +"(" +
				value.argumentString + ") {\n" +
				"		return \"" + value.formatString + "\";\n" +
				"	}\n" )

		} else {
			file.WriteString("	@override\n" +
				"	String get " + dartTransformKeyToDartVarName(value.originalKey) +" => \"" + value.value + "\";\n")
		}

	}
	file.WriteString("}")
	file.Close()

	return nil
}