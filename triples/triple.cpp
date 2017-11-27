#include "triple.hpp"

CxxTriple::CxxTriple(TripleString *triple)
{
	this -> triple = new TripleString(triple -> getSubject(),
			triple -> getPredicate(),
			triple -> getObject());
}

CxxTriple::~CxxTriple()
{
	if ( triple )
	{
		delete triple;
		triple = NULL;
	}
}

string
CxxTriple::get_subject() const
{
	return triple -> getSubject();
}

string
CxxTriple::get_object() const
{
	return triple -> getObject();
}

string
CxxTriple::get_predicate() const
{
	return triple -> getPredicate();
}
