#ifndef TRIPLE_ITERATOR_H
#define TRIPLE_ITERATOR_H
#include "triple.h"
#ifdef __cplusplus
extern "C" {
#endif
	typedef void* _TripleIterator;
	_Triple tripleIteratorNext(_TripleIterator);
	void tripleIteratorFree(_TripleIterator);
#ifdef __cplusplus
}
#endif
#endif
