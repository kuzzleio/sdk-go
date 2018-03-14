#ifndef _DOCUMENT_HPP_
#define _DOCUMENT_HPP_

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {

    class Document {
        document* _document;
        Document();

        public:
            Document(Kuzzle* kuzzle);
            Document(Kuzzle* kuzzle, document *document);
            virtual ~Document();
            int count_(const std::string& index, const std::string& collection, const std::string& body) Kuz_Throw_KuzzleException;
            bool exists(const std::string& index, const std::string& collection, const std::string& id) Kuz_Throw_KuzzleException;
            std::string create(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string createOrReplace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string delete_(const std::string& index, const std::string& collection, const std::string& id, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::vector<std::string> deleteByQuery(const std::string& index, const std::string& collection, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string get(const std::string& index, const std::string& collection, const std::string& id) Kuz_Throw_KuzzleException;
            std::string replace(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string update(const std::string& index, const std::string& collection, const std::string& id, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            bool validate(const std::string& index, const std::string& collection, const std::string& body) Kuz_Throw_KuzzleException;
            search_result* search(const std::string& index, const std::string& collection, const std::string& body, search_options *opts) Kuz_Throw_KuzzleException;
            std::string mCreate(const std::string& index, const std::string& collection, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string mCreateOrReplace(const std::string& index, const std::string& collection, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::vector<std::string> mDelete(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string mGet(const std::string& index, const std::string& collection, const std::vector<std::string>& ids, bool includeTrash) Kuz_Throw_KuzzleException;
            std::string mReplace(const std::string& index, const std::string& collection, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
            std::string mUpdate(const std::string& index, const std::string& collection, const std::string& body, document_options *doc_ops) Kuz_Throw_KuzzleException;
    };
}

#endif
