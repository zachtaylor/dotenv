# Package env

`import "ztaylor.me/env"`

Package env provides runtime environment, using CLI flags, and/or config file, and/or environment variables

## Environment config

```go
func main() {
	// most common case
	env := env.Global()
	server := server.New(env) // clients use env
	...
	// custom env
	e := env.NewDefaultService() // uses os.Getenv
	env.ParseFile(e, "./CONFIG") // set env with custom file name
	env.ParseFlags(e) // set env with CLI flags
	server := server.New(e) // clients use env
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
