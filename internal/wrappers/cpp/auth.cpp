#include "auth.hpp"

namespace kuzzleio {
  Auth::Auth(Kuzzle* kuzzle) {
      _auth = new auth();
      kuzzle_new_auth(_auth, kuzzle->_kuzzle);
  }

  Auth::~Auth() {
      unregisterAuth(_auth);
      delete(_auth);
  }

  token_validity* Auth::checkToken(const std::string& token) {
    return kuzzle_check_token(_auth, const_cast<char*>(token.c_str()));
  }
}