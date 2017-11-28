#ifndef CONNECTOR_HPP
#define CONNECTOR_HPP

#include <HDTManager.hpp>
#include "tripleIterator.hpp"
#include "tripleIDIterator.hpp"

using namespace hdt;
using namespace std;

class CxxConnector
{
private:
	HDT *hdt;
public:
	CxxConnector(const string& hdt_file);
	virtual ~CxxConnector();
	CxxTripleIterator* search(const string& uri1, const string& uri2, const string& uri3);
	CxxTripleIDIterator* search_id(const string& uri1, const string& uri2, const string& uri3);
	CxxTripleIDIterator* search_id(unsigned int id1, unsigned int id2, unsigned int id3);
	CxxTripleIDIterator* search_id(TripleID& triple);
	string id_to_uri(unsigned int id, const TripleComponentRole& role);
	unsigned int uri_to_id(const string& uri, const TripleComponentRole& role);
};
#endif
