package tigerserver

import (
	"io/ioutil"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, clean := createTempFile(t, "xmen")
	defer clean()

	tape := &tape{file}
	tape.Write([]byte("nightcrawler"))

	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)

	got := string(newFileContents)
	want := "nightcrawler"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
