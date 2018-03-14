#ifndef _KUZZLE_HPP_
#define _KUZZLE_HPP_

#include "exceptions.hpp"
#include "core.hpp"
#include "listeners.hpp"
#include "event_emitter.hpp"
#include <string>
#include <iostream>
#include <vector>
#include <map>

namespace kuzzleio {
  class Document;
  class Index;
  class Server;

  class Kuzzle : public KuzzleEventEmitter {
    private:
      std::map<int, EventListener*>  _listener_instances;

    public:
      kuzzle *_kuzzle;
      Document *document;
      Index  *index;
      Server *server;

      Kuzzle(const std::string& host, options *options=NULL);
      virtual ~Kuzzle();

      token_validity* checkToken(const std::string& token);
      char* connect();
      json_object* createMyCredentials(const std::string& strategy, json_object* credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;

      bool deleteMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      json_object* getMyCredentials(const std::string& strategy, query_options *options=NULL) Kuz_Throw_KuzzleException;
      json_object* updateMyCredentials(const std::string& strategy, json_object* credentials, query_options *options=NULL) Kuz_Throw_KuzzleException;
      bool validateMyCredentials(const std::string& strategy, json_object* credentials, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string login(const std::string& strategy, json_object* credentials, int expiresIn) Kuz_Throw_KuzzleException;
      std::string login(const std::string& strategy, json_object* credentials) Kuz_Throw_KuzzleException;
      statistics* getAllStatistics(query_options* options=NULL) Kuz_Throw_KuzzleException;
      statistics* getStatistics(time_t start, time_t end, query_options* options=NULL) Kuz_Throw_KuzzleException;
      bool getAutoRefresh(const std::string& index, query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::string getJwt();
      json_object* getMyRights(query_options* options=NULL) Kuz_Throw_KuzzleException;
      std::vector<std::string> listIndexes(query_options* options=NULL) Kuz_Throw_KuzzleException;
      void disconnect();
      void logout();
      kuzzle_response* query(kuzzle_request* query, query_options* options=NULL) Kuz_Throw_KuzzleException;
      Kuzzle* replayQueue();
      Kuzzle* setAutoReplay(bool autoReplay);
      Kuzzle* setDefaultIndex(const std::string& index);
      Kuzzle* setJwt(const std::string& jwt);
      Kuzzle* startQueuing();
      Kuzzle* stopQueuing();
      Kuzzle* unsetJwt();
      json_object* updateSelf(user_data* content, query_options* options=NULL) Kuz_Throw_KuzzleException;
      user* whoAmI() Kuz_Throw_KuzzleException;
      Kuzzle* flushQueue();
      json_object* getVolatile();
      Kuzzle* setVolatile(json_object* volatiles);
      std::map<int, EventListener*> getListeners();

      virtual KuzzleEventEmitter* addListener(Event event, EventListener* listener);
      virtual KuzzleEventEmitter* removeListener(Event event, EventListener* listener);
      virtual KuzzleEventEmitter* removeAllListeners(Event event);
      virtual KuzzleEventEmitter* once(Event event, EventListener* listener);
      virtual int listenerCount(Event event);

  };
}

#endif
