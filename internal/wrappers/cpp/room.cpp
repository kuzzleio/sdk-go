#include "room.hpp"

namespace kuzzleio {
    Room::Room() {
      this->_kuzzle = new kuzzle();
      kuzzle_new_kuzzle(this->_kuzzle, const_cast<char*>(host.c_str()), (char*)"websocket", opts);
    }
}