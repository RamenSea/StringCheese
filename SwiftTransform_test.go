package main

import "testing"

func TestSwiftTransformKeyToSwiftVarName(t *testing.T) {
	test1 := SwiftTransformKeyToSwiftVarName("test_1")
	if test1 != "test1" {
		t.Error("Error in: 'SwiftTransformKeyToSwiftVarName' func, failed test 1")
	}

	test2 := SwiftTransformKeyToSwiftVarName("test_test_test_2test")
	if test2 != "testTestTest2test" {
		t.Error("Error in: 'SwiftTransformKeyToSwiftVarName' func, failed test 2")
	}

	test3 := SwiftTransformKeyToSwiftVarName("2test")
	if test3 != "2test" {
		t.Error("Error in: 'SwiftTransformKeyToSwiftVarName' func, failed test 3")
	}

	test4 := SwiftTransformKeyToSwiftVarName("testtesttest")
	if test4 != "testtesttest" {
		t.Error("Error in: 'SwiftTransformKeyToSwiftVarName' func, failed test 4")
	}
}