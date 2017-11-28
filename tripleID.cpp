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

	if ( triple ) 
	{
		return triple -> getSubject();
	}
	return 0;
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
