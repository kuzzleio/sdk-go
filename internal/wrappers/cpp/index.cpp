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
#include "index.hpp"
#include <string>

namespace kuzzleio {

    Index::Index(Kuzzle* kuzzle) {
        _index = new kuzzle_index();
        kuzzle_new_index(_index, kuzzle->_kuzzle);
    }

    Index::Index(Kuzzle* kuzzle, kuzzle_index *index) {
      _index = index;
      kuzzle_new_index(index, kuzzle->_kuzzle);
    }

    Index::~Index() {
        unregisterIndex(_index);
        delete(_index);
    }

    void Index::create(const std::string& index, query_options *options) {
        error_result *r = kuzzle_index_create(_index, const_cast<char*>(index.c_str()), options);
        if (r != nullptr)
            throwExceptionFromStatus(r);
        kuzzle_free_error_result(r);
    }

    void Index::delete_(const std::string& index, query_options *options) {
        error_result *r = kuzzle_index_delete(_index, const_cast<char*>(index.c_str()), options);
        if (r != nullptr)
            throwExceptionFromStatus(r);
        kuzzle_free_error_result(r);
    }

    std::vector<std::string> Index::mDelete(const std::vector<std::string>& indexes, query_options *options) {
        char **indexesArray = new char *[indexes.size()];
        int i = 0;
        for (auto& index : indexes) {
          indexesArray[i] = const_cast<char*>(index.c_str());
          i++;
        }
        string_array_result *r = kuzzle_index_mdelete(_index, indexesArray, indexes.size(), options);

        delete[] indexesArray;
        if (r->error != nullptr)
          throwExceptionFromStatus(r);

        std::vector<std::string> v;
        for (int i = 0; i < r->result_length; i++)
          v.push_back(r->result[i]);

        kuzzle_free_string_array_result(r);
        return v;
    }

    bool Index::exists(const std::string& index, query_options *options) {
        bool_result *r = kuzzle_index_exists(_index, const_cast<char*>(index.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);
        bool ret = r->result;
        kuzzle_free_bool_result(r);
        return ret;
    }

    void Index::refresh(const std::string& index, query_options *options) {
        error_result *r = kuzzle_index_refresh(_index, const_cast<char*>(index.c_str()), options);
        if (r != nullptr)
            throwExceptionFromStatus(r);
        kuzzle_free_error_result(r);
    }

    void Index::refreshInternal(query_options *options) {
        error_result *r = kuzzle_index_refresh_internal(_index, options);
        if (r != nullptr)
            throwExceptionFromStatus(r);
        kuzzle_free_error_result(r);
    }

    void Index::setAutoRefresh(const std::string& index, bool autoRefresh, query_options *options) {
      error_result *r = kuzzle_index_set_auto_refresh(_index, const_cast<char*>(index.c_str()), autoRefresh, options);
      if (r != nullptr)
          throwExceptionFromStatus(r);
        kuzzle_free_error_result(r);
    }

    bool Index::getAutoRefresh(const std::string& index, query_options *options) {
        bool_result *r = kuzzle_index_get_auto_refresh(_index, const_cast<char*>(index.c_str()), options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);
        bool ret = r->result;
        kuzzle_free_bool_result(r);
        return ret;
    }

    std::vector<std::string> Index::list(query_options *options) {
        string_array_result *r = kuzzle_index_list(_index, options);
        if (r->error != nullptr)
            throwExceptionFromStatus(r);

        std::vector<std::string> v;
        for (int i = 0; i < r->result_length; i++)
          v.push_back(r->result[i]);

        kuzzle_free_string_array_result(r);
        return v;
    }
}
