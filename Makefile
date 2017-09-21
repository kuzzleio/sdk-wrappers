CC = gcc
CFLAGS = -fPIC -I/home/kblondel/Downloads/android-studio/jre/include/linux -I/home/kblondel/Downloads/android-studio/jre/include/ -I./headers
LDFLAGS = -L./
LIBS = -lgokcore
SRCS = kcore_wrap.c
OBJS = $(SRCS:.c=.o)
TARGET = libkcore.so

GOCC = /usr/local/bin/go
GOFLAGS = -buildmode=c-shared
GOSRC = ./cgo/kuzzle/
GOTARGET = libgokcore.so

SWIG = swig

all: java

kcore_wrap.o: kcore_wrap.c
	$(CC) -ggdb -c $< -o $@ $(CFLAGS) $(LDFLAGS) $(LIBS)

core:
	$(GOCC) build -o $(GOTARGET) $(GOFLAGS) $(GOSRC)

wrapper: $(OBJS)

object:
	gcc -ggdb -shared kcore_wrap.o -o libkcore.so $(LDFLAGS) $(LIBS)

swigjava:
	$(SWIG) -java -package io.kuzzle.sdk -outdir ./io/kuzzle/sdk -o kcore_wrap.c templates/java/core.i

java: 	core swigjava wrapper object

clean:
	rm -rf build *.class *.o *.h *.so io/kuzzle/sdk/*.java *.c *~ *.go

.PHONY: all java wrapper swigjava clean object core

.DEFAULT_GOAL := all
