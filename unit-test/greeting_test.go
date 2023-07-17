package main

import (
	"fmt"
	"testing"
)

func Test_sayHello_ValidArgument(t *testing.T) {
	inputs := []struct {
		name   string
		result string
	}{
		{name: "Yemeksepeti", result: "Hello Yemeksepeti"},
		{name: "Banabi", result: "Hello Banabi"},
		{name: "Yemek", result: "Hello Yemek"},
		{name: "", result: "Hello Anonymous"},
	}

	for _, item := range inputs {

		result := sayHello(item.name)
		if result != item.result {
			t.Errorf("\"sayHello('%s')\" failed, expected -> %v, got -> %v", item.name, item.result, result)
		} else {
			t.Logf("\"sayHello('%s')\" succeded, expected -> %v, got -> %v", item.name, item.result, result)
		}
	}
}

func Test_SayGoodBye(t *testing.T) {
	name := "Mert"
	expected := fmt.Sprintf("Bye Bye %s!", name)
	result := sayGoodBye(name)

	if result != expected {
		t.Errorf("\"sayGoodBye('%s')\" FAILED, expected -> %v, got -> %v", name, expected, result)
	} else {
		t.Logf("\"sayGoodBye('%s')\" SUCCEDED, expected -> %v, got -> %v", name, expected, result)
	}
}

func Example_sayHello() {
	fmt.Println(sayHello("Mert"))
	// Output: Hello Mert
}
