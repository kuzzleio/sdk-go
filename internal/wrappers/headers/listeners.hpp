#ifndef _LISTENERS_HPP
#define _LISTENERS_HPP

#include "kuzzle.hpp"

namespace kuzzleio {
    class NotificationListener {
        public:
            virtual ~NotificationListener(){};
            virtual void onMessage(notification_result*) const = 0;
    };

    class EventListener {
        public:
            virtual ~EventListener(){};
            virtual void trigger(json_object*, json_object*);
    };
}

#endif