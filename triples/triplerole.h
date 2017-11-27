#ifndef TRIPLEROLE_H
#define TRIPLEROLE_H

#ifdef __cplusplus
extern "C" {
#endif
	typedef struct Triplerole {
		int subject;
		int predicate;
		int object;
	} Triplerole;
	Triplerole initializeTriplerole();

#ifdef __cplusplus
}
#endif

#endif
