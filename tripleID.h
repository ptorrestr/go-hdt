#ifndef TRIPLE_ID_H
#define TRIPLE_ID_H
#ifdef __cplusplus
extern "C" {
#endif
	typedef void* _TripleID;
	unsigned int tripleIdGetSubject(_TripleID);
	unsigned int tripleIdGetObject(_TripleID);
	unsigned int tripleIdGetPredicate(_TripleID);
	void tripleIdFree(_TripleID);
#ifdef __cplusplus
}
#endif
#endif
