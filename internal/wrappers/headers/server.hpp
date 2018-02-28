#ifndef _KUZZLE_SERVER_HPP
#define _KUZZLE_SERVER_HPP

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {
  class Server {
    server *_server;
    Server();

    public:
      Server(Kuzzle* kuzzle);
      virtual ~Server();
      std::string getAllStats(query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getLastStats(time_t start, time_t end, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getConfig(query_options* options=NULL) Kuz_Throw_KuzzleException;      
      std::string info(query_options* options=NULL) Kuz_Throw_KuzzleException;
      long long now(query_options* options=NULL) Kuz_Throw_KuzzleException;      
  };
}

#endif