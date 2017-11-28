#include "connector.hpp"
#include "connector.h"

_Connector connectorInit(const char* hdt_file)
{
	CxxConnector *conn = new CxxConnector(	string(hdt_file) );
	return (void*) conn;
}

_TripleIterator connectorSearch(_Connector c, const char *uri1, const char* uri2, const char* uri3)
{
	CxxConnector *conn = (CxxConnector*) c;
	CxxTripleIterator *iter = conn ->	search(string(uri1), string(uri2), string(uri3));
	return (void*) iter;
}

_TripleIDIterator connectorSearchId1(_Connector c, const char* uri1, const char* uri2, const char* uri3)
{
	CxxConnector *conn = (CxxConnector*) c;
	CxxTripleIDIterator *iter = conn -> search_id(string(uri1), string(uri2), string(uri3));
	return (void*) iter;
}

_TripleIDIterator connectorSearchId2(_Connector c, unsigned int id1, unsigned int id2, unsigned int id3)
{
	CxxConnector *conn = (CxxConnector*) c;
	CxxTripleIDIterator *iter = conn -> search_id(id1, id2, id3);
	return (void*) iter;
}

TripleComponentRole typeToRole(const unsigned int type) 
{
	TripleComponentRole role;
	switch( type ) {
		case 0:
			role = SUBJECT;
			break;
		case 1:
			role = PREDICATE;
			break;
		case 2:
			role = OBJECT;
			break;
		default:
			role = SUBJECT;
	}
	return role;
}

const char* connectorIdToUri(_Connector c, unsigned int id, const unsigned int type)
{
	CxxConnector *conn = (CxxConnector*) c;
	return (conn -> id_to_uri(id, typeToRole(type))).c_str();
}

unsigned int connectorUriToId(_Connector c, const char* uri, const unsigned int type)
{
	CxxConnector *conn = (CxxConnector*) c;
	return conn -> uri_to_id(uri, typeToRole(type));
}

void connectorFree(_Connector c) 
{
	CxxConnector *conn = (CxxConnector*) c;
	delete conn;
}
