# StringCheese
### Overview:
StringCheese is a simple string translator that turns Android's strings.xml files into iOS' .strings and a Dart string file. 
It also generates a key class to help implementing the strings in your application.

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

* Increase tests
* Clean up code
* Clean up README
* Consider extending translation to and from XML, CSV, or a spreadsheet. This should be pretty simple

### Argument Bucket List

`./StringCheese -h` for similar information

Arguments for pointing to projects

| Arg | Type    | Required   | Description  | Defaults |
| --- | ------- | ---------- | ------------ | -------- |
| a   | string  | yes          | Path to your Android res folder.  | none |
| i   | string  | yes(for ios) | Root folder of your iOS project. This is where the .strings files will be created.  | none |
| dart | string  | yes(for dart) | This is where the Dart StringCheese classes will be generated  | none |

Arguments for generating the translations

| Arg | Type    | Required   | Description  | Defaults |
| --- | ------- | ---------- | ------------ | -------- |
| lang | string  | no | Id of language to use as the base string xml file. For example en if you want the file 'values-en/strings.xml' to be used as the root language. If left blank the default folder will be values/strings.xml.  | none |
| langId | string  | no | Id of language to use as the base id for -lang if none is specified. This value will be overwritten by -lang. For right now this is only used with Dart.  | en |
| className | string  | no | The class name of the key file that will be generated  | StringCheese |
| logMissing | boolean  | no | Logs to the console if there are any missing strings in your strings.xml files  | true |
| reduce | boolean  | no | Reduces keys in outputted key classes. | true |
| s_xml | string  | no | Name of your xml file containing strings. | strings |
| static | boolean  | no | If the generated key file is static. | false |
| useArgs | boolean  | no | If the key class generated generates a special argument method for strings with format points. | true |
| useValidLangId | boolean  | no | Skips a values-LANGUAGE_ID folder if the id is not valid | true |

iOS specific arguments

| Arg | Type    | Required   | Description  | Defaults |
| --- | ------- | ---------- | ------------ | -------- |
| dotName   | string  | no | Name of the .strings file.  | String |
| swift   | boolean  | no | If a swift key file should be generated.  | true |
| swiftLoc | string  | no | Relative path inside the XCode project where the Swift StringCheese key class will be generated. If empty, this will generate it at the root of the project. | none |
