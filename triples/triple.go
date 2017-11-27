package triples

// #cgo pkg-config: hdt
// #include "triple.h"
import "C"

type Triple struct {
	triple C.Triple
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
