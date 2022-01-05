# nosigsegv

Remove `nil` pointer problems from your go program!

## Usage

If you are running a recent Linux version, you need to run the following command **before** using `nosigsegv`.

```sh
sudo sysctl -w vm.mmap_min_addr=0
```

Now `nil` pointers are not going to crash your program anymore.

## Example

Add the following line at the top of your go program:

```go
import _ "github.com/thebabush/nosigsegv"
```

For a complete example, see [cmd/nosigsegv/main.go](cmd/nosigsegv/main.go).

Test it yourself:

```sh
$ cd ./cmd/nosigsegv
$ go build
$ ./nosigsegv
initialization ok!
testing...
* <nil>  =  123
```

## Limitations

- Only works on Linux
- Does not support pointer values other than `nil`, unless there's already something allocated at that address
