#include <HDTEnums.hpp>
#include "triplerole.h"

using namespace hdt;

Triplerole initializeTriplerole() {
	Triplerole *ret = new Triplerole();
	ret -> subject = SUBJECT;
	ret -> predicate = OBJECT;
	ret -> object = OBJECT;
	return *ret;
}

