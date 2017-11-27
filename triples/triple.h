#ifndef TRIPLE_H
#define TRIPLE_H

#ifdef __cplusplus
extern "C" {
#endif
	typedef void* Triple; 
	Triple tripleInit();
	const char* tripleGetSubject(Triple);
	const char* tripleGetObject(Triple);
	const char* tripleGetPredicate(Triple);
	void tripleFree(Triple);
#ifdef __cplusplus
}
#endif
#endif
