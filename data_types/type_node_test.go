package data_types

import "testing"

func setUp() {
	allNodeTypes.m = make(map[string]nodeType)
}

func tearDown() {
	allNodeTypes.m = make(map[string]nodeType)
}

func TestCreateNewNodeType(t *testing.T) {
	setUp()
	err := CreateNewNodeType("Your", "Moma")
	if err != nil {
		t.Error("Expected no error got", err)
	}

	val, present := allNodeTypes.m["Your"]
	if !present {
		t.Error("Expected saved NodeType was not created.")
	}

	if val.name != "Your" || val.description != "Moma" {
		t.Error("Expected Your Moma and got", val.name,
			val.description)
	}

	err = CreateNewNodeType("Your", "Moma")
	if err == nil {
		t.Error("FAILED duplicated one node type")
	}

	expected := "A NodeType with this name has already been created"
	actual := err.Error()
	if actual != expected {
		t.Error("Wrong error message expected but got",
			expected, actual)
	}

	err = CreateNewNodeType("Moma", "Your")
	if err != nil {
		t.Error("Expected no error for second node got", err)
	}

	val, present = allNodeTypes.m["Moma"]
	if !present {
		t.Error("Expected second saved NodeType was not created.")
	}

	if val.name != "Moma" || val.description != "Your" {
		t.Error("Expected Moma Your and got", val.name,
			val.description)
	}
	tearDown()
}

func TestGetNodeType(t *testing.T) {
	setUp()
	CreateNewNodeType("Your", "Moma")
	val, err := GetNodeType("Your")
	if err != nil {
		t.Error("Failed to get valid node with error", err)
	}

	if val.name != "Your" || val.description != "Moma" {
		t.Error("Expected Your Moma and got", val.name,
			val.description)
	}

	val, err = GetNodeType("FAIL")
	tmp := nodeType{}
	if *val != tmp {
		t.Error("Didn't return empty nodeType")
	}

	expected := "No NodeType with this name exists"
	actual := err.Error()
	if expected != actual {
		t.Error("Failed error message expected but got",
			expected, actual)
	}

	CreateNewNodeType("Moma", "Your")
	val, err = GetNodeType("Moma")
	if err != nil {
		t.Error("Failed to get second valid node with error", err)
	}
	tearDown()
}

func TestGetOrCreateNodeType(t *testing.T) {
	setUp()
	_, err := GetOrCreateNodeType("Your", "Moma")
	if err != nil {
		t.Error("Failed to create with error", err)
	}

	val, present := allNodeTypes.m["Your"]
	if !present {
		t.Error("Expected saved NodeType was not created.")
	}

	if val.name != "Your" || val.description != "Moma" {
		t.Error("Expected Your Moma and got", val.name,
			val.description)
	}

	_, err = GetOrCreateNodeType("Your", "Moma")
	if err != nil {
		t.Error("Failed to get with error", err)
	}

	_, err = GetOrCreateNodeType("Moma", "Your")
	if err != nil {
		t.Error("Failed to create with error", err)
	}
	tearDown()
}
