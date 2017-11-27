#include "triple.hpp"
#include "triple.h"

const char* tripleGetSubject(Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	return tt -> get_subject().c_str();
}

const char* tripleGetObject(Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	return tt -> get_object().c_str();
}

const char* tripleGetPredicate(Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	return tt -> get_predicate().c_str();
}

void tripleFree(Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	delete tt;
}
