# Package env

`import "ztaylor.me/env"`

Package env provides runtime environment, using CLI flags, and/or config file, and/or environment variables

## Environment config

```go
func main() {
	// most common case
	env := env.Global() // overwrite package name why not
	server := server.New(env) // clients use env
	...
	// multiple env #1
	e1, e2 := env.NewService(), env.NewService() // empty services
	env.ParseFile(e1, ".private1.env") // set env with custom file name
	env.ParseFile(e2, ".private2.env") // set env with custom file name
	handler1 := newXHandler(e1) // clients use env
	handler2 := newYHandler(e2) // clients use env
	server := server.New(handler1, handler2) // separate environments
	...
	// multiple env #2
	env := env.Global()
	dbenv := env.Match("DB_")
	pay := env.Match("PAY_")
	handler1 := newXHandler(dbenv) // clients use env
	handler2 := newYHandler(pay) // clients use env
	server := server.New(handler1, handler2) // separate environments
}
```

## `cmd/dotenv`

Executable: print all values in global environment

```sh
$ dotenv
open .env failed
$ touch .env
$ dotenv
env is empty
$ echo ENV=pro > .env
$ dotenv
ENV=pro
$ dotenv -a ENV=dev -flag=x flag=y
ENV=dev
a=true
flag=y
```
