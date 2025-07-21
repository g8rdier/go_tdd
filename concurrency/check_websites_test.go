package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return url != "https://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"https://blog.gypsydave5.com",
		"https://furhurterwe.geds",
	}

	want := map[string]bool{
		"https://www.google.com":      true,
		"https://blog.gypsydave5.com": true,
		"https://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
