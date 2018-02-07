#include <iostream>
#include <vector>
#include "collection.hpp"
#include "document.hpp"

namespace kuzzleio {
    Collection::Collection(Kuzzle* kuzzle, const std::string& col, const std::string& index) {
        _collection = new collection();
        kuzzle_new_collection(_collection, kuzzle->_kuzzle, const_cast<char*>(col.c_str()), const_cast<char*>(index.c_str()));
    }

    Collection::~Collection() {
        unregisterCollection(_collection);
        delete(_collection);
    }

    int Collection::count(search_filters* filters, query_options* options) Kuz_Throw_KuzzleException {
      int_result *r = kuzzle_collection_count(_collection, filters, options);
      if (r->error != NULL)
          throwExceptionFromStatus(r);
      int ret = r->result;
      delete(r);
      return ret;
    }

    Collection* Collection::createDocument(Document* document, const std::string& id, query_options* options) Kuz_Throw_KuzzleException {
      document_result *r = kuzzle_collection_create_document(_collection, const_cast<char*>(id.c_str()), document->_document, options);
      if (r->error != NULL)
          throwExceptionFromStatus(r);
      delete(r);
      return this;
    }

    std::string Collection::deleteDocument(const std::string& id, query_options* options) Kuz_Throw_KuzzleException {
      string_result *r = kuzzle_collection_delete_document(_collection, const_cast<char*>(id.c_str()), options);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      std::string ret = r->result;
      delete(r);
      return ret;
    }

    Document* Collection::fetchDocument(const std::string& id, query_options* options) Kuz_Throw_KuzzleException {
      document_result *r = kuzzle_collection_fetch_document(_collection, const_cast<char*>(id.c_str()), options);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      Document* ret = new Document(this, r->result->id, r->result->content);
      delete(r);
      return ret;
    }
}