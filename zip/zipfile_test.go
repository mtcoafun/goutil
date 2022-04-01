package zip

import (
	"testing"
)

// test zip all files that under current dir's testdata dir to ./testdata
func TestZip(t *testing.T) {
	t.Run("test zip file", func(t *testing.T) {
		err := Zip("./testdata/test.zip", "./testdata/a.txt", "./testdata/b.txt")
		if err != nil {
			t.Error(err)
		}
	})
}
