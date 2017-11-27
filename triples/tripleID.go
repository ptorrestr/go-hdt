package triples

// #cgo pkg-config: hdt
// #include "tripleID.h"
import "C"

type TripleID struct {
	triple C.TripleID_
}

func Free(t TripleID) {
	C.tripleIdFree(t.triple)
}

func getIdSubject(t TripleID) uint {
	return uint(C.tripleIdGetSubject(t.triple))
}

func getIdObject(t TripleID) uint {
	return uint(C.tripleIdGetObject(t.triple))
}

func getIdPredicate(t TripleID) uint {
	return uint(C.tripleIdGetPredicate(t.triple))
}
