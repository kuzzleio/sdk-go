#include "document.hpp"
#include "collection.hpp"

namespace kuzzleio {
    
    Document::Document(Collection *collection, const std::string& id, json_object* content) Kuz_Throw_KuzzleException {
        _document = new document();
        _collection = collection;
        kuzzle_new_document(_document, collection->_collection, const_cast<char*>(id.c_str()), content); 
    }

    Document::Document(Document& document) {
      _document = document._document;
      _collection = document._collection;
      kuzzle_new_document(_document, _collection->_collection, _document->id, _document->content); 
    }

    Document::~Document() {
        unregisterDocument(_document);
        delete(_document);
    }

    std::string Document::delete_(query_options* options) Kuz_Throw_KuzzleException {
        string_result *r = kuzzle_document_delete(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        std::string ret = r->result;
        delete(r);
        return ret;
    }

    bool Document::exists(query_options* options) Kuz_Throw_KuzzleException {
        bool_result *r = kuzzle_document_exists(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        bool ret = r->result;
        delete(r);
        return ret;
    }

    bool Document::publish(query_options* options) Kuz_Throw_KuzzleException {
        bool_result *r = kuzzle_document_publish(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        bool ret = r->result;
        delete(r);
        return ret;
    }

    Document* Document::refresh(query_options* options) Kuz_Throw_KuzzleException {
        document_result *r = kuzzle_document_refresh(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        Document* ret = new Document(_collection, r->result->id, r->result->content);
        delete(r);
        return ret;
    }

    Document* Document::create(query_options* options) Kuz_Throw_KuzzleException {
        document_result *r = kuzzle_document_create(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        Document* ret = new Document(_collection, r->result->id, r->result->content);

        _document->id = r->result->id;
        _document->version = r->result->version;
        delete(r);
        return ret;
    }

    Document* Document::replace(query_options* options) Kuz_Throw_KuzzleException {
        document_result *r = kuzzle_document_replace(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        Document* ret = new Document(_collection, r->result->id, r->result->content);

        _document->id = r->result->id;
        _document->version = r->result->version;
        delete(r);
        return ret;
    }

    Document* Document::update(query_options* options) Kuz_Throw_KuzzleException {
        document_result *r = kuzzle_document_update(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        Document* ret = new Document(_collection, r->result->id, r->result->content);

        _document->id = r->result->id;
        _document->version = r->result->version;
        delete(r);
        return ret;
    }

    Document* Document::setContent(json_object* content, bool replace) {
        kuzzle_document_set_content(_document, content, replace);
        return this;
    }

    json_object* Document::getContent() {
        return kuzzle_document_get_content(_document);
    }

    void call_cb(notification_result* res, void* data) {
        ((Document*)data)->getListener()->onMessage(res);
    }

    Room* Document::subscribe(NotificationListener* listener, room_options* options) {
        room_result* r = kuzzle_document_subscribe(_document, options, &call_cb, this);
        if (r->error != NULL)
          throwExceptionFromStatus(r);
        _listener_instance = listener;
        
        Room* ret = new Room(r->result, NULL, listener);
        free(r);
        return ret;
    }

    NotificationListener* Document::getListener() {
        return _listener_instance;
    }
}
