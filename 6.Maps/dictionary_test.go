package Maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	got := dictionary.Search("test")
	want := "this is just a test"

	assertStrings(t, got, want)

	if got != want {

	}
}

// on page 69

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
