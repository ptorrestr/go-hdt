#include "tripleID.hpp"

CxxTripleID::CxxTripleID(TripleID *triple)
{
	this -> triple = new TripleID(triple -> getSubject(),
			triple -> getPredicate(),
			triple -> getObject());
}

CxxTripleID::~CxxTripleID()
{
	if ( triple )
	{
		delete triple;
		triple = NULL;
	}
}

unsigned int
CxxTripleID::get_subject() const
{
	return triple -> getSubject();
}

unsigned int
CxxTripleID::get_predicate() const
{
	return triple -> getPredicate();
}

unsigned int
CxxTripleID::get_object() const
{
	return triple -> getObject();
}
