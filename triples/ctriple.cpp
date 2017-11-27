#include "triple.hpp"
#include "triple.h"

_Triple tripleInt(_TripleString ts)
{
	TripleString* ts2 = (TripleString*)ts;
	CxxTriple *t = new CxxTriple(ts2);
	return (void*)t;
}

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
