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
      Kuzzle *_kuzzle;

      Auth(Kuzzle *kuzzle);
      Auth(Kuzzle *kuzzle, auth *auth);
      virtual ~Auth();
      token_validity* checkToken(const std::string& token);
      std::string createMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options=nullptr);
      bool credentialsExist(const std::string& strategy, query_options *options=nullptr);
      void deleteMyCredentials(const std::string& strategy, query_options *options=nullptr);
      kuzzle_user* getCurrentUser();
      std::string getMyCredentials(const std::string& strategy, query_options *options=nullptr);      
      user_right* getMyRights(query_options *options=nullptr);
      std::vector<std::string> getStrategies(query_options *options=nullptr);
      std::string login(const std::string& strategy, const std::string& credentials, int expiresIn);
      std::string login(const std::string& strategy, const std::string& credentials);
      void logout() noexcept;
      void setJwt(const std::string& jwt) noexcept;
      std::string updateMyCredentials(const std::string& strategy, const std::string& credentials, query_options *options=nullptr);      
      kuzzle_user* updateSelf(const std::string& content, query_options* options=nullptr);      
      bool validateMyCredentials(const std::string& strategy, const std::string& credentials, query_options* options=nullptr);
  };
}

#endif
