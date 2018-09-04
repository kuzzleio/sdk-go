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
#include "auth.hpp"

namespace kuzzleio {
  Auth::Auth(Kuzzle *kuzzle) {
      _auth = new auth();
      _kuzzle = kuzzle;
      kuzzle_new_auth(_auth, kuzzle->_kuzzle);
  }

  Auth::Auth(Kuzzle *kuzzle, auth *auth) {
    _auth = auth;
    _kuzzle = kuzzle;
    kuzzle_new_auth(_auth, kuzzle->_kuzzle);
  }

  Auth::~Auth() {
      unregisterAuth(_auth);
      delete(_auth);
  }

  token_validity* Auth::checkToken(const std::string& token) {
    return kuzzle_check_token(_auth, const_cast<char*>(token.c_str()));
  }

  std::string Auth::createMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options) {
    string_result* r = kuzzle_create_my_credentials(_auth, const_cast<char*>(strategy.c_str()), const_cast<char*>(credentials.c_str()), options);
    if (r->error)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    delete(r);
    return ret;
  }

  bool Auth::credentialsExist(const std::string& strategy, query_options *options) {
    bool_result* r = kuzzle_credentials_exist(_auth, const_cast<char*>(strategy.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    delete(r);
    return ret;
  }

  void Auth::deleteMyCredentials(const std::string& strategy, query_options *options) {
    error_result *r = kuzzle_delete_my_credentials(_auth, const_cast<char*>(strategy.c_str()), options);
    if (r != nullptr)
        throwExceptionFromStatus(r);
    delete(r);
  }

  kuzzle_user* Auth::getCurrentUser() {
    user_result *r = kuzzle_get_current_user(_auth);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);

    kuzzle_user *u = r->result;
    kuzzle_free_user_result(r);
    return u;
  }

  std::string Auth::getMyCredentials(const std::string& strategy, query_options *options) {
    string_result *r = kuzzle_get_my_credentials(_auth, const_cast<char*>(strategy.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  user_right* Auth::getMyRights(query_options* options) {
    user_rights_result *r = kuzzle_get_my_rights(_auth, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    user_right *ret = r->result;
    kuzzle_free_user_rights_result(r);
    return ret;
  }

  std::vector<std::string> Auth::getStrategies(query_options *options) {
    string_array_result *r = kuzzle_get_strategies(_auth, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);

    std::vector<std::string> v;
    for (int i = 0; r->result[i]; i++)
        v.push_back(r->result[i]);

    kuzzle_free_string_array_result(r);
    return v;
  }

  std::string Auth::login(const std::string& strategy, const std::string& credentials) {
    string_result* r = kuzzle_login(_auth, const_cast<char*>(strategy.c_str()), const_cast<char*>(credentials.c_str()), nullptr);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  std::string Auth::login(const std::string& strategy, const std::string& credentials, int expires_in) {
    string_result* r = kuzzle_login(_auth, const_cast<char*>(strategy.c_str()), const_cast<char*>(credentials.c_str()), &expires_in);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  void Auth::logout() noexcept {
    kuzzle_logout(_auth);
  }

  void Auth::setJwt(const std::string& jwt) noexcept {
    kuzzle_set_jwt(_kuzzle->_kuzzle, const_cast<char*>(jwt.c_str()));
  }

  std::string Auth::updateMyCredentials(const std::string& strategy, const std::string& credentials, query_options *options) {
    string_result *r = kuzzle_update_my_credentials(_auth, const_cast<char*>(strategy.c_str()), const_cast<char*>(credentials.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  kuzzle_user* Auth::updateSelf(const std::string& content, query_options* options) {
    user_result *r = kuzzle_update_self(_auth, const_cast<char*>(content.c_str()), options);
    if (r->error != nullptr)
      throwExceptionFromStatus(r);
    kuzzle_user *ret = r->result;
    kuzzle_free_user_result(r);
    return ret;
  }

  bool Auth::validateMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options) {
    bool_result *r = kuzzle_validate_my_credentials(_auth, const_cast<char*>(strategy.c_str()), const_cast<char*>(credentials.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }

}
