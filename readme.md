# package env
`import "ztaylor.me/env"`

Package env provides easy runtime environment configuration

### Why

A single `.env` file can manage lots of runtime settings, and flags are organized conveniently

## Runtime Environment

```
package...

func main() {
	env := env.Global() // most common case
	server := server.New(env) // create program logic using env
	db := db.New(env) // load db settings using env
	...
}
```

### `Global()` is batteries-included default `env.Provider`

`env.Provider` is a simple interface that is extensible. Functions that use `env env.Provider` arg are thus more adaptable than programs that use `ztaylor.me/env.Global()` directly.

## `dotenv` binary

Print all values in global environment, which uses the current directory's `.env` file by default (as well as cli flags)

```
zach@ztaylor.me / $ go get ztaylor.me/env/dotenv
```

### use `dotenv` command to test

```
zach@ztaylor.me /tmp $ dotenv
[env empty]
zach@ztaylor.me /tmp $ echo "ENV=dev" > .env
zach@ztaylor.me /tmp $ dotenv
ENV=dev
zach@ztaylor.me /tmp $ dotenv -a -b=b -flag=c flag=d
ENV=dev
a=true
b=b
flag=true
```
