#include "kuzzle.hpp"
#include "collection.hpp"

namespace kuzzleio {
    Collection::Collection(Kuzzle* kuzzle) {
        _collection = new collection();
        kuzzle_new_collection(_collection, kuzzle->_kuzzle);
    }

    Collection::~Collection() {
        unregisterCollection(_collection);
        kuzzle_free_collection(_collection);
    }

    void Collection::create(const std::string& index, const std::string& collection) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_collection_create(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        kuzzle_free_void_result(r);
    }

    bool Collection::exists(const std::string& index, const std::string& collection) Kuz_Throw_KuzzleException {
        bool_result *r = kuzzle_collection_exists(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        bool ret = r->result;
        kuzzle_free_bool_result(r);
        return ret;
    }

    std::string Collection::list(const std::string& index, collection_list_options *collectionListOptions) Kuz_Throw_KuzzleException {
        string_result *r = kuzzle_collection_list(_collection, const_cast<char*>(index.c_str()), collectionListOptions);
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    void Collection::truncate(const std::string& index, const std::string& collection) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_collection_truncate(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        kuzzle_free_void_result(r);
    }

    std::string Collection::getMapping(const std::string& index, const std::string& collection) Kuz_Throw_KuzzleException {
        string_result *r = kuzzle_collection_get_mapping(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);

        return ret;
    }

    void Collection::updateMapping(const std::string& index, const std::string& collection, const std::string& body) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_collection_update_mapping(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        kuzzle_free_void_result(r);
    }

    std::string Collection::getSpecifications(const std::string& index, const std::string& collection) Kuz_Throw_KuzzleException {
        string_result *r = kuzzle_collection_get_specifications(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);

        return ret;
    }

    search_result* Collection::searchSpecifications(search_options *searchOptions) Kuz_Throw_KuzzleException {
        search_result *r = kuzzle_collection_search_specifications(_collection, searchOptions);
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        search_result *ret = r;
        kuzzle_free_search_result(r);

        return ret;
    }

    std::string Collection::updateSpecifications(const std::string& index, const std::string& collection, const std::string& body) Kuz_Throw_KuzzleException {
        string_result *r = kuzzle_collection_update_specifications(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        std::string ret = r->result;
        kuzzle_free_string_result(r);

        return ret;
    }

    void Collection::validateSpecifications(const std::string& body) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_collection_validate_specifications(_collection, const_cast<char*>(body.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        kuzzle_free_void_result(r);
    }

    void Collection::deleteSpecifications(const std::string& index, const std::string& collection) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_collection_delete_specifications(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);

        kuzzle_free_void_result(r);
    }
}
