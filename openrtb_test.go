package openrtb

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	t.Parallel()

	RegisterFailHandler(Fail)
	RunSpecs(t, "openrtb")
}

func fixture(fname string, v interface{}) error {
	f, err := os.Open(filepath.Join("testdata", fname+".json"))
	if err != nil {
		return err
	}
	defer f.Close() //nolint:errcheck

	return json.NewDecoder(f).Decode(v)
}
