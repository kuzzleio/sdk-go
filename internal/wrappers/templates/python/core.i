%rename(TokenValidity) token_validity;
%rename(AckResponse) ack_response;
%rename(queueTTL) queue_ttl;
%rename(Options, match="class") options;
%rename(QueryOptions) query_options;
%rename(JsonObject) json_object;
%rename(JsonResult) json_result;
%rename(LoginResult) login_result;
%rename(BoolResult) bool_result;
%rename(Statistics) statistics;
%rename(AllStatisticsResult) all_statistics_result;
%rename(StatisticsResult) statistics_result;
%rename(CollectionsList) collection_entry;
%rename(CollectionsListResult) collection_entry_result;
%rename(StringArrayResult) string_array_result;
%rename(KuzzleResponse) kuzzle_response;
%rename(KuzzleRequest) kuzzle_request;
%rename(ShardsResult) shards_result;
%rename(DateResult) date_result;
%rename(UserData) user_data;
%rename(User, match="class") user;
%rename(RoomOptions) room_options;
%rename(SearchFilters) search_filters;
%rename(SearchResult) search_result;
%rename(NotificationResult) notification_result;
%rename(NotificationContent) notification_content;
%rename(SubscribeToSelf) subscribe_to_self;

%ignore *::error;
%ignore *::status;
%ignore *::stack;

%feature("director") NotificationListener;
%feature("director") EventListener;
%feature("director") SubscribeListener;

%{
#include "kuzzle.cpp"
#include "collection.cpp"
#include "auth.cpp"
#include "index.cpp"
#include "server.cpp"
#include "document.cpp"
#include "realtime.cpp"

#define SWIG_FILE_WITH_INIT
%}

%include "stl.i"
%include "../../kcore.i"

%extend options {
    options() {
        options *o = kuzzle_new_options();
        return o;
    }

    ~options() {
        free($self);
    }
}

%include "kuzzle.cpp"
%include "collection.cpp"
%include "document.cpp"
%include "realtime.cpp"
%include "auth.cpp"
%include "index.cpp"
%include "server.cpp"

// %{
// #define SWIG_FILE_WITH_INIT
// #include <iostream>
// #include <sstream>

// #include "kuzzle.hpp"
// %}

// %exception kuzzleio::Kuzzle::now {
//   try {
//     $action
//   } catch (...) {
//     std::cout << "BBlah";
//     PyObject *o = SWIG_NewPointerObj(new kuzzleio::ForbiddenError, SWIGTYPE_p_kuzzleio__ForbiddenError, SWIG_POINTER_OWN);
//     SWIG_Python_Raise(o, (char *)"MyException", SWIGTYPE_p_kuzzleio__ForbiddenError);
//     SWIG_fail;
//   }
// }

// %include "stl.i"
// %include "kcore.i"

// %{
// #include "core.cpp"
// %}

// namespace kuzzleio {
//   %exceptionclass KuzzleError;
//   %extend KuzzleError {
//     std::string __str__() const {
//       std::ostringstream s;
//       s << "[" << $self->status << "] " << $self->what();
//       if (!$self->stack.empty()) {
//         s << "\n" << $self->stack;
//       }
//       return s.str();
//     }
//   }
//   %exceptionclass BadRequestError;
//   %exceptionclass ForbiddenError;

//   %typemap(throws) BadRequestError "(void)$1; throw;";
// }

// /*
// %exception kuzzleio::Kuzzle::now {
//   try {
//     $action
//   } catch (...) {
//     std::cout << "Blah";
//     PyObject *o = SWIG_NewPointerObj(new kuzzleio::ForbiddenError, SWIGTYPE_p_kuzzleio__ForbiddenError, SWIG_POINTER_OWN);
//     SWIG_Python_Raise(o, (char *)"MyException", SWIGTYPE_p_kuzzleio__ForbiddenError);
//     SWIG_fail;
//   }
// }
// */

// %include "core.cpp"

// //%ignore kuzzle;
// %rename(QueryOptions) query_options;
