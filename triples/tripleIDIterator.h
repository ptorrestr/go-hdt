#ifndef TRIPLE_ID_ITERATOR_H
#define TRIPLE_ID_ITERATOR_H
#include "tripleID.h"
#ifdef __cplusplus
extern "C" {
#endif
	typedef void* _TripleIDIterator;
	_TripleID tripleIDIteratorNext(_TripleIDIterator);
	void tripleIDIteratorFree(_TripleIDIterator);
#ifdef __cplusplus
}
#endif
#endif
