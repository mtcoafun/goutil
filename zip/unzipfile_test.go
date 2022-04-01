package zip

import "testing"

// test unzip current's testdata/test.zip to ./testdata
func TestUnzip(t *testing.T) {
	t.Run("test unzip file", func(t *testing.T) {
		err := Unzip("./testdata", "./testdata/test.zip")
		if err != nil {
			t.Error(err)
		}
	})
}
