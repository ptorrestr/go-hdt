#ifndef CXX_ITERATOR_H
#define CXX_ITERATOR_H
#include <HDTManager.hpp>
#include "triple.hpp"
using namespace hdt;
using namespace std;

class CxxTripleIterator
{
private:
	IteratorTripleString *iter;
public:
	CxxTripleIterator(IteratorTripleString *iter);
	virtual ~CxxTripleIterator();

	CxxTriple* next();
};
#endif
