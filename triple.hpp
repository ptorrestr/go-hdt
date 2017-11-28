#ifndef CXX_TRIPLE_HPP
#define CXX_TRIPLE_HPP
#include <HDTManager.hpp>
using namespace hdt;

class CxxTriple {
private:
	TripleString *triple;
public:
	CxxTriple(TripleString *triple);
	virtual ~CxxTriple();

	string get_subject() const;
	string get_object() const;
	string get_predicate() const;
};
#endif
