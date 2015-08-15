# go-to-exercise
Exercise along with http://tour.golang.org

## Prepare

### Machine

>`# cat /etc/lsb-release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=14.04
DISTRIB_CODENAME=trusty
DISTRIB_DESCRIPTION="Ubuntu 14.04.3 LTS"
`

### Install go

get download link from [golang.org getting started documentation](http://golang.org/doc/install)

extract it into /usr/local

>`#tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`

set PATH environment variable before first use `go` command, or add your bashrc (`/home/your user name/.bashrc`)

>`export PATH=$PATH:/usr/local/go/bin`

## gotour is "A Tour of Go" 

### Get source code

* prepare work dir

more detail in section of __How to write Go code__ at [golang.org doc](http://golang.org/doc/)

1st create a go work dir, for example in my lab

>`#mkdir -p /opt/tfx/google-go`

2nd set GOPATH environment variable to work dir

>`#export GOPATH=/opt/tfx/google-go`

3rd ceate subdirs _src_, _bin_, _pkg_ under work dir, for example of _src_

>`#mkdir -p /opt/tfx/google-go/src`

* now get `go` command to get gotour into _src_ subdir

>`#go get golang.org/x/tour/gotour`

* gotour mirror repo in github [https://github.com/golang/tour](https://github.com/golang/tour) 

### use gotour

* simple run

>`#go run /opt/tfx/google-go/src/golang.org/x/tour/content/welcome/hello.go`

the target binary is built in `/tmp` dir

* build

just replace `run` command with 'build', and the binary will be generated in current dir

* install

same as `run` and `build` command, and the binary will be located in `bin` subdir
