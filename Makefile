.POSIX:

OS = $(shell uname -s)
ifeq ($(OS), Darwin)
  PREFIX = /usr/local
else
  PREFIX = /usr
endif
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

default: build

installdep:
	@go get

build: installdep
	go build
	cd examples/bricks ; go build $(ROOT_DIR)/examples/bricks_src/music.go
	cd examples/bricks ; go build $(ROOT_DIR)/examples/bricks_src/workspaces.go
	cd examples/bricks ; go build $(ROOT_DIR)/examples/bricks_src/battery.go
	mkdir -p ${XDG_CONFIG_HOME}/lego
	cp examples/legorc ${XDG_CONFIG_HOME}/lego/legorc
	mkdir -p ${XDG_CONFIG_HOME}/lego/bricks
	chmod +x -R examples/bricks
	for brick in examples/bricks/* ; do \
		cp -f $$brick  ${XDG_CONFIG_HOME}/lego/bricks; \
	done
	chmod +x -R examples/subscribe

install:
	cp lego /usr/bin/lego
	cp lego_signal /usr/bin/lego_signal
	cp kill_lego /usr/bin/kill_lego
	cp lego_refresh /usr/bin/lego_refresh
	chmod +x /usr/bin/lego
	chmod +x /usr/bin/lego_signal
	chmod +x /usr/bin/kill_lego
	chmod +x /usr/bin/lego_refresh
	chmod +x subscribe_cleanup
	touch /tmp/lego_refresh
	chmod 777 /tmp/lego_refresh
	cp -R examples/subscribe/* /usr/bin
