#ifndef TRIPLE_H
#define TRIPLE_H
#ifdef __cplusplus
extern "C" {
#endif
	typedef void* _Triple; 
	const char* tripleGetSubject(_Triple);
	const char* tripleGetObject(_Triple);
	const char* tripleGetPredicate(_Triple);
	void tripleFree(_Triple);
#ifdef __cplusplus
}
#endif
#endif
