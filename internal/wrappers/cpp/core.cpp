#include <exception>
#include <stdexcept>
#include <string>
#include "kuzzle.hpp"

#include "errors.cpp"

namespace kuzzleio {
  Kuzzle::Kuzzle(std::string host, options *opts) {
    this->_kuzzle = (kuzzle*)malloc(sizeof(kuzzle));
    kuzzle_new_kuzzle(this->_kuzzle, (char*)host.c_str(), (char*)"websocket", opts);
  }

  Kuzzle::~Kuzzle() {
    unregisterKuzzle(this->_kuzzle);
    free(this->_kuzzle);
  }

  long long Kuzzle::now(query_options *options) Kuz_Throw_KuzzleError {
    throw BadRequestError("Toto");
  }
}
