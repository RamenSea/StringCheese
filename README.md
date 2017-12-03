# StringCheese
### Overview:
Translates Android XML strings into iOS and Dart strings, comes with some high quality sugar if desired.
The key class StringCheese generates will handle referencing to the platforms own localization methods as 
well as provide methods to handle formatted strings.

### Quick use:

1. Run StringCheese.

    `./StringCheese -a <PATH_TO_ANDROID_RES_FOLDER> -i <PATH_TO_ROOT_OF_XCODE_PROJECT>`

    `./StringCheese -a <PATH_TO_ANDROID_RES_FOLDER> -dart <PATH_TO_DART_PROJECT>`

    `-a` is required and should be set to your Android's res folder

    `-i` is required if you want to output to XCode

    `-dart` is required if you want to output to Dart

2. If this is the first time running the program add the files to your XCode or Dart project.
3. Import and create the StringCheese class

### Features:

* Turns value-<language id>/strings.xml into <language id>.lproj/String.strings
* Creates turns Android strings into a Dart class structure
* Written in an extendable way to make it easy to add new targets in the future
* For iOS it creates a key class in Swift
* Both Dart and Swift key classes have string formatting methods with arguments. (ex: `stringCheese.someKeyWithFormats("arg1","arg2")`)
* Good defaults

### Todo:

* Add more argument types (currently only strings are supported)
* Pull strings from all xml files in a values folder rather than just one
* Extend translation to and from XML, CSV, or a spreadsheet. This should be pretty simple.
* Increase tests

### How to embed in another Go project:

1. Run `go get github.com/RamenSea/StringCheese`
2. Import StringCheese into your project
3. Create a StringCheese config struct `StringCheeseConfig`
4. Run the `RunStringCheese` function

### Argument Bucket List

`./StringCheese -h` for similar information

Arguments for pointing to projects. Pointing to an Android project is required.

| Arg | Type    | Description  |
| -------------- | ---------- | ------------ |
| a   | string | Path to your Android res folder.  |
| i   | string | Root folder of your iOS project. This is where the .strings files will be created.  |
| dart | string | This is where the Dart StringCheese classes will be generated  | 

Arguments for generating the translations

| Arg | Type    | Description  | Defaults |
| -------------- | ---------- | ------------ | ----------- |
| lang | string | Id of language to use as the base string xml file. For example en if you want the file 'values-en/strings.xml' to be used as the root language. If left blank the default folder will be values/strings.xml.  | none |
| langId | string| Id of language to use as the base id for -lang if none is specified. This value will be overwritten by -lang. For right now this is only used with Dart.  | en |
| className | string | The class name of the key file that will be generated  | StringCheese |
| logMissing | boolean | Logs to the console if there are any missing strings in your strings.xml files  | true |
| reduce | boolean | Reduces keys in outputted key classes. | true |
| s_xml | string | Name of your xml file containing strings. | strings |
| static | boolean | If the generated key file is static. | false |
| useArgs | boolean | If the key class generated generates a special argument method for strings with format points. | true |
| useValidLangId | boolean | Skips a values-LANGUAGE_ID folder if the id is not valid | true |

iOS specific arguments

| Arg | Type    | Description  | Defaults |
| -------------- | ---------- | ------------ | ----------- |
| dotName   | string | Name of the .strings file.  | String |
| swift   | boolean | If a swift key file should be generated.  | true |
| swiftLoc | string | Relative path inside the XCode project where the Swift StringCheese key class will be generated. If empty, this will generate it at the root of the project. | none |
