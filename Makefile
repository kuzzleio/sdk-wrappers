CC = gcc
CFLAGS = -fPIC -I$(PWD)/headers -I${JAVA_HOME}/include -I${JAVA_HOME}/include/linux
LDFLAGS = -L./
LIBS = -lgokcore -ljson-c
SRCS = kcore_wrap.c
OBJS = $(SRCS:.c=.o)
TARGET = libkcore.so

GOROOT ?= /usr/local
GOCC = $(GOROOT)/bin/go
GOFLAGS = -buildmode=c-shared
GOSRC = ./cgo/kuzzle/
GOTARGET = libgokcore.so

SWIG = swig

all: java

kcore_wrap.o: kcore_wrap.c
	$(CC) -ggdb -c $< -o $@ $(CFLAGS) $(LDFLAGS) $(LIBS)

core:
ifeq ($(wildcard $(GOCC)),)
	$(error "Unable to find go compiler")
endif
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)

wrapper: $(OBJS)

object:
	gcc -ggdb -shared kcore_wrap.o -o libkcore.so $(LDFLAGS) $(LIBS)

swigjava:
	$(SWIG) -java -package io.kuzzle.sdk -outdir ./io/kuzzle/sdk -o kcore_wrap.c -I/usr/local/include templates/java/core.i

java: 	core swigjava wrapper object

clean:
	rm -rf build *.class *.o *.h *.so io/kuzzle/sdk/*.java io/kuzzle/sdk/*.class *.c *~ *.go

.PHONY: all java wrapper swigjava clean object core

.DEFAULT_GOAL := all
