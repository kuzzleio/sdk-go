#ifndef _KUZZLE_UTIL_HPP_
#define _KUZZLE_UTIL_HPP_

#include <stdarg.h>
#include <stdio.h>
#define TXT_COLOR_RESET "\e[0m"

#define TXT_COLOR_DEFAULT "\e[39m"
#define TXT_COLOR_BLACK "\e[30m"
#define TXT_COLOR_RED "\e[31m"
#define TXT_COLOR_GREEN "\e[32m"
#define TXT_COLOR_YELLOW "\e[33m"
#define TXT_COLOR_BLUE "\e[34m"
#define TXT_COLOR_MAGENTA "\e[35m"
#define TXT_COLOR_CYAN "\e[36m"
#define TXT_COLOR_LIGHTGREY "\e[37m"
#define TXT_COLOR_DARKGREY "\e[90m"

void kuz_log_sep();

void kuz_log_e(const char *filename, int linenumber, const char *fmt...);
void kuz_log_w(const char *filename, int linenumber, const char *fmt...);
void kuz_log_d(const char *filename, int linenumber, const char *fmt...);
void kuz_log_i(const char *filename, int linenumber, const char *fmt...);

#define K_LOG_D(fmt, ...)                                                      \
  do {                                                                         \
    kuz_log_d(__FILE__, __LINE__, fmt, ##__VA_ARGS__);                         \
  } while (0)

#define K_LOG_W(fmt, ...)                                                      \
  do {                                                                         \
    kuz_log_w(__FILE__, __LINE__, fmt, ##__VA_ARGS__);                         \
  } while (0)

#define K_LOG_E(fmt, ...)                                                      \
  do {                                                                         \
    kuz_log_e(__FILE__, __LINE__, fmt, ##__VA_ARGS__);                         \
  } while (0)

#define K_LOG_I(fmt, ...)                                                      \
  do {                                                                         \
    kuz_log_i(__FILE__, __LINE__, fmt, ##__VA_ARGS__);                         \
  } while (0)

std::string get_login_creds(const std::string &username,
                            const std::string &password);
void kuzzle_user_create(kuzzleio::Kuzzle *kuzzle, const std::string &user_id,
                        const std::string &username,
                        const std::string &password);

#endif
