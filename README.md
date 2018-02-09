# Golang dotenv

Unify env data, by being as simple as possible

```
go get ztaylor.me/env
```

# CLI Usage

Use go get to install the command

```
go get ztaylor.me/env/cmd/dotenv
```

This will install `dotenv` to your `$GOPATH/bin`. This program will print all env values. You may now test any directory for its' env values.

```
zach@ztaylor.me /tmp $ ls
zach@ztaylor.me /tmp $ dotenv
zach@ztaylor.me /tmp $ echo
zach@ztaylor.me /tmp $ echo "ENV=dev" > .env
zach@ztaylor.me /tmp $ dotenv
ENV=dev
zach@ztaylor.me /tmp $ dotenv -a -b=b -flag=a flag=true
ENV=dev
b=b
flag=true
zach@ztaylor.me /tmp $
```

# Department of the Runtime Environment

Import the package

```
import (
	...
	"ztaylor.me/env"
	...
)
```

Bootstrap your CLI options and env files somewhere

```
func main() {
	...
	env.Bootstrap()
	env.Source("overwritepart.env")
	...

	if env.Name() == "dev" {
		// start "fast expire sessions" or other development options
	} else if env.Name() == "pro" {
		// start tls or minification or other production options
	}
}
```

Any value that is detected by `dotenv` will be available if you call `env.Bootstrap()` at least once.

# Why

I want to use a `.env` file to configure dependencies which should build together in a single program.

For example, an email server could be pluggable into another server that handles other services. In this example, the email configuration becomes a part of the parent program configuration.
