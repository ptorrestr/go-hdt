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

func getSubject(t Triple) string {
	return C.GoString(C.tripleGetSubject(t.triple))
}

func getObject(t Triple) string {
	return C.GoString(C.tripleGetObject(t.triple))
}

func getPredicate(t Triple) string {
	return C.GoString(C.tripleGetPredicate(t.triple))
}
