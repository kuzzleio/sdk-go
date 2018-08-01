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

#ifndef _COLLECTION_HPP_
#define _COLLECTION_HPP_

#include <iostream>
#include <list>
#include "core.hpp"
#include "exceptions.hpp"

namespace kuzzleio {
    class Kuzzle;
    class Collection {
        collection* _collection;
        Collection();

        public:
            Collection(Kuzzle* kuzzle);
            Collection(Kuzzle* kuzzle, collection *collection);
            virtual ~Collection();
            void create(const std::string& index, const std::string& collection, const std::string* body=NULL, query_options *options=NULL) Kuz_Throw_KuzzleException;
            bool exists(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string list(const std::string& index, query_options *options=NULL) Kuz_Throw_KuzzleException;
            void truncate(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string getMapping(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            void updateMapping(const std::string& index, const std::string& collection, const std::string& body, query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string getSpecifications(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            search_result* searchSpecifications(query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string updateSpecifications(const std::string& index, const std::string& collection, const std::string& body, query_options *options=NULL) Kuz_Throw_KuzzleException;
            bool validateSpecifications(const std::string& body, query_options *options=NULL) Kuz_Throw_KuzzleException;
            void deleteSpecifications(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
    };
}

#endif
