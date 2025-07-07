package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 3)
	fmt.Println(repeated)
	// Output: aaa
}

func TestStringContains(t *testing.T) {
	got := strings.Contains("Hello, World", "World")
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestStringsHasPrefix(t *testing.T) {
	got := strings.HasPrefix("Hello, World", "Hello")
	want := true
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestStringsHasSuffix(t *testing.T) {
	got := strings.HasSuffix("Hello, World", "World")
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestStringsJoin(t *testing.T) {
	words := []string{"Hello", "World", "from", "Go"}
	got := strings.Join(words, " ")
	want := "Hello World from Go"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestStringsSplit(t *testing.T) {
	got := strings.Split("apple,banana,orange", ",")
	want := []string{"apple", "banana", "orange"}

	if len(got) != len(want) {
		t.Errorf("got %v want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v want %v", got, want)
		}
	}
}

func TestSpringsToUpper(t *testing.T) {
	got := strings.ToUpper("hello, world")
	want := "HELLO, WORLD"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestStringsToLower(t *testing.T) {
	got := strings.ToLower("HELLO WORLD")
	want := "hello world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestStringsTrimSpace(t *testing.T) {
	got := strings.Trim(" Hello World  ", " ")
	want := "Hello World"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
