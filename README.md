gen-gitrev
==========

Automatically generates source code to retrieve the HEAD revision of the git repository.

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
# build

```sh
$ make # => ./gen-gitrev
```
