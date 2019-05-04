default: build

installdep:
	@go get

build: installdep
	mkdir -p ${XDG_CONFIG_HOME}/lego
	cp defaults/legorc ${XDG_CONFIG_HOME}/lego/legorc
	mkdir -p ${XDG_CONFIG_HOME}/lego/bricks
	mkdir -p ${XDG_CONFIG_HOME}/lego/subscribe
	cp -R defaults/bricks ${XDG_CONFIG_HOME}/lego
	chmod +x -R ${XDG_CONFIG_HOME}/lego/bricks
	cp -R defaults/subscribe ${XDG_CONFIG_HOME}/lego
	@go build

install:
	cp lego_signal /usr/bin/lego_signal
	cp kill_lego /usr/bin/kill_lego
	cp lego /usr/bin/lego
	touch /tmp/lego_refresh
	chmod 777 /tmp/lego_refresh
	chmod +x -R  defaults/subscribe
	cp -R defaults/subscribe/* /usr/bin
	cp lego_refresh /usr/bin
	chmod +x /usr/bin/lego_refresh
	chmod +x subscribe_cleanup
	chown root /tmp/lego_refresh
	chgrp root /tmp/lego_refresh

clean:
	sudo rm -f /usr/bin/lego_refresh
	sudo rm -f /usr/bin/lego_signal
	sudo rm -f /usr/bin/kill_lego
	./subscribe_cleanup
