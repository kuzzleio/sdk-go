#ifndef _KUZZLE_REALTIME_HPP
#define _KUZZLE_REALTIME_HPP

#include "exceptions.hpp"
#include "core.hpp"
#include <map>

namespace kuzzleio {

  class EventListener;
  class NotificationListener;
  class SubscribeListener;

  class Realtime {
    realtime *_realtime;
    std::map<std::string, NotificationListener*> _listener_instances;
    Realtime();

    public:
      Realtime(Kuzzle* kuzzle);
      Realtime(Kuzzle* kuzzle, realtime* realtime);
      virtual ~Realtime();
      int count(const std::string& index, const std::string& collection, const std::string& roomId) Kuz_Throw_KuzzleException;
      void join(const std::string& index, const std::string& collection, const std::string& roomId, room_options* options, NotificationListener* cb) Kuz_Throw_KuzzleException;
      std::string list(const std::string& index, const std::string &collection) Kuz_Throw_KuzzleException;
      void publish(const std::string& index, const std::string& collection, const std::string& body) Kuz_Throw_KuzzleException;
      std::string subscribe(const std::string& index, const std::string& collection, const std::string& body, NotificationListener* cb, room_options* options=NULL) Kuz_Throw_KuzzleException;
      void unsubscribe(const std::string& roomId) Kuz_Throw_KuzzleException;
      bool validate(const std::string& index, const std::string& collection, const std::string& body) Kuz_Throw_KuzzleException;

      NotificationListener* getListener(const std::string& roomId);
  };
}

#endif
