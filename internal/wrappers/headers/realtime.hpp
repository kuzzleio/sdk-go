#ifndef _KUZZLE_REALTIME_HPP
#define _KUZZLE_REALTIME_HPP

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {
  class Realtime {
    realtime *_realtime;
    Realtime();

    public:
      Realtime(Kuzzle* kuzzle);
      virtual ~Realtime();
      int count(const std::string& index, const std::string collection, const std::string roomId) Kuz_Throw_KuzzleException;
      void join(const std::string& index, const std::string collection, const std::string roomId, callback cb) Kuz_Throw_KuzzleException;
      std::string list(const std::string& index, const std::string collection) Kuz_Throw_KuzzleException;
      void publish(const std::string& index, const std::string collection, const std::string body) Kuz_Throw_KuzzleException;
      std::string subscribe(const std::string& index, const std::string collection, const std::string body, callback cb, room_options* options) Kuz_Throw_KuzzleException;
      void unsubscribe(const std::string& roomId) Kuz_Throw_KuzzleException;
      bool validate(const std::string& index, const std::string collection, const std::string body) Kuz_Throw_KuzzleException;
  };
}

#endif
