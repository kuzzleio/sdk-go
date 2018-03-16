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