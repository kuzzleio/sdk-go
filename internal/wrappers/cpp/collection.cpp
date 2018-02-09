#include "collection.hpp"
#include "document.hpp"
#include "room.hpp"

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
      document **docs = new document *[documents.size()];
      int i = 0;
      for(auto const& doc: documents) {
        docs[i] = doc->_document;
        i++;
      }
      document_array_result *r = kuzzle_collection_m_create_document(_collection, docs, documents.size(), options);

      delete[] docs;

      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

    std::vector<Document*> Collection::mCreateOrReplaceDocument(std::vector<Document*>& documents, query_options* options) Kuz_Throw_KuzzleException {
      document **docs = new document *[documents.size()];
      int i = 0;
      for (auto const& doc : documents) {
        docs[i] = doc->_document;
        i++;
      }
      document_array_result *r = kuzzle_collection_m_create_or_replace_document(_collection, docs, documents.size(), options);

      delete[] docs;
      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

    std::vector<std::string> Collection::mDeleteDocument(std::vector<std::string>& ids, query_options* options) Kuz_Throw_KuzzleException {
      char **docsIds = new char *[ids.size()];
      int i = 0;
      for (auto const& id : ids) {
        docsIds[i] = const_cast<char*>(id.c_str());
        i++;
      }
      string_array_result *r = kuzzle_collection_m_delete_document(_collection, docsIds, ids.size(), options);

      delete[] docsIds;
      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<std::string> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(r->result[i]);

      delete(r);
      return v;
    }

    std::vector<Document*> Collection::mGetDocument(std::vector<std::string>& ids, query_options* options) Kuz_Throw_KuzzleException {
      char **docsIds = new char *[ids.size()];
      int i = 0;
      for (auto const& id : ids) {
        docsIds[i] = const_cast<char*>(id.c_str());
        i++;
      }
      document_array_result *r = kuzzle_collection_m_get_document(_collection, docsIds, ids.size(), options);

      delete[] docsIds;
      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

    std::vector<Document*> Collection::mReplaceDocument(std::vector<Document*>& documents, query_options* options) Kuz_Throw_KuzzleException {
      document **docs = new document *[documents.size()];
      int i = 0;
      for (auto const& doc : documents) {
        docs[i] = doc->_document;
        i++;
      }
      document_array_result *r = kuzzle_collection_m_replace_document(_collection, docs, documents.size(), options);

      delete[] docs;
      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

    std::vector<Document*> Collection::mUpdateDocument(std::vector<Document*>& documents, query_options* options) Kuz_Throw_KuzzleException {
      document **docs = new document *[documents.size()];
      int i = 0;
      for (auto const& doc : documents) {
        docs[i] = doc->_document;
        i++;
      }
      document_array_result *r = kuzzle_collection_m_update_document(_collection, docs, documents.size(), options);

      delete[] docs;
      if (r->error != NULL)
        throwExceptionFromStatus(r);

      std::vector<Document*> v;
      for (int i = 0; i < r->result_length; i++)
        v.push_back(new Document(this, (r->result + i)->id, (r->result + i)->content));

      delete(r);
      return v;
    }

    bool Collection::publishMessage(json_object* content, query_options* options) Kuz_Throw_KuzzleException {
      bool_result *r = kuzzle_collection_publish_message(_collection, content, options);
      if (r->error != NULL)
          throwExceptionFromStatus(r);
      bool ret = r->result;
      delete(r);
      return ret;
    }

    Document* Collection::replaceDocument(const std::string& id, Document* document, query_options* options) Kuz_Throw_KuzzleException {
      document_result* r = kuzzle_collection_update_document(_collection, const_cast<char*>(id.c_str()), document->_document, options);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      Document* ret = new Document(this, r->result->id, r->result->content);
      delete(r);
      return ret;
    }

    search_result* Collection::scroll(const std::string& id, query_options* options) Kuz_Throw_KuzzleException {
      search_result* r = kuzzle_collection_scroll(_collection, const_cast<char*>(id.c_str()), options);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      return r;
    }

    search_result* Collection::search(search_filters* filters, query_options* options) Kuz_Throw_KuzzleException {
      search_result* r = kuzzle_collection_search(_collection, filters, options);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      return r;
    }

    void call_collection_cb(notification_result* res, void* data) {
      ((Collection*)data)->getListener()->onMessage(res);
    }

    NotificationListener* Collection::getListener() {
      return _listener_instance;
    }

    Room* Collection::subscribe(search_filters* filters, NotificationListener *listener, room_options* options) Kuz_Throw_KuzzleException {
      room_result* r = kuzzle_collection_subscribe(_collection, filters, options, &call_collection_cb, this);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      _listener_instance = listener;
      
      Room* ret = new Room(r->result, NULL, listener);
      free(r);
      return ret;
    }

    Document* Collection::updateDocument(const std::string& id, Document* document, query_options* options) Kuz_Throw_KuzzleException {
      document_result* r = kuzzle_collection_update_document(_collection, const_cast<char*>(id.c_str()), document->_document, options);
      if (r->error != NULL)
        throwExceptionFromStatus(r);
      Document* ret = new Document(this, r->result->id, r->result->content);
      delete(r);
      return ret;
    }

}