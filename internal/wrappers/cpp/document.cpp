#include "document.hpp"

namespace kuzzleio {
    
    Document::Document(Collection *collection, const std::string& id, json_object* content) Kuz_Throw_KuzzleException {
        _document = new document();
        kuzzle_new_document(_document, collection->_collection, const_cast<char*>(id.c_str()), content);
        _document->_collection = collection->_collection;
        _document->id = const_cast<char*>(id.c_str());
        this->id = id;
        _document->content = content;
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
}
