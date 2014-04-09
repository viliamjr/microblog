
**How to setup the environment**

1) initialize the apps:

- a login app instance:
$ PORT=3000 go run login.go

- the first instance of blog app:
$ PORT=3001 go run blog.go

- optionally, you can start a second instance of blog app:
$ PORT=3002 go run blog.go

2) Start nginx with the follow configuration:

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

That is it. The apps are available at http://your-host-name.com/ as if it were a single app.
