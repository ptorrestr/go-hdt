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

func TestShouldIterateHDTFileGettingStrings(t *testing.T) {
	it := hdtMap.Search("", "", "")
	sl := it.Get(10)
	assert.Equal(t, 10, len(sl))
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
	assert.IsType(t, uint(1), sl[0].getSubject())
	assert.IsType(t, uint(1), sl[0].getPredicate())
	assert.IsType(t, uint(1), sl[0].getObject())
	assert.NotZero(t, sl[0].getSubject())
	assert.NotZero(t, sl[0].getPredicate())
	assert.NotZero(t, sl[0].getObject())
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

func TestShouldIterateTheWholeHDTFile(t *testing.T) {
	it := hdtMap.Search("", "", "")
	sl := it.GetAll()
	assert.Equal(t, len(sl), 100656, "Get everything from hdt")
	it.Free()
}

func TestShouldConsumeTheWholeIterator(t *testing.T) {
	it := hdtMap.Search("", "", "")
	triple := it.Get(1)[0]
	it2 := hdtMap.Search("", "", triple.getObject()) // Search everthing for this uri
	sl := it2.GetAll()
	assert.Equal(t, 1873, len(sl))
	it2.Free()
	it.Free()
}

func TestShouldNotIterateForNonExistingUri(t *testing.T) {
	it := hdtMap.Search("this://uri.does.not/exist", "", "")
	sl := it.GetAll()
	assert.Equal(t, 0, len(sl))
	it.Free()
}

func TestShouldNotIterateForNonExistingUri1(t *testing.T) {
	it := hdtMap.SearchID1("this://uri.does.not/exist", "", "")
	sl := it.GetAll()
	assert.Equal(t, 0, len(sl))
	it.Free()
}

func TestShouldNotIterateForNonExistingUri2(t *testing.T) {
	it := hdtMap.SearchID2(123123123, 0, 0)
	sl := it.GetAll()
	assert.Equal(t, 0, len(sl))
	it.Free()
}

func TestShouldTransformIdIntoUrl(t *testing.T) {
	it := hdtMap.SearchID1("", "", "")
	triple := it.Get(1)[0]
	s_uri := hdtMap.IdToUri(triple.getSubject(), subject)
	p_uri := hdtMap.IdToUri(triple.getPredicate(), predicate)
	o_uri := hdtMap.IdToUri(triple.getObject(), object)
	s_id := hdtMap.UriToId(s_uri, subject)
	p_id := hdtMap.UriToId(p_uri, predicate)
	o_id := hdtMap.UriToId(o_uri, object)
	assert.Equal(t, triple.getSubject(), s_id)
	assert.Equal(t, triple.getPredicate(), p_id)
	assert.Equal(t, triple.getObject(), o_id)
	it.Free()
}

func TestShouldTransformUrlIntoId(t *testing.T) {
	it := hdtMap.Search("", "", "")
	triple := it.Get(1)[0]
	s_id := hdtMap.UriToId(triple.getSubject(), subject)
	p_id := hdtMap.UriToId(triple.getPredicate(), predicate)
	o_id := hdtMap.UriToId(triple.getObject(), object)
	s_uri := hdtMap.IdToUri(s_id, subject)
	p_uri := hdtMap.IdToUri(p_id, predicate)
	o_uri := hdtMap.IdToUri(o_id, object)
	assert.Equal(t, triple.getSubject(), s_uri)
	assert.Equal(t, triple.getPredicate(), p_uri)
	assert.Equal(t, triple.getObject(), o_uri)
	it.Free()
}
