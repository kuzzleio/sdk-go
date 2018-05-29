cmake_minimum_required(VERSION 3.0.0)
project(KuzzleSKDTests)


include(ExternalProject)

ExternalProject_Add(
    project_cucumber_cpp
    GIT_REPOSITORY https://github.com/cucumber/cucumber-cpp
    INSTALL_COMMAND ""

    CMAKE_CACHE_ARGS
      -DCUKE_DISABLE_QT:BOOL=ON 
      -DCUKE_ENABLE_EXAMPLES:BOOL=OFF 
      -DCUKE_DISABLE_GTEST:BOOL=ON
      -DCUKE_DISABLE_UNIT_TESTS:BOOL=ON
      -DCUKE_DISABLE_E2E_TESTS:BOOL=ON
)

# Library
ExternalProject_Get_Property(project_cucumber_cpp binary_dir)
set(CUCUMBER_CPP_LIBRARY_PATH ${binary_dir}/${CMAKE_FIND_SOURCE_PREFIXES}/src/libcucumber-cpp.a)
list( APPEND CUCUMBER_CPP_INCLUDE_PATH "${binary_dir}/${CMAKE_FIND_SOURCE_PREFIXES}/src")
set(CUCUMBER_CPP_LIBRARY cucumber-cpp)


ExternalProject_Get_property(project_cucumber_cpp SOURCE_DIR)
list(APPEND CUCUMBER_CPP_INCLUDE_PATH "${SOURCE_DIR}/include")

message("cucumber-cpp: include path  = ${CUCUMBER_CPP_INCLUDE_PATH}")
message("cucumber-cpp: library parh = ${CUCUMBER_CPP_LIBRARY_PATH}")

add_library(cucumber-cpp STATIC IMPORTED) # or STATIC instead of SHARED
add_dependencies(cucumber-cpp project_cucumber_cpp)
set_target_properties(cucumber-cpp PROPERTIES
    IMPORTED_LOCATION ${CUCUMBER_CPP_LIBRARY_PATH}
    INTERFACE_INCLUDE_DIRECTORIES  
        "${CUCUMBER_CPP_INCLUDE_PATH}"
    DEPENDS project_cucumber_cpp
)

#
# JSon Spirit
#

add_subdirectory(3rdparty/json_spirit)

#
# Kuzzle SDK 
#

get_filename_component(KUZZLE_SDK_ROOT "${CMAKE_SOURCE_DIR}/../../" ABSOLUTE)
message("Kuzzle SDK root: ${KUZZLE_SDK_ROOT}")

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
    set (CMAKE_CXX_STANDARD 11)

    add_executable(KuzzleSDKStepDefs kuzzle-sdk-steps kuzzle_utils)
    add_dependencies(KuzzleSDKStepDefs project_cucumber_cpp)


    add_library(kuzzlesdk STATIC IMPORTED) # or STATIC instead of SHARED
    set_target_properties(kuzzlesdk PROPERTIES
    IMPORTED_LOCATION "${KUZZLE_SDK_ROOT}/build/c/libkuzzlesdk.a" 
    INTERFACE_INCLUDE_DIRECTORIES "${KUZZLE_SDK_ROOT}/build/c"
    )

    add_library(cpp STATIC IMPORTED) # or STATIC instead of SHARED
    set_target_properties(cpp PROPERTIES
    IMPORTED_LOCATION  "${KUZZLE_SDK_ROOT}/build/cpp/libcpp.a"
    INTERFACE_INCLUDE_DIRECTORIES "${KUZZLE_SDK_ROOT}/headers"
    )

    target_link_libraries(KuzzleSDKStepDefs cpp kuzzlesdk  pthread cucumber-cpp  json_spirit ${CUKE_EXTRA_LIBRARIES} ${CUKE_LIBRARIES} ${Boost_UNIT_TEST_FRAMEWORK_LIBRARY} )
else()
    message(FATAL_ERROR)
endif()

