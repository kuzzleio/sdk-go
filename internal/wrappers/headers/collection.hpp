#ifndef _COLLECTION_HPP_
#define _COLLECTION_HPP_

#include <iostream>
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
    };
}

#endif