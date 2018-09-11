// Copyright 2015-2018 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#include "kuzzle.hpp"
#include "collection.hpp"

namespace kuzzleio {
    Collection::Collection(Kuzzle* kuzzle) {
        _collection = new collection();
        kuzzle_new_collection(_collection, kuzzle->_kuzzle);
    }

    Collection::Collection(Kuzzle* kuzzle, collection *collection) {
        _collection = collection;
        kuzzle_new_collection(collection, kuzzle->_kuzzle);
    }

    Collection::~Collection() {
        unregisterCollection(_collection);
        delete(_collection);
    }

    void Collection::create(const std::string& index, const std::string& collection, const std::string* body, query_options *options) {
        error_result *r = kuzzle_collection_create(
            _collection,
            const_cast<char*>(index.c_str()),
            const_cast<char*>(collection.c_str()),
            body != nullptr ? const_cast<char*>(body->c_str()) : nullptr,
            options);

        if (r != nullptr)
            throwExceptionFromStatus(r);

        kuzzle_free_error_result(r);
    }

    bool Collection::exists(const std::string& index, const std::string& collection, query_options *options) {
        bool_result *r = kuzzle_collection_exists(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        bool ret = r->result;
        kuzzle_free_bool_result(r);
        return ret;
    }

    std::string Collection::list(const std::string& index, query_options *options) {
        string_result *r = kuzzle_collection_list(_collection, const_cast<char*>(index.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    void Collection::truncate(const std::string& index, const std::string& collection, query_options *options) {
        error_result *r = kuzzle_collection_truncate(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), options);
        if (r != nullptr)
            throwExceptionFromStatus(r);

        kuzzle_free_error_result(r);
    }

    std::string Collection::getMapping(const std::string& index, const std::string& collection, query_options *options) {
        string_result *r = kuzzle_collection_get_mapping(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);

        return ret;
    }

    void Collection::updateMapping(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        error_result *r = kuzzle_collection_update_mapping(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r != nullptr)
            throwExceptionFromStatus(r);

        kuzzle_free_error_result(r);
    }

    std::string Collection::getSpecifications(const std::string& index, const std::string& collection, query_options *options) {
        string_result *r = kuzzle_collection_get_specifications(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);

        return ret;
    }

    search_result* Collection::searchSpecifications(query_options *options) {
        search_result *r = kuzzle_collection_search_specifications(_collection, options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        search_result *ret = r;
        kuzzle_free_search_result(r);

        return ret;
    }

    std::string Collection::updateSpecifications(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        string_result *r = kuzzle_collection_update_specifications(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);
        std::string ret = r->result;
        kuzzle_free_string_result(r);

        return ret;
    }

    bool Collection::validateSpecifications(const std::string& body, query_options *options) {
        bool_result *r = kuzzle_collection_validate_specifications(_collection, const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        bool ret = r->result;
        kuzzle_free_bool_result(r);

        return ret;
    }

    void Collection::deleteSpecifications(const std::string& index, const std::string& collection, query_options *options) {
        error_result *r = kuzzle_collection_delete_specifications(_collection, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), options);
        if (r != nullptr)
            throwExceptionFromStatus(r);

        kuzzle_free_error_result(r);
    }
}
