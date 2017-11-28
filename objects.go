package hdt

// #cgo pkg-config: hdt
// #include "triple.h"
// #include "tripleID.h"
// #include "tripleIterator.h"
// #include "tripleIDIterator.h"
// #include "connector.h"
import "C"

/* Define Triple*/
type Triple struct {
	triple C._Triple
}

func (t Triple) Free() {
	C.tripleFree(t.triple)
}

func (t Triple) getSubject() string {
	return C.GoString(C.tripleGetSubject(t.triple))
}

func (t Triple) getObject() string {
	return C.GoString(C.tripleGetObject(t.triple))
}

func (t Triple) getPredicate() string {
	return C.GoString(C.tripleGetPredicate(t.triple))
}

/* define triple ID */
type TripleID struct {
	triple C._TripleID
}

func (t TripleID) Free() {
	C.tripleIdFree(t.triple)
}

func (t TripleID) getSubject() uint {
	return uint(C.tripleIdGetSubject(t.triple))
}

func (t TripleID) getObject() uint {
	return uint(C.tripleIdGetObject(t.triple))
}

func (t TripleID) getPredicate() uint {
	return uint(C.tripleIdGetPredicate(t.triple))
}

/* Define triple iterator*/
type TripleIterator struct {
	iter C._TripleIterator
}

func (ti TripleIterator) Free() {
	C.tripleIteratorFree(ti.iter)
}

func (ti TripleIterator) Next() Triple {
	var ret Triple
	ret.triple = C.tripleIteratorNext(ti.iter)
	return ret
}

/* Define triple ID iterator*/
type TripleIDIterator struct {
	iter C._TripleIDIterator
}

func (ti TripleIDIterator) Free() {
	C.tripleIDIteratorFree(ti.iter)
}

func (ti TripleIDIterator) Next() TripleID {
	var ret TripleID
	ret.triple = C.tripleIDIteratorNext(ti.iter)
	return ret
}

/* Define triple role */
type TripleRole uint

const (
	subject TripleRole = iota
	predicate
	object
)

/* Define connector*/
type HDTConnector struct {
	conn C._Connector
}

func OpenHDT(file string) HDTConnector {
	var ret HDTConnector
	ret.conn = C.connectorInit(C.CString(file))
	return ret
}

func (c HDTConnector) Free() {
	C.connectorFree(c.conn)
}

func (c HDTConnector) Search(uri1 string, uri2 string, uri3 string) TripleIterator {
	var ret TripleIterator
	ret.iter = C.connectorSearch(c.conn, C.CString(uri1), C.CString(uri2), C.CString(uri3))
	return ret
}

func (c HDTConnector) SearchID1(uri1 string, uri2 string, uri3 string) TripleIDIterator {
	var ret TripleIDIterator
	ret.iter = C.connectorSearchId1(c.conn, C.CString(uri1), C.CString(uri2), C.CString(uri3))
	return ret
}

func (c HDTConnector) SearchID2(id1 uint, id2 uint, id3 uint) TripleIDIterator {
	var ret TripleIDIterator
	ret.iter = C.connectorSearchId2(c.conn, C.uint(id1), C.uint(id2), C.uint(id3))
	return ret
}

func (c HDTConnector) IdToUri(id uint, t TripleRole) string {
	return C.GoString(C.connectorIdToUri(c.conn, C.uint(id), C.uint(t)))
}

func (c HDTConnector) UriToId(uri string, t TripleRole) uint {
	return uint(C.connectorUriToId(c.conn, C.CString(uri), C.uint(t)))
}
