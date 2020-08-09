# Chapter10: Deploying Go

Web application by Go composed of a single compiled binary file so that is easy to deploy.
However, web application often does not built on only a single executable file but requires template, JavaScript, images, ond others. In this chapter, we are going to look through a way to deploy go web application with cloud providers;

- Completely on-premise server or IaaS provider server
- Cloud PaaS providers, e.g. Heroku or Google App Engine
- Docker container, deploy on local docker server and virtual machine in DigitalOcean.

This chapter basically introduces a way to deploy by a single user, but production environments require much more complicated ptocesses (test for each components, CI, server staging, etc).

## 10.1: Deploy to a server

`$ ./ws-s` starts an app on foreground but other task can't be done. `nohup ./ws-s &` force the OS to ignore HUP (hung-up) signal to the app. We can check the job by `ps aux | grep ws-s` and kill it by `kill (PID)`

We can do this by other init-equivalent daemons like Upstart or systemd. `init` is the first process launched at boot time on Unix-based systems, continues to run until the system shuts down. This us automatically launched by the kernel.

This chapter uses Upstart. Upstart is an event-driven program developed for Ubuntu Linux. Though systemd are popular, Upstart is easy to use. With Upstart, we create job-description file (ws.conf) and put it on `etc/init` directory.

``` 
respawn
respawn limit 10 5

setuid sausheong
setgid sausheong

exec /go/src/github.com/sausheong/ws-s/ws-s
```

To start teh Upstart job, runs

```
$ sudo start ws
ws start/running, process (PID)
```

Upstart automatically restart the job by this. For example;

```
$ ps -ef | grep ws
sausheo+ (PID) x x hh:mm ? 00:00:00 /go/src/github.com/sausheong/ws-s/ws-s

$ sudo kill -0 (PID)

$ ps -ef | grep ws
sausheo+ (newPID) x x hh:mm ? 00:00:00 /go/src/github.com/sausheong/ws-s/ws-s
```

## 10.2: Deploy to Heroku
Heroku enables applications built on manu programming languages. The *application* in Heroku is a set of source code and depending files.

Heroku assumes just two preparations;

- Depenedency file, like Gemfile in Ruby, package.json in Node.js, pom.xml in Java.
- Procfile defining the execution target. We can execute multiple files.

A procedure to deploy to Heroku;

1. Modify codes, get port from environment variables
2. Process dependency files by Godep
3. Create Heroku application
4. Push codes to Heroku

Problems faced in my case and solutions;

- `heroku create` commands did not automatically create remote `heroku` repository in my git environment
  - Manually add remote repository: `git remote add heroku https://git.heroku.com/<heroku-app-name>.git`
- Error at Heroku deployment: `No default language could be detected for this app.`
  - This is caused by hierarchical directory structure. Normal Heroku deployment (`git push heroku master`) assumes an application to be placed at a root directory
  - push subdirectory: `git subtree push --prefix ch10/ws-h heroku master` at the project root directory
- Error at Heroku deployment: `data.go:5:2: cannot find package "github.com/lib/pq"`
  - dependency files (`ch10/ws-h/vendor`) must be pushed (I ignored this in .gitignore)
  - push `ch10/ws-h/vender` (`ch10/ws-h/Godeps` also)
- Connection with database (postgres)
  - Activate Heroku Postgres in *resouce* tab in Heroku dashboard
  - Write the database setting in `data.go`
  - Create table from local `$ heroku pg:psql <database-name> --app <heroku-app-name>`
	- `--> Connecting to <database-name>`
	- `DATABASE=> CREATE TABLE posts (id serial PRIMARY KEY, content TEXT, author VARCHAR(255))`
  - Post some content; `$ curl -i -X POST -H "Content-Type: application/json" -d '{"content":"My first post", "author":"Sau Sheong"}' https://<heroku-app-name>.herokuapp.com:/post/`
  - Get (retrieve) content: `$ curl -i -X GET https://<heroku-app-name>.herokuapp.com/post/1`
