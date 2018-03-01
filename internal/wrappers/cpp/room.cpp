#include "room.hpp"
#include "collection.hpp"

namespace kuzzleio {
    Room::Room(room* r, SubscribeListener* subscribe_listener, NotificationListener* notification_listener) {
      this->_room = r;
      this->_listener_instance = subscribe_listener;
      this->_notification_listener_instance = notification_listener;
    }

    Room::Room(Collection *collection, json_object *filters, room_options* options) {
      this->_room = new room();
      room_new_room(this->_room, collection->_collection, filters, options);
    }

    Room::~Room() {
      unregisterRoom(this->_room);
      delete(this->_room);
    }

    void call_cb(room_result* res, void* data) {
        ((Room*)data)->getSubscribeListener()->onSubscribe(res);
    }

    void notify(notification_result* res, void* data) {
        ((Room*)data)->getNotificationListener()->onMessage(res);
    }

    SubscribeListener* Room::getSubscribeListener() {
        return _listener_instance;
    }

    NotificationListener* Room::getNotificationListener() {
        return _notification_listener_instance;
    }

    Room* Room::onDone(SubscribeListener *listener) {
        _listener_instance = listener;
        room_on_done(_room, &call_cb, this);
    }

    Room* Room::subscribe(NotificationListener* listener) {
        _notification_listener_instance = listener;
        room_subscribe(_room, &notify, this);
    }

    void Room::unsubscribe() Kuz_Throw_KuzzleException {
        void_result *r = room_unsubscribe(_room);
        if (r && r->error != NULL)
            throwExceptionFromStatus(r);
        _notification_listener_instance = NULL;
        _listener_instance = NULL;
        delete(r);
    }
}