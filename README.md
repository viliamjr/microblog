A very simple prototype (proof of concept) of Microservice architecture using [Go](http://golang.org/doc/) and [Nginx](http://nginx.org/en/docs/).

###What we have
There is two apps, *login* and *blog*. Both control the user by session and the second one let the user post a blog entry.

###What we want
As expected behaviour, the applications will be available at http://your-host-name.com/ as if it were one single app.

*Note: both apps and the package core are organized into a single repository to make easier to visualize the prototype.*

###Setup

1. Install [Go environment](http://golang.org/doc/install) and [Nginx server](http://nginx.org/en/docs/install.html).
2. And get the source/dependencies using Go:
```
$ go get \
github.com/viliamjr/microblog \
github.com/go-martini/martini \
github.com/martini-contrib/sessions \
github.com/martini-contrib/render
```

###Run

Start the apps:
- a login app instance:
```
$ PORT=3000 go run login.go
```

- the first instance of blog app:
```
$ PORT=3001 go run blog.go
```

- optionally, you can start a second instance of blog app:
```
$ PORT=3002 go run blog.go
```

- start nginx with the follow configuration:
```
upstream loginservice {
    ip_hash;
    server 127.0.0.1:3000;
}

upstream blogservice {
    ## ip_hash;
    server 127.0.0.1:3001;
    server 127.0.0.1:3002 weight=3;
}

server {
    listen       80;
    server_name  your-host-name.com;

    location /blog/ {
        proxy_pass http://blogservice/;
    }

    location / {
        proxy_pass http://loginservice;
    }
}
```
*Remember to define 'your-host-name.com' as a valid hostname at your hosts file.*

> Have fun!

###References:
* Microservices: http://martinfowler.com/articles/microservices.html
* Nginx: http://nginx.org/en/docs/
* Go: http://golang.org/doc/
* Martini: https://github.com/go-martini/martini

###TODO
* Change shared variables (keys and directories) to environment variables.
* Provide some documentation: readme and golang style.
