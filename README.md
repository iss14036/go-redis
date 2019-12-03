# Go Redis

## Problem Statement
Simple project implementation of redis pool size and routines in Go.

## Dependencies
In this time, I use mac as an os to running the system.

#### 1. Golang
First of all export some paths, and save them in your `.zshrc` or `.bashrc` files for easy use. Use sudo if you get error.

```
# Go development
```

```
export GOPATH="${HOME}/.go"
```

```
export GOROOT="$(brew --prefix golang)/libexec"
```

```
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
```

```
test -d "${GOPATH}" || mkdir "${GOPATH}"
```

And then, you can install the Go
```
brew install go

```

To setup Go mod, you can follow link [HERE](https://github.com/golang/go/wiki/Modules#how-to-install-and-activate-module-support).

#### 2. Redis

Well, well, well. I use mac in this develompenmt, if you do. You can type this command for installing the redis.
```
brew install redis
```

## Start the application

#### Start redis
You can type this command below for starting redis,
```
redis-server
```
Check either redis is already running or not
```
redis-cli ping
```

expected output will be `Pong`

#### Run the Go
And if redis works, you can run the golang by typing this following command :
```
go run main.go
```
