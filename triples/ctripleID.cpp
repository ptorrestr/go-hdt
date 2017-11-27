#include "tripleID.hpp"
#include "tripleID.h"

unsigned int tripleIdGetSubject(_TripleID t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	return tt -> get_subject();
}

unsigned int tripleIdGetObject(_TripleID t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	return tt -> get_object();
}

unsigned int tripleIdGetPredicate(_TripleID t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	return tt -> get_predicate();
}

void tripleIdFree(_TripleID t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	delete tt;
}
