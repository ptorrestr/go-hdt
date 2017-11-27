#ifndef TRIPLE_ID_H
#define TRIPLE_ID_H
#ifdef __cplusplus
extern "C" {
#endif
	typedef void* TripleID_; //TripleID already exists on hdt
	unsigned int tripleIdGetSubject(TripleID_);
	unsigned int tripleIdGetObject(TripleID_);
	unsigned int tripleIdGetPredicate(TripleID_);
	void tripleIdFree(TripleID_);
#ifdef __cplusplus
}
#endif
#endif
