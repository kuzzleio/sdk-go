#include "collection.hpp"

namespace kuzzleio {
    Collection::Collection(Kuzzle* kuzzle, const std::string& col, const std::string& index) {
        _collection = new collection();
        kuzzle_new_collection(_collection, kuzzle->_kuzzle, const_cast<char*>(col.c_str()), const_cast<char*>(index.c_str()));
    }

    Collection::~Collection() {
        unregisterCollection(_collection);
        delete(_collection);
    }
}