#ifndef _COLLECTION_HPP_
#define _COLLECTION_HPP_

#include <iostream>

#include "core.hpp"

namespace kuzzleio {
    class Collection {
        public:
            collection* _collection;        
            Collection(Kuzzle* kuzzle, const std::string& collection, const std::string& index);
            virtual ~Collection();
    };
}

#endif