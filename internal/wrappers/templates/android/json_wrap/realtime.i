%rename (_list) kuzzleio::Realtime::list(const std::string& index, const std::string& collection, const std::string& id, query_options *options);
%rename (_list) kuzzleio::Realtime::list(const std::string& index, const std::string& collection, const std::string& id);
%javamethodmodifiers kuzzleio::Realtime::list(const std::string& index, const std::string& collection, const std::string& id, query_options *options) "private";
%javamethodmodifiers kuzzleio::Realtime::list(const std::string& index, const std::string& collection, const std::string& id) "private";
