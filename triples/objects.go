package triples

// #cgo pkg-config: hdt
// #include "triple.h"
// #include "tripleID.h"
// #include "tripleIterator.h"
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
