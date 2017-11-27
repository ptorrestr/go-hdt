#include "tripleID.hpp"
#include "tripleID.h"

unsigned int tripleIdGetSubject(TripleID_ t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	return tt -> get_subject();
}

unsigned int tripleIdGetObject(TripleID_ t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	return tt -> get_object();
}

unsigned int tripleIdGetPredicate(TripleID_ t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	return tt -> get_predicate();
}

void tripleIdFree(TripleID_ t) {
	CxxTripleID *tt = (CxxTripleID*)t;
	delete tt;
}
