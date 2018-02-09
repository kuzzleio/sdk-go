#ifndef _COLLECTION_HPP_
#define _COLLECTION_HPP_

#include <iostream>
#include <list>
#include "core.hpp"

namespace kuzzleio {
    class Document;

    class Collection {
        public:
            collection* _collection;

            Collection(Kuzzle* kuzzle, const std::string& collection, const std::string& index);
            virtual ~Collection();
            int count(search_filters* filters, query_options* options=NULL) Kuz_Throw_KuzzleException;
            Collection* createDocument(Document* document, const std::string& id="", query_options* options=NULL) Kuz_Throw_KuzzleException;
            std::string deleteDocument(const std::string& id, query_options* options=NULL) Kuz_Throw_KuzzleException;
            Document* fetchDocument(const std::string& id, query_options* options=NULL) Kuz_Throw_KuzzleException;
            std::vector<Document*> mCreateDocument(std::vector<Document*>& documents, query_options* options=NULL) Kuz_Throw_KuzzleException;
            std::vector<Document*> mCreateOrReplaceDocument(std::vector<Document*>& documents, query_options* options=NULL) Kuz_Throw_KuzzleException;
            std::vector<std::string> mDeleteDocument(std::vector<std::string>& ids, query_options* options=NULL) Kuz_Throw_KuzzleException;
            std::vector<Document*> mGetDocument(std::vector<std::string>& ids, query_options* options=NULL) Kuz_Throw_KuzzleException;            
            std::vector<Document*> mReplaceDocument(std::vector<Document*>& documents, query_options* options=NULL) Kuz_Throw_KuzzleException;
            std::vector<Document*> mUpdateDocument(std::vector<Document*>& documents, query_options* options=NULL) Kuz_Throw_KuzzleException;
            bool publishMessage(Document* document, query_options* options=NULL) Kuz_Throw_KuzzleException;
            //bool publishMessage(json_object* content, query_options* options=NULL) Kuz_Throw_KuzzleException;
    };
}

#endif