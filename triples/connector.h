#ifndef CONNECTOR_H
#define CONNECTOR_H
#include "tripleIterator.h"
#include "tripleIDIterator.h"
#ifdef __cplusplus
extern "C" {
#endif
	typedef void* _Connector;
	_Connector connectorInit(const char* hdt_file);
	_TripleIterator connectorSearch(_Connector, const char* uri1, const char* uri2, const char* uri3);
	_TripleIDIterator connectorSearchId1(_Connector, const char* uri1, const char* uri2, const char* uri3);
	_TripleIDIterator connectorSearchId2(_Connector, unsigned int id1, unsigned int id2, unsigned int id3);
	const char* connectorIdToUri(_Connector, unsigned int id, const unsigned int type);
	unsigned int connectorUriToId(_Connector, const char* uri, const unsigned int type);
	void connectorFree(_Connector);
#ifdef __cplusplus
}
#endif
#endif
