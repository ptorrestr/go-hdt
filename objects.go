package hdt

// #cgo pkg-config: hdt
// #cgo CFLAGS: -I${SRCDIR}/include
// #include "triple.h"
// #include "tripleID.h"
// #include "tripleIterator.h"
// #include "tripleIDIterator.h"
// #include "connector.h"
import "C"
import "fmt"

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

func (t Triple) String() string {
	return fmt.Sprintf("{%s, %s, %s}", t.getSubject(), t.getPredicate(), t.getObject())
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

func (t TripleID) String() string {
	return fmt.Sprintf("{%b, %b, %b}", t.getSubject(), t.getPredicate(), t.getObject())
}

/* Define triple iterator*/
type TripleIterator struct {
	iter C._TripleIterator
}

func (ti TripleIterator) Free() {
	C.tripleIteratorFree(ti.iter)
}

func (ti TripleIterator) Get(n uint) []Triple {
	var ret []Triple
	ret = make([]Triple, n) //if cap is ommited, the default is "n"
	var i uint
	for i = 0; i < n; i = i + 1 {
		var i_tri Triple
		i_tri.triple = C.tripleIteratorNext(ti.iter)
		ret[i] = i_tri
	}
	return ret
}

func (ti TripleIterator) GetAll() []Triple {
	var ret []Triple
	for true {
		var i_tri Triple
		i_tri.triple = C.tripleIteratorNext(ti.iter)
		if i_tri.triple == nil {
			break
		}
		ret = append(ret, i_tri)
	}
	return ret
}

/* Define triple ID iterator*/
type TripleIDIterator struct {
	iter C._TripleIDIterator
}

func (ti TripleIDIterator) Free() {
	C.tripleIDIteratorFree(ti.iter)
}

func (ti TripleIDIterator) Get(n uint) []TripleID {
	var ret []TripleID
	ret = make([]TripleID, n)
	var i uint
	for i = 0; i < n; i = i + 1 {
		var i_tri TripleID
		i_tri.triple = C.tripleIDIteratorNext(ti.iter)
		ret[i] = i_tri
	}
	return ret
}

func (ti TripleIDIterator) GetAll() []TripleID {
	var ret []TripleID
	for true {
		var i_tri TripleID
		i_tri.triple = C.tripleIDIteratorNext(ti.iter)
		if i_tri.triple == nil {
			break
		}
		ret = append(ret, i_tri)
	}
	return ret
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

func (c HDTConnector) NeighboursOut(id uint) []uint {
	it := c.SearchID2(id, 0, 0)
	var ret []uint
	for true {
		var i_tri TripleID
		i_tri.triple = C.tripleIDIteratorNext(it.iter)
		if i_tri.triple == nil {
			break
		}
		ret = append(ret, i_tri.getObject())
		i_tri.Free()
	}
	it.Free()
	return ret
}

func (c HDTConnector) NeighboursIn(id uint) []uint {
	it := c.SearchID2(0, 0, id)
	var ret []uint
	for true {
		var i_tri TripleID
		i_tri.triple = C.tripleIDIteratorNext(it.iter)
		if i_tri.triple == nil {
			break
		}
		ret = append(ret, i_tri.getSubject())
		i_tri.Free()
	}
	it.Free()
	return ret
}

func (c HDTConnector) Neighbours(id uint) []uint {
	in := c.NeighboursIn(id)
	out := c.NeighboursOut(id)
	return append(in, out...)
}
