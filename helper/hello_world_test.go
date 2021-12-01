package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Jamey")
	if result != "Hello Jamey" {
		// error
		panic("result is not 'Hello Jamey'")
	}
}

func TestHelloWorldAnjali(t *testing.T) {
	result := HelloWorld("Anjali")
	if result != "Hello Anjali" {
		// error
		panic("result is not 'Hello Anjali'")
	}
}
