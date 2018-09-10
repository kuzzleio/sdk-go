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

#ifndef _DOCUMENT_HPP_
#define _DOCUMENT_HPP_

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {

    class Document {
        document* _document;
        Document();

        public:
            Document(Kuzzle* kuzzle);
            Document(Kuzzle* kuzzle, document *document);
            virtual ~Document();
            int count(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            bool exists(const std::string& index, const std::string& collection, const std::string& id, query_options *options=nullptr);
            std::string create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options=nullptr);
            std::string createOrReplace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options=nullptr);
            std::string delete_(const std::string& index, const std::string& collection, const std::string& id, query_options *options=nullptr);
            std::vector<std::string> deleteByQuery(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            std::string get(const std::string& index, const std::string& collection, const std::string& id, query_options *options=nullptr);
            std::string replace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options=nullptr);
            std::string update(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, query_options *options=nullptr);
            bool validate(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            search_result* search(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            std::string mCreate(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            std::string mCreateOrReplace(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            std::vector<std::string> mDelete(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, query_options *options=nullptr);
            std::string mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash, query_options *options=nullptr);
            std::string mReplace(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
            std::string mUpdate(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
    };
}

#endif
