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

#ifndef _KUZZLE_INDEX_HPP
#define _KUZZLE_INDEX_HPP
#include <stdlib.h>
#include <vector>
#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {
  class Kuzzle;
  class Index {
    kuzzle_index *_index;
    Index();

    public:
      Index(Kuzzle* kuzzle);
      Index(Kuzzle* kuzzle, kuzzle_index *index);
      virtual ~Index();
      void create(const std::string& index, query_options *options=nullptr);
      void delete_(const std::string& index, query_options *options=nullptr);
      std::vector<std::string> mDelete(const std::vector<std::string>& indexes, query_options *options=nullptr);
      bool exists(const std::string& index, query_options *options=nullptr);
      void refresh(const std::string& index, query_options *options=nullptr);
      void refreshInternal(query_options *options=nullptr);
      void setAutoRefresh(const std::string& index, bool autoRefresh, query_options *options=nullptr);
      bool getAutoRefresh(const std::string& index, query_options *options=nullptr);
      std::vector<std::string> list(query_options *options=nullptr);
  };
}

#endif
