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

# Department of the Runtime Environment
`import "ztaylor.me/env"`
### package env automatically bootstraps during init

```
func main() {
	...
	env.Source("overwritepart.env")
	...
	if env.Name() == "pro" {
		// start tls 
		// start minification
	}
}
```

# Why

A single `.env` file can manage lots of runtime settings, and runtime flags are interchangeable.