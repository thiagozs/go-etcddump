# go-etcddump

A simple tool used to dump / restore etcd KV.

### Install

```sh
go get github.com/thiagozs/go-etcddump
```

Or download a compiled version at [release](https://github.com/thiagozs/go-etcddump/releases) page

### Usage

```sh
# help
etcddump -h

# dump
etcddump dump \
	--address=127.0.0.1:2379 \
	--prefix="/micro/config/jm" \
	--output=test.out

# restore
etcddump restore \
	--address=127.0.0.1:2379 \
	--file=test.out
```

## Versioning and license

We use SemVer for versioning. You can see the versions available by checking the tags on this repository.

For more details about our license model, please take a look at the LICENSE file

---

2022, thiagozs
