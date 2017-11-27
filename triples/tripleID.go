package triples

// #cgo pkg-config: hdt
// #include "tripleID.h"
import "C"

type TripleID struct {
	triple C.TripleID_
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
