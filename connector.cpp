#include "connector.hpp"

CxxConnector::CxxConnector(const string& hdt_file) : hdt(NULL)
{
	try {
		hdt = HDTManager::mapIndexedHDT(hdt_file.c_str());
	}
	catch (...)
	{
		cerr << "Error opening HDT file" << endl;
	}
}

CxxConnector::~CxxConnector()
{
	if ( hdt )
	{
		delete hdt;
		hdt = NULL;
	}
}

CxxTripleIterator*
CxxConnector::search(const string& uri1, const string& uri2, const string& uri3)
{
	IteratorTripleString *iter = hdt -> search(uri1.c_str(), uri2.c_str(), uri3.c_str());
	return new CxxTripleIterator(iter);
}

CxxTripleIDIterator*
CxxConnector::search_id(const string& uri1, const string& uri2, const string& uri3)
{
	TripleString ts(uri1, uri2, uri3);
	TripleID tid;
	hdt -> getDictionary() -> tripleStringtoTripleID(ts, tid);
	if ( ( tid.getSubject() == 0 && !uri1.empty() ) ||
			 ( tid.getPredicate() == 0 && !uri2.empty() ) ||
			 ( tid.getObject() == 0 && !uri3.empty() ) )
	{
		// If couldn't found the uris, return empty iterator
		return new CxxTripleIDIterator(new IteratorTripleID() );
	}
	return search_id(tid);
}

CxxTripleIDIterator*
CxxConnector::search_id(unsigned int id1, unsigned int id2, unsigned int id3)
{
	TripleID tid(id1, id2, id3);
	return search_id(tid);
}

CxxTripleIDIterator*
CxxConnector::search_id(TripleID& triple)
{
	// search should receive a const tripleID&
	return new CxxTripleIDIterator( hdt -> getTriples() -> search(triple) );
}

string
CxxConnector::id_to_uri(unsigned int id, const TripleComponentRole& role)
{
	return hdt -> getDictionary() -> idToString(id, role);
}

unsigned int
CxxConnector::uri_to_id(const string& uri, const TripleComponentRole& role)
{
	return hdt -> getDictionary() -> stringToId(uri, role);
}
