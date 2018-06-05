ROOT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ifeq ($(OS),Windows_NT)
    STATICLIB := .lib
    DYNLIB := .dll
    GOROOT ?= C:/Go
    GOCC ?= $(GOROOT)bin\go
    SEP = \\
    RM = del /Q /F /S
    RRM = rmdir /S /Q
    MV = rename
    MKDIR = mkdir
    CMDSEP=&
    ROOT_DIR_CLEAN = $(subst /,\,$(ROOT_DIR))
    LIB_PREFIX = 
else
    STATICLIB := .a
    DYNLIB := .so
    GOROOT ?= /usr/local/go
    GOCC ?= $(GOROOT)/bin/go
    SEP = \
    RM = rm -f
    RRM = rm -f -r
    MV = mv -f
    MKDIR = mkdir -p
    CMDSEP=;
    ROOT_DIR_CLEAN = $(ROOT_DIR)
    LIB_PREFIX = lib
endif
PATHSEP = $(strip $(SEP))
ROOTOUTDIR = $(ROOT_DIR)build
JAVA_HOME ?= /usr/local
PYTHON_INCLUDE_DIR ?= `pkg-config --cflags python3`
GOFLAGS = -buildmode=c-archive
GOSRC = .$(PATHSEP)cgo$(PATHSEP)kuzzle$(PATHSEP)
GOTARGETDIR = $(ROOTOUTDIR)/c
GOTARGET = $(GOTARGETDIR)$(PATHSEP)$(LIB_PREFIX)kuzzlesdk$(STATICLIB)
GOTARGETSO = $(GOTARGETDIR)$(PATHSEP)$(LIB_PREFIX)kuzzlesdk$(DYNLIB)
CPPDIR = $(ROOT_DIR)cpp
HEADERSDIR = $(ROOT_DIR)headers
CPP = $(CXX)
CPPFLAGS = -g -fPIC
INCS = -I$(HEADERSDIR) -I$(CPPDIR) -I$(ROOT_DIR)templates -I$(GOTARGETDIR)
LDFLAGS = -L$(GOTARGETDIR)
LIBS = -lkuzzlesdk
SRCS = kcore_wrap.cxx
OBJS = $(SRCS:.cxx=.o)
JAVAINCLUDE = -I$(JAVA_HOME)/include -I$(JAVA_HOME)/include/linux
PYTHONINCLUDE = $(PYTHON_INCLUDE_DIR) -I$(ROOT_DIR)templates/python
CPP_SDK_SRCS = kuzzle.cpp \
                    document.cpp \
                    collection.cpp \
                    realtime.cpp \
                    auth.cpp \
                    server.cpp \
                    index.cpp
CPPSRCS = $(CPP_SDK_SRCS) \
                    kcore_wrap.cxx
CPPSDK = $(CPP_SDK_SRCS:.cpp=.o)
SWIG = swig
all: python java csharp
kcore_wrap.o: kcore_wrap.cxx
    $(CPP) -c $< -o $@ $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(LANGINCLUDE)
%.o: cpp/%.cpp
    $(CPP) -fPIC -c $< -o $(OUTDIR)/$@ $(CPPFLAGS) -I./headers -I./build/c/ -L./build/c -lkuzzlesdk
core_java:
    cd ..$(PATHSEP)..$(PATHSEP) && go get .$(PATHSEP)... $(CMDSEP) cd internal$(PATHSEP)wrappers
    $(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)
    $(GOCC) build -o $(GOTARGETSO) $(GOFLAGS) $(GOSRC)
ifeq ($(OS),Windows_NT)
    $(MV) $(subst /,\,$(GOTARGETDIR)$(PATHSEP)$(LIB_PREFIX)kuzzlesdk.h) kuzzle.h
else
    $(MV) $(GOTARGETDIR)$(PATHSEP)$(LIB_PREFIX)kuzzlesdk.h $(GOTARGETDIR)$(PATHSEP)kuzzle.h
endif
core_python: core_java
core_csharp: core_java
wrapper_java: $(OBJS)
wrapper_python: $(OBJS)
wrapper_csharp: $(OBJS)
java_object:
    $(CPP) -shared kcore_wrap.o -o $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-java$(DYNLIB) $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(JAVAINCLUDE)
    strip $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-java$(DYNLIB)
python_object:
    $(CPP) -shared kcore_wrap.o -o $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-python$(DYNLIB) $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) $(PYTHONINCLUDE)
    strip $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-python$(DYNLIB)
csharp_object:
ifeq ($(OS),Windows_NT)
    $(CPP) -shared kcore_wrap.o -o $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-csharp$(DYNLIB) $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS) -lWinmm -lntdll -lWS2_32
else
    $(CPP) -shared kcore_wrap.o -o $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-csharp$(DYNLIB) $(CPPFLAGS) $(LDFLAGS) $(LIBS) $(INCS)
endif
    strip $(OUTDIR)/$(LIB_PREFIX)kuzzle-wrapper-csharp$(DYNLIB)
swig:
    $(SWIG) -Wall -c++ -$(LANGUAGE) $(SWIGOPTS) -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(LANGINCLUDE) templates/$(LANGUAGE)/core.i
swig_java:
    $(SWIG) -Wall -c++ -java -package io.kuzzle.sdk -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(JAVAINCLUDE) templates/java/core.i
swig_python:
    $(SWIG) -Wall -c++ -python -py3 -outdir $(OUTDIR) -o $(SRCS) $(INCS) $(PYTHONINCLUDE) templates/python/core.i
swig_csharp:
    $(SWIG) -Wall -c++ -csharp -namespace Kuzzleio -dllimport kuzzle-wrapper-csharp -outdir $(OUTDIR) -o $(SRCS) $(INCS)  templates/csharp/core.i
makedir:
ifeq ($(OS),Windows_NT)
    @if not exist $(subst /,\,$(OUTDIR)) mkdir $(subst /,\,$(OUTDIR))
    @if not exist $(ROOT_DIR_CLEAN) mkdir $(ROOT_DIR_CLEAN)
else
    mkdir $(ROOT_DIR_CLEAN)
    mkdir $(OUTDIR)
endif
java: LANGINCLUDE = $(JAVAINCLUDE)
java: OUTDIR = $(ROOTOUTDIR)/java/io/kuzzle/sdk
java: makedir core_java swig_java wrapper_java java_object
#rm -rf $(OUTDIR)/kcore_wrap.o $(OUTDIR)/libkuzzlesdk.a $(OUTDIR)/libkuzzlesdk.so
    taskset -c 1 $(JAVA_HOME)/bin/javac $(OUTDIR)/*.java
    mkdir $(ROOTOUTDIR)/io/kuzzle/sdk/resources
    mv build/java/io/kuzzle/sdk/$(LIB_PREFIX)kuzzle-wrapper-java.so build/io/kuzzle/sdk/resources
    mkdir $(ROOTOUTDIR)/java/src/main/java  
    ln -sf $(ROOTOUTDIR)/java/io/kuzzle/sdk/* $(ROOTOUTDIR)/java/src/main/java/
    cd build/java && taskset -c 1 sh gradlew sourcesJar jar javadocJar && cd -
    rm -f $(OUTDIR)/*.java $(OUTDIR)/*.class
    cp -p $(GOTARGET) $(OUTDIR)
    cp -p $(GOTARGETSO) $(OUTDIR)
csharp: OUTDIR = $(ROOTOUTDIR)/csharp
csharp: makedir core_csharp swig_csharp wrapper_csharp csharp_object
    mcs -t:library $(OUTDIR)/*.cs -sdk:2 -out:$(OUTDIR)/$(LIB_PREFIX)kuzzle-csharp$(DYNLIB)
    $(RM) "$(subst /,\,$(OUTDIR))$(PATHSEP)*.cs"
python: LANGINCLUDE = $(PYTHONINCLUDE)
python: OUTDIR = $(ROOTOUTDIR)/python
python: CC = g++
python: CFLAGS = -fPIC
python: makedir core_python swig_python wrapper_python python_object
    cp $(ROOT_DIR)templates/python/*.py $(OUTDIR)
    cp $(ROOT_DIR)$(SRCS) $(OUTDIR)/
    cp $(ROOT_DIR)kcore_wrap.h $(OUTDIR)/
    python3 $(OUTDIR)/setup.py build_ext -I $(HEADERSDIR):$(CPPDIR):$(OUTDIR):$(GOTARGETDIR):$(OUTDIR)/templates/python -L $(OUTDIR):$(GOTARGETDIR) -R $(GOTARGETDIR) -l kuzzlesdk  -b $(OUTDIR) -t $(OUTDIR)/tmp
    cp -p $(GOTARGET) $(OUTDIR)
    cp -p $(GOTARGETSO) $(OUTDIR)
cpp: OUTDIR = $(ROOTOUTDIR)/cpp
cpp: makedir $(CPPSDK)
cpp: LANGUAGE = c++
cpp: core
        ar rvs $(OUTDIR)/libcpp$(STATICLIB) $(OUTDIR)/*.o
        g++ -shared -fPIC -o $(OUTDIR)/libcpp$(DYNLIB) -Wl,--whole-archive $(OUTDIR)/libcpp$(STATICLIB) build/c/$(LIB_PREFIX)kuzzlesdk$(STATICLIB) -Wl,--no-whole-archive
clean:
    $(RM) kcore_wrap.cxx
    $(RM) kcore_wrap.h 
    $(RRM) build$(PATHSEP)c build$(PATHSEP)java$(PATHSEP)src build$(PATHSEP)java$(PATHSEP)io build$(PATHSEP)java$(PATHSEP)build build$(PATHSEP)csharp build$(PATHSEP)python
    $(RRM) build$(PATHSEP)java
.PHONY: all java csharp python wrapper swig clean object core
.DEFAULT_GOAL := all
