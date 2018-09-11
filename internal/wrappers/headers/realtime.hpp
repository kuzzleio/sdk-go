#ifndef _KUZZLE_REALTIME_HPP
#define _KUZZLE_REALTIME_HPP

#include "exceptions.hpp"
#include "core.hpp"
#include <map>
#include <functional>

typedef std::function<void(const kuzzleio::notification_result*)> NotificationListener;

namespace kuzzleio {

  class Kuzzle;

  class Realtime {
    realtime *_realtime;
    std::map<std::string, NotificationListener*> _listener_instances;
    Realtime();

    public:
      Realtime(Kuzzle* kuzzle);
      Realtime(Kuzzle* kuzzle, realtime* realtime);
      virtual ~Realtime();
      int count(const std::string& index, const std::string& collection, const std::string& roomId, query_options *options=nullptr);
      std::string list(const std::string& index, const std::string& collection, query_options *options=nullptr);
      void publish(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);
      std::string subscribe(const std::string& index, const std::string& collection, const std::string& body, NotificationListener* cb, room_options* options=nullptr);
      void unsubscribe(const std::string& roomId, query_options *options=nullptr);
      bool validate(const std::string& index, const std::string& collection, const std::string& body, query_options *options=nullptr);

      NotificationListener* getListener(const std::string& roomId);
  };
}

#endif
