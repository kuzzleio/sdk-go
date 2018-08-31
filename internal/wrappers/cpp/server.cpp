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
#include "server.hpp"

namespace kuzzleio {
  Server::Server(Kuzzle* kuzzle) {
      _server = new server();
      kuzzle_new_server(_server, kuzzle->_kuzzle);
  }

  Server::Server(Kuzzle* kuzzle, server *server) {
      _server = server;
      kuzzle_new_server(_server, kuzzle->_kuzzle);
  }

  Server::~Server() {
      unregisterServer(_server);
      delete(_server);
  }

  bool Server::adminExists(query_options *options) {
    bool_result* r = kuzzle_admin_exists(_server, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getAllStats(query_options* options) {
    string_result* r = kuzzle_get_all_stats(_server, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getStats(time_t start, time_t end, query_options* options) {
    string_result* r = kuzzle_get_stats(_server, start, end, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getLastStats(query_options* options) {
    string_result* r = kuzzle_get_last_stats(_server, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getConfig(query_options* options) {
    string_result* r = kuzzle_get_config(_server, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::info(query_options* options) {
    string_result* r = kuzzle_info(_server, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  // java wrapper for this method is in typemap.i
  long long Server::now(query_options* options) {
    date_result *r = kuzzle_now(_server, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    long long ret = r->result;
    delete(r);
    return ret;
  }

}
