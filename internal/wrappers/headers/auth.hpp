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

#ifndef _AUTH_HPP
#define _AUTH_HPP

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {

  class Kuzzle;

  class Auth {
    auth *_auth;
    Auth();

    public:
      Auth(Kuzzle *kuzzle);
      Auth(Kuzzle *kuzzle, auth *auth);
      virtual ~Auth();
      token_validity* checkToken(const std::string& token);
      std::string createMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;
      bool credentialsExist(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      void deleteMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      user* getCurrentUser() Kuz_Throw_KuzzleException;
      std::string getMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      user_right* getMyRights(query_options *options=NULL) Kuz_Throw_KuzzleException;
      std::vector<std::string> getStrategies(query_options *options=NULL) Kuz_Throw_KuzzleException;
      std::string login(const std::string& strategy, const std::string& credentials, int expiresIn) Kuz_Throw_KuzzleException;
      std::string login(const std::string& strategy, const std::string& credentials) Kuz_Throw_KuzzleException;
      void logout();
      std::string updateMyCredentials(const std::string& strategy, const std::string& credentials, query_options *options=NULL) Kuz_Throw_KuzzleException;
      user* updateSelf(const std::string& content, query_options* options=NULL) Kuz_Throw_KuzzleException;
      bool validateMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;
  };
}

#endif
