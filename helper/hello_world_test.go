package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("============================= BEFORE RUNNING TEST =============================")

	m.Run()

	// after
	fmt.Println("============================= AFTER RUNNING TEST =============================")
}

func TestSubTest(t *testing.T) {
	t.Run("Ogi", func(t *testing.T) {
		result := HelloWorld("ogi")
		require.Equal(t, "hello ogi", result, "must be hello ogi")
	})

	t.Run("Aji", func(t *testing.T) {
		result := HelloWorld("aji")
		require.Equal(t, "hello aji", result, "must be hello aji")
	})
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("osgi")
	assert.Equal(t, "hello ogi", result, "error! result must be 'hello ogi'")
	fmt.Println("hello ini TestHelloWorldAssert")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("osgi")
	require.Equal(t, "hello ogi", result, "error! result must be 'hello ogi'")
}

func TestHelloWorldSkip(t *testing.T) {

	if runtime.GOOS == "linux" {
		t.Skip("unit test tidak bisa dijalankan di linux")
	}

	result := HelloWorld("osgi")
	require.Equal(t, "hello ogi", result, "error! result must be 'hello ogi'")
}

func TestHelloWorldOgi(t *testing.T) {
	result := HelloWorld("ogi")

	if result != "hello ogi" {
		t.Error("result must be hello ogi")
	}

	fmt.Println("hello TestHelloWorldOgi")
}

func TestHelloWorldAji(t *testing.T) {
	result := HelloWorld("aji")

	if result != "hello aji" {
		t.Error("result must be hello aji")
	}

	fmt.Println("hello TestHelloWorldaji")
}
