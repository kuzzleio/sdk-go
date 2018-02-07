#ifndef _ROOM_HPP_
#define _ROOM_HPP_

#include "listeners.hpp"
#include "exceptions.hpp"
#include "core.hpp"
#include "room.hpp"

namespace kuzzleio {
    class Collection;

    class Room {
        room *_room;
        SubscribeListener *_listener_instance;
        NotificationListener *_notification_listener_instance;

        Room(){}

        public:
            Room(Collection *collection, json_object *filters=NULL, room_options* options=NULL);
            Room(room* r, SubscribeListener* subscribe_listener, NotificationListener* notification_listener);            
            virtual ~Room();
            int count() Kuz_Throw_KuzzleException;
            SubscribeListener *getSubscribeListener();
            NotificationListener *getNotificationListener();
            Room* onDone(SubscribeListener* listener);
            Room* subscribe(NotificationListener* listener);
            void unsubscribe() Kuz_Throw_KuzzleException;
    };
}

#endif