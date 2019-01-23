
all: build

build:
	go build -v

mount:
	sudo mount -t proc proc /proc

install:
	mkdir /dockerc
	cp bysybox.tar /dockerc

busybox:
	# get the busybox.tar
	# docker run -d busybox top -b
	# docker export -o busybox.tar cid
	# sudo mkdir /root/busybox
	# tar -xvf busybox.tar -C /busybox

sh:
	sudo ./dockerc run -ti /bin/sh
