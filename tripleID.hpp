#ifndef CXX_TRIPLE_ID_HPP
#define CXX_TRIPLE_ID_HPP
#include <HDTManager.hpp>
using namespace hdt;

class CxxTripleID {
private:
	TripleID *triple;
public:
	CxxTripleID(TripleID *triple);
	virtual ~CxxTripleID();

	unsigned int get_subject() const;
	unsigned int get_predicate() const;
	unsigned int get_object() const;
};
#endif
