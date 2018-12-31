# GoLearnGo

I am pretty interested in [Kubernetes](https://kubernetes.io/) and a few other
Go based software like [Docker](https://www.docker.com/) and
[Terraform](https://www.terraform.io/). To be able to understand the workings
of these projects and possibly work on them, I have started to learn Go.

## Setup Steps

*Note* There's a simpler method at the bottom which doesn't install any Go
dependencies on the host - [Docker Method](#docker-method)

1. Install [gvm](https://github.com/moovweb/gvm)
```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

2. Install Go 1.11.4 (Latest as of this writing)
```bash
gvm install go1.4 -B
gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
gvm install go1.11.4
gvm use go1.11.4 --default
```

3. Create workspace
```bash
mkdir ~/workspace/code/go && cd $_
export GOPATH=$PWD
export PATH=$PATH:$(go env GOPATH)/bin
```

4. Write a "Hello World" program
```bash
mkdir -p $GOPATH/src/github.com/mbtamuli/GoLearnGo/hello
cat << EOF > helloworld.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
}
EOF
```

5. Run "Hello World"
```bash
go install github.com/mbtamuli/GoLearnGo/hello
hello
```

FYI, once you have installed go, which you can verify using `go version`, you
can simply run
```bash
go get github.com/mbtamuli/GoLearnGo/hello && hello
```

### Docker Method

1. Clone the repo
```
git clone https://github.com/mbtamuli/GoLearnGo.git
```

2. Run the following command to start the container
```
docker run --rm -ti -v "$PWD":/usr/src/GoLearnGo -w /usr/src/GoLearnGo golang:1.11-alpine /bin/sh
```

3. Now you can run the  programs using
```
go install GoLearnGo/calculator
```

## Resources

https://dave.cheney.net/resources-for-new-go-programmers

I think this is a pretty good enough curation of all the resources one will
ever need to "Get Started" with Go.

## Progress

1. I have started with https://tour.golang.org
