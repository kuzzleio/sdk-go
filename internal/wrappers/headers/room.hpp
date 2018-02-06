#ifndef _ROOM_HPP_
#define _ROOM_HPP_

namespace kuzzleio {
    class Room {
        room *_room;
        SubscribeListener *_listener_instance;

        Room(){}

        public:
            Room(Collection *collection, json_object *filters=NULL, room_options* options=NULL);
            Room(room* r);
            virtual ~Room();
            int count() Kuz_Throw_KuzzleException;
            SubscribeListener *getSubscribeListener();
            Room* onDone(SubscribeListener* listener);
    };
}

#endif