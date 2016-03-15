all:
	go build .

install: all
	go install .


.PHONY:
	all install
