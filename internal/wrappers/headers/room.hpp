#ifndef _ROOM_HPP_
#define _ROOM_HPP_

#include "listeners.hpp"
#include "exceptions.hpp"
#include "collection.hpp"
#include "core.hpp"
#include "room.hpp"

namespace kuzzleio {
    class Room {
        room *_room;
        SubscribeListener *_listener_instance;
        NotificationListener *_notification_listener_instance;

        Room(){}

        public:
            Room(Collection *collection, json_object *filters=NULL, room_options* options=NULL);
            Room(room* r);
            virtual ~Room();
            int count() Kuz_Throw_KuzzleException;
            SubscribeListener *getSubscribeListener();
            NotificationListener *getNotificationListener();
            Room* onDone(SubscribeListener* listener);
            Room* subscribe(NotificationListener* listener);
    };
}

#endif