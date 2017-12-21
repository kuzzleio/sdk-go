#include <exception>
#include <stdexcept>
#include <string>
#include "kuzzle.hpp"
#include <iostream>
#include <vector>

namespace kuzzleio {

  KuzzleException::KuzzleException(int status, const std::string& error)
    : std::runtime_error(error), status(status) {}

  std::string KuzzleException::getMessage() const {
    return what();
  }

  Kuzzle::Kuzzle(const std::string& host, options *opts) {
    this->_kuzzle = (kuzzle*)malloc(sizeof(kuzzle));
    kuzzle_new_kuzzle(this->_kuzzle, const_cast<char*>(host.c_str()), (char*)"websocket", opts);
  }

  Kuzzle::~Kuzzle() {
    unregisterKuzzle(this->_kuzzle);
    free(this->_kuzzle);
  }

  token_validity* Kuzzle::checkToken(const std::string& token) {
    return kuzzle_check_token(_kuzzle, const_cast<char*>(token.c_str()));
  }

  char* Kuzzle::connect() {
    return kuzzle_connect(_kuzzle);
  }

  bool Kuzzle::createIndex(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_create_index(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  json_object* Kuzzle::createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    json_result* r = kuzzle_create_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  bool Kuzzle::deleteMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_delete_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  json_object* Kuzzle::getMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_get_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  json_object* Kuzzle::updateMyCredentials(const std::string& strategy, json_object* credentials, query_options *options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_update_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  bool Kuzzle::validateMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_validate_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  std::string Kuzzle::login(const std::string& strategy, json_object* credentials) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_login(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, NULL);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }
  std::string Kuzzle::login(const std::string& strategy, json_object* credentials, int expires_in) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_login(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, &expires_in);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  statistics* Kuzzle::getAllStatistics(query_options* options) Kuz_Throw_KuzzleException {
    all_statistics_result* r = kuzzle_get_all_statistics(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  statistics* Kuzzle::getStatistics(time_t start, time_t end, query_options* options) Kuz_Throw_KuzzleException {
    statistics_result* r = kuzzle_get_statistics(_kuzzle, start, end, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  bool Kuzzle::getAutoRefresh(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_get_auto_refresh(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  std::string Kuzzle::getJwt() {
    return kuzzle_get_jwt(_kuzzle);
  }

  json_object* Kuzzle::getMyRights(query_options* options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_get_my_rights(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  json_object* Kuzzle::getServerInfo(query_options* options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_get_server_info(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  collection_entry* Kuzzle::listCollections(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    collection_entry_result *r = kuzzle_list_collections(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  std::vector<std::string> Kuzzle::listIndexes(query_options* options) Kuz_Throw_KuzzleException {
    string_array_result *r = kuzzle_list_indexes(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);

    std::vector<std::string> v;
    for (int i = 0; r->result[i]; i++)
        v.push_back(r->result[i]);

    return v;
  }

  void Kuzzle::disconnect() {
    kuzzle_disconnect(_kuzzle);
  }

  void Kuzzle::logout() {
    kuzzle_logout(_kuzzle);
  }

  kuzzle_response* Kuzzle::query(kuzzle_request* query, query_options* options) Kuz_Throw_KuzzleException {
    kuzzle_response *r = kuzzle_query(_kuzzle, query, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);

    return r;
  }

  shards* Kuzzle::refreshIndex(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    shards_result *r = kuzzle_refresh_index(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }

  // java wrapper for this method is in typemap.i
  long long Kuzzle::now(query_options* options) Kuz_Throw_KuzzleException {
    date_result *r = kuzzle_now(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(*r);
    return r->result;
  }
}
