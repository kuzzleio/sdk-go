#ifndef _AUTH_HPP
#define _AUTH_HPP

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {
  class Auth {
    auth *_auth;
    Auth();

    public:
      Auth(Kuzzle *kuzzle);
      virtual ~Auth();
      token_validity* checkToken(const std::string& token);
      json_object* createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;
      bool credentialsExist(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      void deleteMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
  };
}

#endif