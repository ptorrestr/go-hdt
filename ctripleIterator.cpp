#include "tripleIterator.hpp"
#include "tripleIterator.h"

_Triple tripleIteratorNext(_TripleIterator it)
{
	CxxTripleIterator *iter = (CxxTripleIterator*) it;
	CxxTriple *t = iter -> next();
	return (void*) t;
}

void tripleIteratorFree(_TripleIterator it)
{
	CxxTripleIterator *iter = (CxxTripleIterator*) it;
	delete iter;
}
