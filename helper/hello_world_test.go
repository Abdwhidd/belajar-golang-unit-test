package helper

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {

	//BEFORE TEST
	fmt.Println("BEFORE TEST")

	m.Run()

	//AFTER TEST
	fmt.Println("AFTER TEST")

}

func TestSubTestSayHello(t *testing.T) {
	t.Run("Wahid", func(t *testing.T) {
		result := SayHello("Wahid")
		require.Equal(t, "Hello Wahid", result, "Must be result 'Hello Wahid'")
	})
	t.Run("Kahar", func(t *testing.T) {
		result := SayHello("Kahar")
		require.Equal(t, "Hello Kahar", result, "Must be result 'Hello Kahar'")
	})
}
