#include <exception>
#include <stdexcept>
#include "kuzzle.hpp"
#include "auth.hpp"
#include "index.hpp"
#include "server.hpp"
#include "collection.hpp"
#include "document.hpp"
#include "auth.hpp"
#include "index.hpp"
#include "server.hpp"
#include "realtime.hpp"
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
    this->document = new Document(this, kuzzle_get_document_controller(this->_kuzzle));
    this->auth = new Auth(this, kuzzle_get_auth_controller(_kuzzle));
    this->index = new Index(this, kuzzle_get_index_controller(_kuzzle));
    this->server = new Server(this, kuzzle_get_server_controller(_kuzzle));
    this->collection = new Collection(this, kuzzle_get_collection_controller(this->_kuzzle));
    this->realtime = new Realtime(this, kuzzle_get_realtime_controller(this->_kuzzle));
    kuzzle_new_kuzzle(this->_kuzzle, const_cast<char*>(host.c_str()), (char*)"websocket", opts);
  }

  Kuzzle::~Kuzzle() {
    unregisterKuzzle(this->_kuzzle);
    delete(this->_kuzzle);
    delete(this->document);
    delete(this->auth);
    delete(this->index);
    delete(this->server);
    delete(this->collection);
    delete(this->realtime);
  }

  char* Kuzzle::connect() {
    return kuzzle_connect(_kuzzle);
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

  Kuzzle* Kuzzle::setVolatile(const std::string& volatiles) {
    kuzzle_set_volatile(_kuzzle, const_cast<char*>(volatiles.c_str()));
    return this;
  }

  std::string Kuzzle::getVolatile() {
    return std::string(kuzzle_get_volatile(_kuzzle));
  }

  void trigger_event_listener(int event, char* res, void* data) {
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
