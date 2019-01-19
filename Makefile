
all: build

build:
	go build -v

mount:
	sudo mount -t proc proc /proc

sh:
	sudo ./dockerc run -ti /bin/sh
