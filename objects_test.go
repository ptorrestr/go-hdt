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
	hdtMap = OpenHDT("./test.hdt")
	return nil
}

func After() (err error) {
	/*err = os.Remove("test.hdt")
	if err != nil {
		return err
	}*/
	hdtMap.Free()
	return nil
}

var hdtMap HDTConnector

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
	it := hdtMap.Search("", "", "")
	sl := it.Get(10)
	assert.Equal(t, 10, len(sl))
	it.Free()
}

func TestShouldObtainStringFromAUriIterator(t *testing.T) {
	it := hdtMap.Search("", "", "")
	sl := it.Get(1)
	assert.IsType(t, string(""), sl[0].getSubject())
	assert.IsType(t, string(""), sl[0].getPredicate())
	assert.IsType(t, string(""), sl[0].getObject())
	assert.NotZero(t, len(sl[0].getSubject()))
	assert.NotZero(t, len(sl[0].getPredicate()))
	assert.NotZero(t, len(sl[0].getObject()))
	it.Free()
}

func TestShouldIterateHDTFileGettingIds(t *testing.T) {
	it := hdtMap.SearchID1("", "", "") //This is not working properly.
	sl := it.Get(10)
	assert.Equal(t, 10, len(sl))
	it.Free()
}

func TestShouldIterateHDTFileGettingIds2(t *testing.T) {
	it := hdtMap.SearchID2(0, 0, 0)
	sl := it.Get(10)
	assert.Equal(t, 10, len(sl))
	assert.IsType(t, uint(1), sl[0].getSubject())
	assert.IsType(t, uint(1), sl[0].getPredicate())
	assert.IsType(t, uint(1), sl[0].getObject())
	assert.NotZero(t, sl[0].getSubject())
	assert.NotZero(t, sl[0].getPredicate())
	assert.NotZero(t, sl[0].getObject())
	it.Free()
}

func TestShouldIterate(t *testing.T) {
	it := hdtMap.Search("", "", "")
	sl := it.GetAll()
	assert.Equal(t, len(sl), 100656, "Get everything from hdt")
	it.Free()
}
