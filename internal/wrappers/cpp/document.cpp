#include "document.hpp"

namespace kuzzleio {
    
    Document::Document(Collection *collection, const std::string& id, json_object* content) Kuz_Throw_KuzzleException {
        _document = new document();
        _collection = collection;
        kuzzle_new_document(_document, collection->_collection, const_cast<char*>(id.c_str()), content);
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

    Document* Document::save(query_options* options) Kuz_Throw_KuzzleException {
        document_result *r = kuzzle_document_save(_document, options);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        //Document* ret = new Document(_collection, r->result->id, r->result->content);
        delete(r);
        return this;
    }

    Document* Document::setContent(json_object* content, bool replace) {
        kuzzle_document_set_content(_document, content, replace);
        return this;
    }
}
