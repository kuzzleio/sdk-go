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

#ifndef _KUZZLE_SERVER_HPP
#define _KUZZLE_SERVER_HPP

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {
  class Server {
    server *_server;
    Server();

    public:
      Server(Kuzzle* kuzzle);
      Server(Kuzzle *kuzzle, server *server);
      virtual ~Server();
      bool adminExists(query_options *options) Kuz_Throw_KuzzleException;
      std::string getAllStats(query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getStats(time_t start, time_t end, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getLastStats(query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getConfig(query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string info(query_options* options=NULL) Kuz_Throw_KuzzleException;
      long long now(query_options* options=NULL) Kuz_Throw_KuzzleException;
  };
}

#endif
