#ifndef _DOCUMENT_HPP_
#define _DOCUMENT_HPP_

#include "listeners.hpp"
#include "exceptions.hpp"
#include "core.hpp"
#include "room.hpp"

#include <string>
#include <iostream>

namespace kuzzleio {    
    class Collection;

    class Document {
        Document(){};
        Collection *_collection;
        NotificationListener* _listener_instance;

        public:
            document *_document;
            Document(Collection *collection, const std::string& id="", json_object* content=NULL) Kuz_Throw_KuzzleException;
            virtual ~Document();
            std::string delete_(query_options* options=NULL) Kuz_Throw_KuzzleException;
            bool exists(query_options* options=NULL) Kuz_Throw_KuzzleException;
            bool publish(query_options* options=NULL) Kuz_Throw_KuzzleException;
            Document* refresh(query_options* options=NULL) Kuz_Throw_KuzzleException;
            Document* create(query_options* options=NULL) Kuz_Throw_KuzzleException;
            Document* setContent(json_object* content, bool replace=false);
            json_object* getContent();
            Room* subscribe(NotificationListener* listener, room_options* options=NULL);
            NotificationListener* getListener();
    };
}

#endif