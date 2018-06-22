#ifndef _STEPS_HPP_
#define _STEPS_HPP_

#include <boost/test/unit_test.hpp>
#define EXPECT_EQ BOOST_CHECK_EQUAL
#include <cucumber-cpp/autodetect.hpp>
#include <iostream>

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

class CustomNotificationListener;

struct KuzzleCtx {
  Kuzzle* kuzzle = NULL;
  options kuzzle_options;

  string user_id;
  string index;
  string collection;
  string jwt;

  string room_id;

  user*                   currentUser        = NULL;
  json_spirit::Value_type customUserDataType = json_spirit::null_type;

  bool success;
  int hits;

  notification_result *notif_result = NULL;
  CustomNotificationListener *listener;
};

class CustomNotificationListener : public NotificationListener {
  public:
    virtual void onMessage(notification_result *res) const {
      ScenarioScope<KuzzleCtx> ctx;

      ctx->notif_result = res;
    }
};

#endif
