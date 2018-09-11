// Copyright 2015-2018 Kuzzle
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
#include "event_emitter.hpp"
#include <string>
#include <iostream>
#include <vector>
#include <map>
#include <functional>

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

      Kuzzle(const std::string& host, options *options=nullptr);
      virtual ~Kuzzle();

      char* connect() noexcept;

      std::string getJwt() noexcept;
      void disconnect() noexcept;
      kuzzle_response* query(kuzzle_request* query, query_options* options=nullptr);
      Kuzzle* playQueue() noexcept;
      Kuzzle* setAutoReplay(bool autoReplay) noexcept;
      Kuzzle* startQueuing() noexcept;
      Kuzzle* stopQueuing() noexcept;
      Kuzzle* flushQueue() noexcept;
      std::string getVolatile() noexcept;
      Kuzzle* setVolatile(const std::string& volatiles) noexcept;
      std::map<int, EventListener*> getListeners() noexcept;
      void emitEvent(Event& event, const std::string& body) noexcept;

      virtual KuzzleEventEmitter* addListener(Event event, EventListener* listener) override;
      virtual KuzzleEventEmitter* removeListener(Event event, EventListener* listener) override;
      virtual KuzzleEventEmitter* removeAllListeners(Event event) override;
      virtual KuzzleEventEmitter* once(Event event, EventListener* listener) override;
      virtual int listenerCount(Event event) override;

  };
}

#endif
