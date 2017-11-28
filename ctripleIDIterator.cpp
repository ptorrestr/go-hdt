#include "tripleIDIterator.hpp"
#include "tripleIDIterator.h"

_TripleID tripleIDIteratorNext(_TripleIDIterator it)
{
	CxxTripleIDIterator *iter = (CxxTripleIDIterator*) it;
	CxxTripleID *t = iter -> next();
	return (void*) t;
}

void tripleIDIteratorFree(_TripleIDIterator it)
{
	CxxTripleIDIterator *iter = (CxxTripleIDIterator*) it;
	delete iter;
}
