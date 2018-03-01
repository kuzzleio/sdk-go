#include "server.hpp"

namespace kuzzleio {
  Server::Server(Kuzzle* kuzzle) {
      _server = new server();
      kuzzle_new_server(_server, kuzzle->_kuzzle);
  }

  Server::~Server() {
      unregisterServer(_server);
      delete(_server);
  }

  bool Server::adminExists(query_options *options) Kuz_Throw_KuzzleException {
    bool_result* r = kuzzle_admin_exists(_server, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getAllStats(query_options* options) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_get_all_stats(_server, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getStats(time_t start, time_t end, query_options* options) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_get_stats(_server, start, end, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getLastStats(query_options* options) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_get_last_stats(_server, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::getConfig(query_options* options) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_get_config(_server, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  std::string Server::info(query_options* options) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_info(_server, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  // java wrapper for this method is in typemap.i
  long long Server::now(query_options* options) Kuz_Throw_KuzzleException {
    date_result *r = kuzzle_now(_server, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    long long ret = r->result;
    delete(r);
    return ret;
  }
  
}