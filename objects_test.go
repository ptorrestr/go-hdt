package hdt

import (
	"github.com/op/go-logging"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"testing"
)

var log = logging.MustGetLogger("test")

func DownloadFile(filePath string, url string) (err error) {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	// get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func Before() (err error) {
	err = DownloadFile("test.hdt", "http://srvgal80.deri.ie/~pabtor/test.hdt")
	if err != nil {
		return err
	}
	return nil
}

func After() (err error) {
	/*err = os.Remove("test.hdt")
	if err != nil {
		return err
	}*/
	return nil
}

func TestMain(m *testing.M) {
	err := Before()
	if err != nil {
		log.Error("Before function failed")
		log.Fatal(err)
	}
	retCode := m.Run()
	After()
	os.Exit(retCode)
}

/* TESTS */

func TestShouldIterateHDTFileGettingUrls(t *testing.T) {
	hdtMap := OpenHDT("./test.hdt")
	it := hdtMap.Search("", "", "")
	sl := it.Get(10)
	assert.Equal(t, len(sl), 10, "Get 10 elements from hdt")
}

func TestShouldIterate(t *testing.T) {
	hdtMap := OpenHDT("./test.hdt")
	it := hdtMap.Search("", "", "")
	sl := it.GetAll()
	assert.Equal(t, len(sl), 100656, "Get everything from hdt")
}
