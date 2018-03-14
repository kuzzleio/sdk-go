#include <exception>
#include <stdexcept>
#include "kuzzle.hpp"
#include "document.hpp"
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
    this->document = new Document(this, kuzzle_get_document_controller(this->_kuzzle));
    this->index = new Index(this, kuzzle_get_index_controller(_kuzzle));
    this->server = new Server(this, kuzzle_get_server_controller(_kuzzle));
    kuzzle_new_kuzzle(this->_kuzzle, const_cast<char*>(host.c_str()), (char*)"websocket", opts);
  }

  Kuzzle::~Kuzzle() {
    unregisterKuzzle(this->_kuzzle);
    delete(this->_kuzzle);
    delete(this->document);
    delete(this->index);
    delete(this->server);
  }

  token_validity* Kuzzle::checkToken(const std::string& token) {
    return kuzzle_check_token(_kuzzle, const_cast<char*>(token.c_str()));
  }

  char* Kuzzle::connect() {
    return kuzzle_connect(_kuzzle);
  }

  json_object* Kuzzle::createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    json_result* r = kuzzle_create_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error)
        throwExceptionFromStatus(r);
    json_object *ret = r->result;
    kuzzle_free_json_result(r);
    return ret;
  }

  bool Kuzzle::deleteMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_delete_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }

  json_object* Kuzzle::getMyCredentials(const std::string& strategy, query_options *options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_get_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    json_object *ret = r->result;
    kuzzle_free_json_result(r);
    return ret;
  }

  json_object* Kuzzle::updateMyCredentials(const std::string& strategy, json_object* credentials, query_options *options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_update_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    json_object *ret = r->result;
    kuzzle_free_json_result(r);
    return ret;
  }

  bool Kuzzle::validateMyCredentials(const std::string& strategy, json_object* credentials, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_validate_my_credentials(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }

  std::string Kuzzle::login(const std::string& strategy, json_object* credentials) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_login(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, NULL);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }
  std::string Kuzzle::login(const std::string& strategy, json_object* credentials, int expires_in) Kuz_Throw_KuzzleException {
    string_result* r = kuzzle_login(_kuzzle, const_cast<char*>(strategy.c_str()), credentials, &expires_in);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    std::string ret = r->result;
    kuzzle_free_string_result(r);
    return ret;
  }

  bool Kuzzle::getAutoRefresh(const std::string& index, query_options* options) Kuz_Throw_KuzzleException {
    bool_result *r = kuzzle_get_auto_refresh(_kuzzle, const_cast<char*>(index.c_str()), options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    bool ret = r->result;
    kuzzle_free_bool_result(r);
    return ret;
  }

  std::string Kuzzle::getJwt() {
    return kuzzle_get_jwt(_kuzzle);
  }

  json_object* Kuzzle::getMyRights(query_options* options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_get_my_rights(_kuzzle, options);
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    json_object *ret = r->result;
    kuzzle_free_json_result(r);
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

  void Kuzzle::logout() {
    kuzzle_logout(_kuzzle);
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

  Kuzzle* Kuzzle::setJwt(const std::string& jwt) {
    kuzzle_set_jwt(_kuzzle, const_cast<char*>(jwt.c_str()));
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

  Kuzzle* Kuzzle::unsetJwt() {
    kuzzle_unset_jwt(_kuzzle);
    return this;
  }

  json_object* Kuzzle::updateSelf(user_data* content, query_options* options) Kuz_Throw_KuzzleException {
    json_result *r = kuzzle_update_self(_kuzzle, content, options);
    if (r->error != NULL)
      throwExceptionFromStatus(r);
    json_object* ret = r->result;
    kuzzle_free_json_result(r);
    return ret;
  }

  user* Kuzzle::whoAmI() Kuz_Throw_KuzzleException {
    user_result *r = kuzzle_who_am_i(_kuzzle);
    if (r->error != NULL)
      throwExceptionFromStatus(r);
    user *ret = r->user;
    kuzzle_free_user_result(r);
    return ret;
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
