#include "index.hpp"

namespace kuzzleio {

    Index::Index(Kuzzle* kuzzle) Kuz_Throw_KuzzleException {
        _index = new index();
        kuzzle_new_index(_index, kuzzle->_kuzzle);
    }

    Index::~Index() {
        unregisterIndex(_index);
        delete(_index);
    }

    void Index::create(const std::string& index) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_index_create(_index, const_cast<char*>(index.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        delete(r);
    }

    void Index::delete_(const std::string& index) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_index_delete(_index, const_cast<char*>(index.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        delete(r);
    }

    string* Index::mDelete(string* indexes) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_index_mdelete(_index, const_cast<char**>(index.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        delete(r);
    }

    bool Index::exists(const std::string& index) Kuz_Throw_KuzzleException {
        bool_result *r = kuzzle_index_exists(_index, const_cast<char*>(index.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        bool ret = r->result;
        delete(r);
        return ret;
    }

    void Index::refresh(const std::string& index) Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_index_refresh(_index, const_cast<char*>(index.c_str()));
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        delete(r);
    }

    void Index::refreshInternal() Kuz_Throw_KuzzleException {
        void_result *r = kuzzle_index_refresh_internal(_index);
        if (r->error != NULL)
            throwExceptionFromStatus(r);
        delete(r);
    }

    void Index::setAutoRefresh(const std::string& index, bool autoRefresh) Kuz_Throw_KuzzleException {
      void_result *r = kuzzle_index_set_auto_refresh(_index, const_cast<char*>(index.c_str()), autoRefresh);
      if (r->error != NULL)
          throwExceptionFromStatus(r);
      delete(r);
    }
}
