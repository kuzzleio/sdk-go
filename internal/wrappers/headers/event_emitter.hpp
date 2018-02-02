#ifndef _EVENT_EMITTER_HPP_
#define _EVENT_EMITTER_HPP_

#include "kuzzle.hpp"
#include "listeners.hpp"

namespace kuzzleio {
  class KuzzleEventEmitter {
    public:
      virtual KuzzleEventEmitter* addListener(Event e, EventListener* listener) = 0;
      virtual KuzzleEventEmitter* removeListener(Event e, EventListener* listener) = 0;
      virtual KuzzleEventEmitter* removeAllListeners(Event e) = 0;
      virtual KuzzleEventEmitter* once(Event e, EventListener* listener) = 0;
      virtual int listenerCount(Event e) = 0;
  };
}

#endif