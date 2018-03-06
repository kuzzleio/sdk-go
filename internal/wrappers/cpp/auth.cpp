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

  std::string Auth::createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_create_my_credentials(_auth, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  bool Auth::credentialsExist(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    bool_result* r = kuzzle_credentials_exist(_auth, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    delete(r);
    return ret;
  }

  void Auth::deleteMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    void_result *r = kuzzle_delete_my_credentials(_auth, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    delete(r);
  }

  user* Auth::getCurrentUser() Kuz_Throw_KuzzleException {
    user_result *r = kuzzle_get_current_user(_auth);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    kuzzle_free_user_result(r);
  }

  std::string Auth::getMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    string_result *r = kuzzle_get_my_credentials(_auth, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  user_right* Auth::getMyRights(query_options* options) Kuz_Throw_KuzzleException {
    user_rights_result *r = kuzzle_get_my_rights(_auth, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    user_right *ret = r->result;
    kuzzle_free_user_rights_result(r);
    return ret;
  }

  std::vector<std::string> Auth::getStrategies(query_options *options) Kuz_Throw_KuzzleException {
    string_array_result *r = kuzzle_get_strategies(_auth, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);

    std::vector<std::string> v;
    for (int i = 0; r->result[i]; i++)
        v.push_back(r->result[i]);

    kuzzle_free_string_array_result(r);
    return v;
  }
  
  std::string Auth::login(const std::string& strategy, json_object* credentials) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_login(_auth, const_cast<char*>(strategy.c_str()), credentials, NULL);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  std::string Auth::login(const std::string& strategy, json_object* credentials, int expires_in) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_login(_auth, const_cast<char*>(strategy.c_str()), credentials, &expires_in);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  void Auth::logout() {
    kuzzle_logout(_auth);
  }

  std::string Auth::updateMyCredentials(const std::string& strategy, json_object* credentials, query_options *options) Kuz_Throw_KuzzleException {
    string_result *r = kuzzle_update_my_credentials(_auth, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  user* Auth::updateSelf(const std::string& content, query_options* options) Kuz_Throw_KuzzleException {
    user_result *r = kuzzle_update_self(_auth, const_cast<char*>(content.c_str()), options);
    if (r->error != NULL)
      throwExceptionFromStatus(r);
    user *ret = r->result;
    kuzzle_free_user_result(r);
    return ret;
  }

  bool Auth::validateMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_validate_my_credentials(_auth, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }

}