#ifndef _ROOM_HPP_
#define _ROOM_HPP_

namespace kuzzleio {
    class Room {
        room *_room;
        Room(){}

        public:
            Room(Collection *c, json_object *filters=NULL, room_options* options=NULL);
            virtual ~Room();
    };
}

#endif