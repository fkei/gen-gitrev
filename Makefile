__out=gitrev2go

all: build

clean:
	rm -f $(__out)

build:
	go build -o $(__out) main.go

