#ifndef CXX_TRIPLE_ID_ITERATOR_H
#define CXX_TRIPLE_ID_ITERATOR_H
#include <HDTManager.hpp>
#include "tripleID.hpp"
using namespace hdt;
using namespace std;

class CxxTripleIDIterator
{
private:
	IteratorTripleID *iter;
public:
	CxxTripleIDIterator(IteratorTripleID *iter);
	virtual ~CxxTripleIDIterator();

	CxxTripleID* next();
};
#endif
