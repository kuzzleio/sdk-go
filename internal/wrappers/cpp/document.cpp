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
#include "document.hpp"

namespace kuzzleio {
    Document::Document(Kuzzle* kuzzle) {
        _document = new document();
        kuzzle_new_document(_document, kuzzle->_kuzzle);
    }

    Document::Document(Kuzzle* kuzzle, document *document) {
        _document = document;
        kuzzle_new_document(_document, kuzzle->_kuzzle);
    }

    Document::~Document() {
        unregisterDocument(_document);
        delete(_document);
    }

    int Document::count(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
      int_result *r = kuzzle_document_count(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
      if (r->error != nullptr)
          throwExceptionFromStatus(r);
      int ret = r->result;
      kuzzle_free_int_result(r);
      return ret;
    }

    bool Document::exists(const std::string& index, const std::string& collection, const std::string& id, query_options *options) {
        bool_result *r = kuzzle_document_exists(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        bool ret = r->result;
        kuzzle_free_bool_result(r);
        return ret;
    }

    std::string Document::create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_create(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::string Document::createOrReplace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_create_or_replace(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::string Document::delete_(const std::string& index, const std::string& collection, const std::string& id, query_options *options) {
        string_result *r = kuzzle_document_delete(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::vector<std::string> Document::deleteByQuery(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        string_array_result *r = kuzzle_document_delete_by_query(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);

        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::vector<std::string> v;
        for (int i = 0; i < r->result_length; i++)
          v.push_back(r->result[i]);

        kuzzle_free_string_array_result(r);
        return v;
    }

    std::string Document::get(const std::string& index, const std::string& collection, const std::string& id, query_options *options) {
        string_result *r = kuzzle_document_get(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::string Document::replace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_replace(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::string Document::update(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_update(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(id.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    bool Document::validate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        bool_result *r = kuzzle_document_validate(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        bool ret = r->result;
        kuzzle_free_bool_result(r);
        return ret;
    }

    search_result* Document::search(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        search_result *r = kuzzle_document_search(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        return r;
    }

    std::string Document::mCreate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_mcreate(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::string Document::mCreateOrReplace(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_mcreate_or_replace(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::vector<std::string> Document::mDelete(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, query_options *options) {
        char **idsArray = new char *[ids.size()];
        int i = 0;
        for (auto& id : ids) {
          idsArray[i] = const_cast<char*>(id.c_str());
          i++;
        }

        string_array_result *r = kuzzle_document_mdelete(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), idsArray, ids.size(), options);
        delete[] idsArray;
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::vector<std::string> v;
        for (int i = 0; i < r->result_length; i++)
          v.push_back(r->result[i]);

        kuzzle_free_string_array_result(r);
        return v;
    }

    std::string Document::mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash, query_options *options) {
        char **idsArray = new char *[ids.size()];
        int i = 0;
        for (auto& id : ids) {
          idsArray[i] = const_cast<char*>(id.c_str());
          i++;
        }

        string_result *r = kuzzle_document_mget(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), idsArray, ids.size(), includeTrash, options);
        delete[] idsArray;
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;

    }
    std::string Document::mReplace(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_mreplace(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

    std::string Document::mUpdate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
        string_result *r = kuzzle_document_mupdate(_document, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::string ret = r->result;
        kuzzle_free_string_result(r);
        return ret;
    }

}
