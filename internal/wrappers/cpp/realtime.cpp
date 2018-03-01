#include "realtime.hpp"

namespace kuzzleio {
  Realtime::Realtime(Kuzzle *kuzzle) {
    _realtime = new realtime();
    kuzzle_new_realtime(_realtime, kuzzle->_kuzzle);
  }

  Realtime::~Realtime() {
    unregisterRealtime(_realtime);
    delete(_realtime);
  }

  int Realtime::count(const std::string& index, const std::string collection, const std::string roomId) Kuz_Throw_KuzzleException {
    int_result *r = kuzzle_realtime_count(_realtime, const_cast<char*>(index.c_str()), const_cast<char*>(collection.c_str()), const_cast<char*>(roomId.c_str()));
    if (r->error != NULL)
        throwExceptionFromStatus(r);
    int ret = r->result;
    delete(r);
    return ret;
  }
}