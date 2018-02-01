#ifndef _EVENT_EMITTER_HPP_
#define _EVENT_EMITTER_HPP_

#include "kuzzle.hpp"
#include "listeners.hpp"

namespace kuzzleio {
  class KuzzleEventEmitter {
    virtual Kuzzle* addListener(Event e, EventListener* listener) const = 0;
    virtual Kuzzle* removeListener(Event e, EventListener* listener) const = 0;
    virtual Kuzzle* 
  };
}

#endif