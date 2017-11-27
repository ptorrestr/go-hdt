#include "tripleIterator.hpp"

CxxTripleIterator::CxxTripleIterator(IteratorTripleString *iter) : iter(iter)
{
}

CxxTripleIterator::~CxxTripleIterator()
{
	if ( iter) 
	{
		delete iter;
		iter = NULL;
	}
}

CxxTriple*
CxxTripleIterator::next()
{
	if ( iter -> hasNext() )
	{
		return new CxxTriple( iter -> next() );
	}
	return nullptr;
}
