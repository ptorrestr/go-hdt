#include "triple.hpp"
#include "triple.h"

const char* tripleGetSubject(_Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	return tt -> get_subject().c_str();
}

const char* tripleGetObject(_Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	return tt -> get_object().c_str();
}

const char* tripleGetPredicate(_Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	return tt -> get_predicate().c_str();
}

void tripleFree(_Triple t) {
	CxxTriple *tt = (CxxTriple*)t;
	delete tt;
}
