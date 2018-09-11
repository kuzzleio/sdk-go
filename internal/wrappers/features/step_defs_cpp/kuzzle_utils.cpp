#include <boost/property_tree/json_parser.hpp>
#include <boost/property_tree/ptree.hpp>
#include <json_spirit/json_spirit_reader.h>
#include <json_spirit/json_spirit_writer.h>
#include <map>
#include <sstream>

#include <exceptions.hpp>
#include <kuzzle.hpp>
#include <kuzzlesdk.h>

#include "kuzzle_utils.h"

using boost::property_tree::ptree;
using boost::property_tree::read_json;
using boost::property_tree::write_json;

using kuzzleio::Kuzzle;
using kuzzleio::KuzzleException;
using std::cout;
using std::endl;
using std::string;

using json_spirit::Array;
using json_spirit::Object;
using json_spirit::Pair;
using json_spirit::Value;

std::string get_login_creds(const std::string &username,
                            const std::string &password)
{
  // Write json.
  ptree pt;
  pt.put("username", username);
  pt.put("password", password);
  std::ostringstream buf;
  write_json(buf, pt, false);
  return buf.str();
}

string get_createUser_body(const std::string &username,
                           const std::string &password)
{
  // Write json.

  Object body;
  Object content;
  Object creds;
  Object local;

  local.push_back(Pair("username", Value(username)));
  local.push_back(Pair("password", Value(password)));

  creds.push_back(Pair("local", Value(local)));

  Array profileIds;
  profileIds.push_back("default");

  content.push_back(json_spirit::Pair("profileIds", Value(profileIds)));
  content.push_back(json_spirit::Pair("name", Value(username)));

  body.push_back(Pair("content", Value(content)));
  body.push_back(Pair("credentials", Value(creds)));

  std::ostringstream buf;
  json_spirit::write_formatted(body, buf);

  return buf.str();
}

bool kuzzle_user_exists(Kuzzle *kuzzle, const string &user_id)
{
  bool user_exists = false;
  try
  {
    kuzzle_request req = {0};
    req.controller = "security";
    req.action = "getUser";
    req.id = user_id.c_str();
    kuzzle->query(&req);
    user_exists = true;
  }
  catch (kuzzleio::NotFoundException e)
  {
    user_exists = false;
  }
  catch (KuzzleException e)
  {
    K_LOG_W(e.getMessage().c_str());
  }
  return user_exists;
}

void kuzzle_user_delete(Kuzzle *kuzzle, const string &user_id)
{
  K_LOG_D("Deleting user with ID: '%s'", user_id.c_str());
  try
  {
    kuzzle_request req = {0};
    req.controller = "security";
    req.action = "deleteUser";
    req.id = user_id.c_str();

    query_options options = {0};
    options.refresh = const_cast<char*>("wait_for");
    options.volatiles = const_cast<char*>("{}");
    kuzzle->query(
        &req, &options); // TODO: test if we can delete with options

    K_LOG_D("Deleted user \"%s\"", user_id.c_str());
  }
  catch (KuzzleException e)
  {
    K_LOG_E("Failed to delete user \"%s\"", user_id.c_str());
  }
}

void kuzzle_credentials_delete(Kuzzle *kuzzle, const string &strategy,
                               const string &user_id)
{
  try
  {
    kuzzle_request req = {0};
    req.controller = "security";
    req.action = "deleteCredentials";
    req.strategy = "local";
    req.id = user_id.c_str();

    kuzzle->query(&req);

    K_LOG_D("Deleted '%s' credentials for userId '%s'", strategy.c_str(),
            user_id.c_str());
  }
  catch (KuzzleException e)
  {
    K_LOG_E("Failed to delete '%s' credentials for userId '%s'",
            strategy.c_str(), user_id.c_str());
    K_LOG_E(e.getMessage().c_str());
  }
}

void kuzzle_user_create(Kuzzle *kuzzle, const string &user_id,
                        const string &username, const string &password)
{

  kuzzle_request req = {0};
  req.controller = "security";
  req.action = "createUser";
  req.strategy = "local";
  req.id = user_id.c_str();
  string body = get_createUser_body(username, password);
  req.body = body.c_str();

 // if (kuzzle_user_exists(kuzzle, user_id.c_str()))
  {
    K_LOG_W("An user with id: '%s' already exists, deleteting it...",
            user_id.c_str());
    kuzzle_credentials_delete(kuzzle, "local", user_id);
    kuzzle_user_delete(kuzzle, user_id);
  }

    // kuzzle_credentials_delete(kuzzle, "local", user_id);

  K_LOG_D("Creating user with id: '%s' and 'local' creds: %s:%s",
          user_id.c_str(), username.c_str(), password.c_str());
  K_LOG_D("Req body: %s", req.body);
  try
  {
    kuzzle_response *resp = kuzzle->query(&req);
    K_LOG_D("createUser ended with status: %d", resp->status);
  }
  catch (KuzzleException e)
  {
    K_LOG_E("Status (%d): %s", e.status, e.getMessage().c_str());
    if (kuzzle_user_exists(kuzzle, user_id.c_str())) {
      K_LOG_W("But user seems to exist anyway?????");
    }
  }
}

void kuz_log_sep()
{
  cout << "\x1b(0\x6c\x1b(B";
  for (int i = 0; i < 79; i++)
  {
    cout << "\x1b(0\x71\x1b(B";
  }
  cout << endl;
}

void kuz_log_e(const char *filename, int linenumber, const char *fmt...)
{
  va_list args;
  va_start(args, fmt);

  cout << "\x1b(0\x78\x1b(B" << TXT_COLOR_RED << "E:";
  cout << basename(filename) << ":" << linenumber << ":";
  vprintf(fmt, args);
  std::cout << TXT_COLOR_RESET << std::endl;

  va_end(args);
}

void kuz_log_w(const char *filename, int linenumber, const char *fmt...)
{
  va_list args;
  va_start(args, fmt);
  cout << "\x1b(0\x78\x1b(B" << TXT_COLOR_YELLOW << "W:";
  cout << basename(filename) << ":" << linenumber << ":";
  vprintf(fmt, args);
  cout << TXT_COLOR_RESET << endl;

  va_end(args);
}

void kuz_log_d(const char *filename, int linenumber, const char *fmt...)
{
  va_list args;
  va_start(args, fmt);

  cout << "\x1b(0\x78\x1b(B" << TXT_COLOR_DEFAULT << "D:";
  cout << basename(filename) << ":" << linenumber << ":";
  vprintf(fmt, args);
  cout << TXT_COLOR_RESET << endl;

  va_end(args);
}

void kuz_log_i(const char *filename, int linenumber, const char *fmt...)
{
  va_list args;
  va_start(args, fmt);

  cout << "\x1b(0\x78\x1b(B" << TXT_COLOR_BLUE << "I:";
  cout << basename(filename) << ":" << linenumber << ":";
  vprintf(fmt, args);
  cout << TXT_COLOR_RESET << endl;

  va_end(args);
}
