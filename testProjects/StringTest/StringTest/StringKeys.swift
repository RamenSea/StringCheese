// Last generated at: 2017-04-19 21:27:22.395626885 -0700 PDT
import Foundation

class StringKeys { 
	private static var _shared: StringKeys? = nil
	static var shared: StringKeys {
		if let s = _shared {
			return s
		}
		let s = StringKeys()
		_shared = s
		return s
	}

	private func localize(_ key: String) -> String { 
		return NSLocalizedString(key, comment: "")
	}
	var raw_argTest2: String {
		return localize("11")
	}
	func argTest2(_ arg1: String, arg2: String) -> String {
		let s = localize("11")
		return String(format: s, arg1, arg2)
	}
	var test3: String {
		return localize("0")
	}
	var raw_argTest1: String {
		return localize("1")
	}
	func argTest1(_ arg1: String) -> String {
		let s = localize("1")
		return String(format: s, arg1)
	}
	var raw_namedArgTest2: String {
		return localize("2")
	}
	func namedArgTest2(anotherValue: String, value2: String) -> String {
		let s = localize("2")
		return String(format: s, anotherValue, value2)
	}
	var appName: String {
		return "StringValueTest"
	}
	var test2: String {
		return localize("6")
	}
	var iAmMissingInOther: String {
		return localize("8")
	}
	var raw_namedArgTest3: String {
		return localize("3")
	}
	func namedArgTest3(anotherValue: String, value2: String) -> String {
		let s = localize("3")
		return String(format: s, anotherValue, value2)
	}
	var iAmMissingInRoot: String {
		return localize("4")
	}
	var raw_argTest3: String {
		return localize("7")
	}
	func argTest3(_ arg1: String, arg2: String, arg3: String) -> String {
		let s = localize("7")
		return String(format: s, arg1, arg2, arg3)
	}
	var raw_namedArgTest1: String {
		return localize("9")
	}
	func namedArgTest1(value1: String) -> String {
		let s = localize("9")
		return String(format: s, value1)
	}
	var test1: String {
		return localize("10")
	}
}