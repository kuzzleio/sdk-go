#include "kuzzle.hpp"
#include "realtime.hpp"

namespace kuzzleio {
  Realtime::Realtime(Kuzzle *kuzzle) {
    _realtime = new realtime();
    kuzzle_new_realtime(_realtime, kuzzle->_kuzzle);
  }

  Realtime::Realtime(Kuzzle *kuzzle, realtime *realtime) {
    _realtime = realtime;
    kuzzle_new_realtime(_realtime, kuzzle->_kuzzle);
  }

  Realtime::~Realtime() {
    unregisterRealtime(_realtime);
    delete(_realtime);
  }

  int Realtime::count(const std::string& index, const std::string& collection, const std::string& roomId, query_options *options) {
    int_result *r = kuzzle_realtime_count(_realtime, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(roomId.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    int ret = r->result;
    kuzzle_free_int_result(r);
    return ret;
  }

  NotificationListener* Realtime::getListener(const std::string& roomId) {
    return _listener_instances[roomId];
  }

  void call_subscribe_cb(notification_result* res, void* data) {
    if (data) {
      NotificationListener* listener = static_cast<Realtime*>(data)->getListener(res->room_id);

      if (listener) {
        (*listener)(res);
      }
    }
  }

  std::string Realtime::list(const std::string& index, const std::string& collection, query_options *options) {
    string_result *r = kuzzle_realtime_list(_realtime, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  void Realtime::publish(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
    error_result *r = kuzzle_realtime_publish(_realtime, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(body.c_str()), options);
    if (r != nullptr)
        throwExceptionFromStatus(r);
    kuzzle_free_error_result(r);
  }

  std::string Realtime::subscribe(const std::string& index, const std::string& collection, const std::string& body, NotificationListener* cb, room_options* options) {
    subscribe_result *r = kuzzle_realtime_subscribe(_realtime, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()),  const_cast<char*>(body.c_str()), call_subscribe_cb, this, options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);

    std::string roomId = r->room;
    std::string channel = r->channel;

    _listener_instances[channel] = cb;
    kuzzle_free_subscribe_result(r);
    return roomId;
  }

  void Realtime::unsubscribe(const std::string& roomId, query_options *options) {
    error_result *r = kuzzle_realtime_unsubscribe(_realtime, const_cast<char*>(roomId.c_str()), options);
    if (r != nullptr)
        throwExceptionFromStatus(r);

    _listener_instances[roomId] = nullptr;
    kuzzle_free_error_result(r);
  }

  bool Realtime::validate(const std::string& index, const std::string& collection, const std::string& body, query_options *options) {
    bool_result *r = kuzzle_realtime_validate(_realtime, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()),  const_cast<char*>(body.c_str()), options);
    if (r->error != nullptr)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }
}
