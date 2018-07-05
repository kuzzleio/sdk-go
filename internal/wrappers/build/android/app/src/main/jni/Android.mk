# File: Android.mk
LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)
LOCAL_MODULE := kuzzlesdk
LOCAL_SRC_FILES := ../../../../../../build/c/libkuzzlesdk.so
include $(PREBUILT_SHARED_LIBRARY)


include $(CLEAR_VARS)
LOCAL_MODULE    := kuzzle
LOCAL_CPP_FEATURES := rtti exceptions stdio
LOCAL_CPPFLAGS += -I../../../../../../headers -I../../../../../../cpp/ -I../../../../../../build/c
LOCAL_SHARED_LIBRARIES := kuzzlesdk
LOCAL_SRC_FILES := ../../../../../../kcore_wrap.cxx
include $(BUILD_SHARED_LIBRARY)