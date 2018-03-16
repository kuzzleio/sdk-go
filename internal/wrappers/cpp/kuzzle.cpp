#include <exception>
#include <stdexcept>
#include "kuzzle.hpp"
#include "auth.hpp"
#include "index.hpp"
#include "server.hpp"
#include <iostream>
#include <vector>

namespace kuzzleio {

  KuzzleException::KuzzleException(int status, const std::string& error)
    : std::runtime_error(error), status(status) {}

  std::string KuzzleException::getMessage() const {
    return what();
  }

  Kuzzle::Kuzzle(const std::string& host, options *opts) {
    this->_kuzzle = new kuzzle();
    this->auth = new Auth(this, kuzzle_get_auth_controller(_kuzzle));
    this->index = new Index(this, kuzzle_get_index_controller(_kuzzle));
    this->server = new Server(this, kuzzle_get_server_controller(_kuzzle));
    kuzzle_new_kuzzle(this->_kuzzle, const_cast<char*>(host.c_str()), (char*)"websocket", opts);
  }

  Kuzzle::~Kuzzle() {
    unregisterKuzzle(this->_kuzzle);
    delete(this->_kuzzle);
    delete(this->auth);
    delete(this->index);
    delete(this->server);
  }

  char* Kuzzle::connect() {
    return kuzzle_connect(_kuzzle);
  }

  bool Kuzzle::getAutoRefresh(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_get_auto_refresh(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }

  collection_entry* Kuzzle::listCollections(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    collection_entry_result *r = kuzzle_list_collections(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    collection_entry *ret = r->result;
    kuzzle_free_collection_entry_result(r);
    return ret;
  }

  std::vector<std::string> Kuzzle::listIndexes(query_options* options) Kuz_Throw_KuzzleException {
    string_array_result *r = kuzzle_list_indexes(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);

    std::vector<std::string> v;
    for (int i = 0; r->result[i]; i++)
        v.push_back(r->result[i]);

    kuzzle_free_string_array_result(r);
    return v;
  }

  void Kuzzle::disconnect() {
    kuzzle_disconnect(_kuzzle);
  }

  kuzzle_response* Kuzzle::query(kuzzle_request* query, query_options* options) Kuz_Throw_KuzzleException {
    kuzzle_response *r = kuzzle_query(_kuzzle, query, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    return r;
  }

  Kuzzle* Kuzzle::replayQueue() {
    kuzzle_replay_queue(_kuzzle);
    return this;
  }

  Kuzzle* Kuzzle::setAutoReplay(bool autoReplay) {
    kuzzle_set_auto_replay(_kuzzle, autoReplay);
    return this;
  }

  Kuzzle* Kuzzle::setDefaultIndex(const std::string& index) {
    kuzzle_set_default_index(_kuzzle, const_cast<char*>(index.c_str()));
    return this;
  }

  Kuzzle* Kuzzle::startQueuing() {
    kuzzle_start_queuing(_kuzzle);
    return this;
  }

  Kuzzle* Kuzzle::stopQueuing() {
    kuzzle_stop_queuing(_kuzzle);
    return this;
  }

  Kuzzle* Kuzzle::flushQueue() {
    kuzzle_flush_queue(_kuzzle);
    return this;
  }

  Kuzzle* Kuzzle::setVolatile(json_object *volatiles) {
    kuzzle_set_volatile(_kuzzle, volatiles);
    return this;
  }

  json_object* Kuzzle::getVolatile() {
    return kuzzle_get_volatile(_kuzzle);
  }

  void trigger_event_listener(int event, json_object* res, void* data) {
    ((Kuzzle*)data)->getListeners()[event]->trigger(res);
  }

  std::map<int, EventListener*> Kuzzle::getListeners() {
    return _listener_instances;
  }

  KuzzleEventEmitter* Kuzzle::addListener(Event event, EventListener* listener) {
    kuzzle_add_listener(_kuzzle, event, &trigger_event_listener, this);
    _listener_instances[event] = listener;

    return this;
  }

  KuzzleEventEmitter* Kuzzle::removeListener(Event event, EventListener* listener) {
    kuzzle_remove_listener(_kuzzle, event, (void*)&trigger_event_listener);
    _listener_instances[event] = NULL;

    return this;
  }

  KuzzleEventEmitter* Kuzzle::removeAllListeners(Event event) {
    kuzzle_remove_all_listeners(_kuzzle, event);

    return this;
  }

  KuzzleEventEmitter* Kuzzle::once(Event event, EventListener* listener) {
    kuzzle_once(_kuzzle, event, &trigger_event_listener, this);
  }

  int Kuzzle::listenerCount(Event event) {
    return kuzzle_listener_count(_kuzzle, event);
  }

}
