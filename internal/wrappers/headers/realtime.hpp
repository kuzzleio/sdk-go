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
  };
}

#endif