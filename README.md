gen-gitrev
==========

Automatically generates source code to retrieve the HEAD revision of the git repository.

# Download

- Release page: [https://github.com/fkei/gen-gitrev/releases](https://github.com/fkei/gen-gitrev/releases)
- Download binary file(version 0.1): [https://github.com/fkei/gen-gitrev/releases/download/0.1/gen-gitrev](https://github.com/fkei/gen-gitrev/releases/download/0.1/gen-gitrev)

# install

```sh
$ go get github.com/fkei/gen-gitrev
```

# Use

```sh
$ gen-gitrev -out /tmp/hoge_gen.go -pkgname main
Output generate file: /tmp/hoge_gen.go
	pacakge name: main
	revision: aa41f29462e685ebb199a205d97f938bc3789489

$ cat /tmp/hoge_gen.go
package main

func GetRevision() (string){
	return "aa41f29462e685ebb199a205d97f938bc3789489"
}
```

# Options

```
$ gen-gitrev -h

Usage of gen-gitrev:
   gen-gitrev [OPTIONS] ARGS...

Options  -out="/Users/fkei/repository/github/gen-gitrev/gitrev_gen.go": output file path
  -pkgname="main": package name
  
```

# go generate <= go1.4

```go
//go:generate gen-gitrev -pkgname gen -out ./gen/gitrev_gen.go
```

# build

```sh
$ make # => ./gen-gitrev
```
