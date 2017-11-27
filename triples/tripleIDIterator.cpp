#include "tripleIDIterator.hpp"

CxxTripleIDIterator::CxxTripleIDIterator(IteratorTripleID *iter) : iter(iter)
{
}

CxxTripleIDIterator::~CxxTripleIDIterator()
{
	if ( iter )
	{
		delete iter;
		iter = NULL;
	}
}

CxxTripleID*
CxxTripleIDIterator::next()
{
	if ( iter -> hasNext() )
	{
		return new CxxTripleID(iter -> next() );
	}
	return nullptr;
}
