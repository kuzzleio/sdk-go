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

    std::vector<Document*> Collection::mCreateDocument(std::vector<Document*>& documents, query_options* options) Kuz_Throw_KuzzleException {
      document **docs = (document**)calloc(1, sizeof(*docs) * documents.size());
      int i = 0;
      for(auto const& doc: documents) {
        docs[i] = doc->_document;
        i++;
      }
      document_array_result *r = kuzzle_collection_m_create_document(_collection, docs, documents.size(), options);

      for (int j = 0; j < documents.size(); j++)
        free(docs[j]);

      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

    std::vector<Document*> Collection::mCreateOrReplaceDocument(std::vector<Document*>& documents, query_options* options) Kuz_Throw_KuzzleException {
      document **docs = (document**)calloc(1, sizeof(*docs) * documents.size());
      int i = 0;
      for(auto const& doc: documents) {
        docs[i] = doc->_document;
        i++;
      }
      document_array_result *r = kuzzle_collection_m_create_or_replace_document(_collection, docs, documents.size(), options);

      for (int j = 0; j < documents.size(); j++)
        free(docs[j]);

      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

}