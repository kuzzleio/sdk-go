#include <exception>
#include <stdexcept>
#include <string>
#include "kuzzle.hpp"

namespace kuzzleio {

  KuzzleException::KuzzleException(int status, const std::string& error, const std::string& stack)
    : std::runtime_error(error) {
    this->status = status;
    this->stack = stack;
  }

  std::string KuzzleException::getMessage() {
    return what();
  }

  Kuzzle::Kuzzle(const std::string& host, options *opts) {
    this->_kuzzle = (kuzzle*)malloc(sizeof(kuzzle));
    kuzzle_new_kuzzle(this->_kuzzle, (char*)host.c_str(), (char*)"websocket", opts);
  }

  Kuzzle::~Kuzzle() {
    unregisterKuzzle(this->_kuzzle);
    free(this->_kuzzle);
  }

  token_validity* Kuzzle::checkToken(const std::string& token) {
    return kuzzle_check_token(_kuzzle, (char*)token.c_str());
  }

  char* Kuzzle::connect() {
    return kuzzle_connect(_kuzzle);
  }

  bool_result* Kuzzle::createIndex(const std::string& index, query_options* options) {
    return kuzzle_create_index(_kuzzle, (char*)index.c_str(), options);
  }

  json_result* Kuzzle::createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) {
    return kuzzle_create_my_credentials(_kuzzle, (char*)strategy.c_str(), credentials, options);
  }

  bool Kuzzle::deleteMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_delete_my_credentials(_kuzzle, (char*)strategy.c_str(), options);
    throwExceptionFromStatus(*r);
  }

}
