cmake_minimum_required(VERSION 3.0.0)
project(KuzzleSKDTests)

add_subdirectory(3rdparty/json_spirit)

#
# Boost
#
set(BOOST_MIN_VERSION "1.46")

set(Boost_USE_STATIC_RUNTIME OFF)
set(CUKE_CORE_BOOST_LIBS thread system regex date_time program_options filesystem)
if(NOT CUKE_DISABLE_BOOST_TEST)
    # "An external test runner utility is required to link with dynamic library" (Boost User's Guide)
    set(Boost_USE_STATIC_LIBS OFF)
    set(CMAKE_CXX_FLAGS "-DBOOST_TEST_DYN_LINK ${CMAKE_CXX_FLAGS}")
    find_package(Boost ${BOOST_MIN_VERSION} COMPONENTS unit_test_framework)
endif()

if(CUKE_USE_STATIC_BOOST)
    set(Boost_USE_STATIC_LIBS ON)
    find_package(Boost ${BOOST_MIN_VERSION} COMPONENTS ${CUKE_CORE_BOOST_LIBS} REQUIRED)

    if(NOT MSVC)
        find_package(Threads)
        set(CUKE_EXTRA_LIBRARIES ${CUKE_EXTRA_LIBRARIES} ${CMAKE_THREAD_LIBS_INIT})
    endif()
else()
    set(CMAKE_CXX_FLAGS "-DBOOST_ALL_DYN_LINK ${CMAKE_CXX_FLAGS}")
    set(Boost_USE_STATIC_LIBS OFF)
    find_package(Boost ${BOOST_MIN_VERSION} COMPONENTS ${CUKE_CORE_BOOST_LIBS} REQUIRED)
endif()

include_directories(${Boost_INCLUDE_DIRS})
set(CUKE_EXTRA_LIBRARIES ${CUKE_EXTRA_LIBRARIES} ${Boost_THREAD_LIBRARY} ${Boost_SYSTEM_LIBRARY} ${Boost_REGEX_LIBRARY} ${Boost_DATE_TIME_LIBRARY} ${Boost_PROGRAM_OPTIONS_LIBRARY} ${Boost_FILESYSTEM_LIBRARY})


if(Boost_UNIT_TEST_FRAMEWORK_FOUND)
    link_directories("/home/etrousset/test/test_cucumber-cpp/cucumber-cpp/build/src")
    link_directories("/home/etrousset/go/src/github.com/kuzzleio/sdk-go/internal/wrappers/build/cpp")
    include_directories("/home/etrousset/test/test_cucumber-cpp/cucumber-cpp/include")
    include_directories("/home/etrousset/go/src/github.com/kuzzleio/sdk-go/internal/wrappers/headers")
    include_directories("/home/etrousset/go/src/github.com/kuzzleio/sdk-go/internal/wrappers/build/c")
    add_executable(KuzzleSDKStepDefs kuzzle-sdk-steps kuzzle_utils)
    target_link_libraries(KuzzleSDKStepDefs cpp cucumber-cpp  json_spirit ${CUKE_EXTRA_LIBRARIES} ${CUKE_LIBRARIES} ${Boost_UNIT_TEST_FRAMEWORK_LIBRARY})
endif()

