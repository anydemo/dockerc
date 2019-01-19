
all: build

build:
	go build -v

mount:
	sudo mount -t proc proc /proc

busybox:
	# docker run -d busybox top -b
	# docker export -o busybox.tar cid
	# sudo mkdir /root/busybox
	# tar -xvf busybox.tar -C /busybox

sh:
	sudo ./dockerc run -ti /bin/sh
