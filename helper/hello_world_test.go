package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dan")
	require.Equal(t, "Hello Dan", result, "Result must be 'Hello Dan'")
	fmt.Println("TestHelloWorld with require done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Dan")
	assert.Equal(t, "Hello Dan", result, "Result must be 'Hello Dan'")
	fmt.Println("TestHelloWorld with assert done")
}

func TestHelloWorldJamey(t *testing.T) {
	result := HelloWorld("Jamey")
	if result != "Hello Jamey" {
		// error
		t.Error("Result must be 'Hello Jamey'")
	}

	fmt.Println("TestHelloWorldJamey done")
}

func TestHelloWorldAnjali(t *testing.T) {
	result := HelloWorld("Anjali")
	if result != "Hello Anjali" {
		// error
		t.Fatal("Result must be 'Hello Anjali'")
	}

	fmt.Println("TestHelloWorldAnjali done")
}
