// Copyright 2015-2017 Kuzzle
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
  class Collection;
  class Document;
  class Auth;
  class Index;
  class Server;
  class Realtime;

  class Kuzzle : public KuzzleEventEmitter {
    private:
      std::map<int, EventListener*>  _listener_instances;

    public:
      kuzzle *_kuzzle;
      Auth *auth;
      Index  *index;
      Server *server;
      Collection *collection;
      Document *document;
      Realtime *realtime;

      Kuzzle(const std::string& host, options *options=NULL);
      virtual ~Kuzzle();

      char* connect();
      bool createIndex(const std::string& index, query_options* options=NULL) Kuz_Throw_KuzzleException;

      statistics* getAllStatistics(query_options* options=NULL) Kuz_Throw_KuzzleException;
      statistics* getStatistics(time_t start, time_t end, query_options* options=NULL) Kuz_Throw_KuzzleException;
      void disconnect();
      kuzzle_response* query(kuzzle_request* query, query_options* options=NULL) Kuz_Throw_KuzzleException;
      Kuzzle* replayQueue();
      Kuzzle* setAutoReplay(bool autoReplay);
      Kuzzle* setDefaultIndex(const std::string& index);
      Kuzzle* startQueuing();
      Kuzzle* stopQueuing();
      Kuzzle* flushQueue();
      std::string getVolatile();
      Kuzzle* setVolatile(const std::string& volatiles);
      std::map<int, EventListener*> getListeners();

      virtual KuzzleEventEmitter* addListener(Event event, EventListener* listener);
      virtual KuzzleEventEmitter* removeListener(Event event, EventListener* listener);
      virtual KuzzleEventEmitter* removeAllListeners(Event event);
      virtual KuzzleEventEmitter* once(Event event, EventListener* listener);
      virtual int listenerCount(Event event);

  };
}

#endif
