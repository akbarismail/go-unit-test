package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubTest(t *testing.T) {
	t.Run("Cleveland", func(t *testing.T) {
		result := HelloWorld("Cleveland")
		require.Equal(t, "Hello Cleveland", result, "Result must be 'Hello Cleveland'")
	})

	t.Run("Anastasia", func(t *testing.T) {
		result := HelloWorld("Anastasia")
		require.Equal(t, "Hello Anastasia", result, "Result must be 'Hello Anastasia'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Cleveland")
	require.Equal(t, "Hello Cleveland", result, "Result must be 'Hello Cleveland'")
	fmt.Println("TestSkip done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Cleveland")
	require.Equal(t, "Hello Cleveland", result, "Result must be 'Hello Cleveland'")
	fmt.Println("TestHelloWorld with require Cleveland")
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
