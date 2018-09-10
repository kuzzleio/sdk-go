#ifndef _STEPS_HPP_
#define _STEPS_HPP_

#include <boost/test/unit_test.hpp>
#define EXPECT_EQ BOOST_CHECK_EQUAL
#include <cucumber-cpp/autodetect.hpp>
#include <cstdlib>
#include <iostream>
#include <functional>

#include "auth.hpp"
#include "collection.hpp"
#include "document.hpp"
#include "index.hpp"
#include "realtime.hpp"
#include "kuzzle.hpp"

#include "kuzzle_utils.h"

#include "json_spirit/json_spirit.h"

using cucumber::ScenarioScope;

using namespace kuzzleio;
using std::cout;
using std::endl;
using std::string;

struct KuzzleCtx {
  Kuzzle* kuzzle = NULL;
  options kuzzle_options;

  string user_id;
  string index;
  string collection;
  string jwt;
  string document_id;
  search_result *documents;

  string room_id;

  kuzzle_user*                   currentUser        = NULL;
  json_spirit::Value_type customUserDataType = json_spirit::null_type;

  // 1 mean success, 0 failure and -1 is base state
  int success = -1;
  string error_message;
  int hits = -1;
  string content;
  // 1 mean yes, 0 no and -1 is base state
  int partial_exception = -1;
  std::vector<string> string_array;

  notification_result *notif_result = NULL;
};

class CustomNotificationListener {
  private:
    CustomNotificationListener() {
      listener = [](const kuzzleio::notification_result* res) {
        ScenarioScope<KuzzleCtx> ctx;
        ctx->notif_result = const_cast<notification_result*>(res);
      };
    };
    static CustomNotificationListener* _singleton;
  public:
    NotificationListener listener;
    static CustomNotificationListener* getSingleton() {
      if (!_singleton) {
        _singleton = new CustomNotificationListener();
      }
      return _singleton;
    }
};

#endif
