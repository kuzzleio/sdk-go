#include "room.hpp"
#include "collection.hpp"

namespace kuzzleio {
    Room::Room(Collection *c, json_object *filters, room_options* options) {
      this->_room = new room();
      room_new_room(this->_room, c->_collection, filters, options);
    }

    Room::~Room() {
      unregisterRoom(this->_room);
      delete(this->_room);
  }
}