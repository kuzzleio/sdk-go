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

  json_object* Auth::createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    json_result* r = kuzzle_create_my_credentials(_auth, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error)
        throwExceptionFromStatus(r);
    json_object *ret = r->result;
    delete(r);
    return ret;
  }
}