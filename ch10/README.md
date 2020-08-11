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


## 10.3: Deploy to Google App Engine

GAE has some advantages in performance and scalability over other PaaS like Heroku. 
GAE automatically scales out the application as needed, implements many tooles and utilities (Google account authorization, sending mail, log generation, etc).
GAE also has some disadvantages; file system is read-only, request is keep for 60 min, does not provide direct network access. This means that GAE can not normally access to the other services out of the application environment.

Procedures in GAE deployment;

- Code modification: add statement to call Google library
- Create `app.yaml`
- Create GAE application
- Push codes to GAE application

Problems and solutions;

- Table creation
  - Create Google CloudSQL instance (MySQL)
  - Launch Google Cloud Console and access into CloudSQL: `$ gcloud sql connect <instance-id> --user=root`, and put password
  - Create database and table in it
	- `mysql> CREATE DATABASE <database-name>`
	- `mysql> CREATE TABLE posts (id serial PRIMARY KEY, content TEXT, author VARCHAR(255))`
- Connection between GAE and CloudSQL
  - [Official documentation](https://cloud.google.com/sql/docs/mysql/connect-app-engine#go)
  - Description in this book is for an old version of Go with GAE and I fix `app.yaml` and `init()` in `data.go` for the correct connection.
  
`app.yaml`;

``` yaml
runtime: go111

env_variables:
  INSTANCE_CONNECTION_NAME: <project>:<region>:<instance-id>
  DB_USER: root
  DB_PASS: <password>
  DB_NAME: <database-name>
```

Database connection part of `data.go`;

``` go
func init() {
	// [START cloud_sql_mysql_databasesql_create_socket]
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	
	var err error
	Db, err = sql.Open("mysql", dbURI)
	if err != nil {
		panic(err)
	}
}
```

Then I have successfully deployed the app by `$ gcloud app deploy`. We can create/retrieve/update/delete post with REST-API. For example;

- `$ curl -i -X POST -H "Content-Type: application/json" -d '{"content":"Hello World!", "author": "Sau Sheong"}' https://xxx.appspot.com/post/` posts a new content
- `$ curl -i -X GET https://ws-g-dhirooka.uc.r.appspot.com/post/1` retrieves the specified content

## 10.4: Deploy to Docker

This chapter explains about Docker and how to deploy and run a Go web application as a Docker container.

Docker is an open platform enables applications to be built, move, execute on a container.
Container is a type of infrastructure virtualization. VM emurates a whole computer system including OS. 
Container virtualizes OS-level operations. This splits the computer resources by multiple user-space instances. As a result, container requires much less resource than VM, can be launched and deployed faster.

Docker concepts and components;

- **Docker engine** (or simply Docker) is composed of multiple components
- **Docke client** is a command-line interface allows users to interact with Docker daemon
- **Docker daemon** is a process working on the host OS, responds a service call, manages containers
- **Docker container** (or container) is a whole program (including OS) composing the target application
  - The reason why container is light is that applications and other bundled programms only behaves like occupying OS, but actually shares host OS
  - Docker container is built based on Docker image
  - Writing Dockerfile is a way to create Docker image
- **Docker image** can be contained in the local environment same as the Docker daemon, or host on Docker registry
  - We can use private Docker registry or use Docker Hub as ones registry

When installing Docker into Linux OS like Ubuntu, Docker daemon and Docker client is installed in the same machine.
If other OS (maxOS, Windows, etc) exists, daemon is normally installed in the VM on the OS.
Docker container are built from Docker image, run on a Docker host.

Docker-ize Go web application;

1. Create Dockerfile
2. Create Docker image from the Dockerfile
3. Build Docke container from Docker image
4. Create Docker host on a cloud provider
5. Connect to the remote Docker host
6. Built Docker image on the remote host
7. Launch Docker container on the remote host

I have skipped the hands-on Docker application because this is less exmplained in the book. Deep explanations about deployment would be better to read other books or courses.
Code `ws-d` is copied from the [official repository](https://github.com/mushahiroyuki/gowebprog).

## 10.5: Comparison of deployment solutions

|      |Standalone      |Heroku   |GAE   |Docker   |
|:-----|:-----|:--|:--|:--|
|Type      |Public/Private  |public   |public   |public   |
|Code modification |None      |Low      |Middle   |None   |
|System operation      |High      |None   |None   |Middle   |
|Maintenance      |High      |None   |None   |Middle   |
|Easy deployment |Low |High |Middle |Low      |
|Platform support      |None      |Low   |High   |Low   |
|Platform dependency      |None      |Low   |High   |Low   |
|Scalability      |None      |Middle   |High   |High   |

## 10.6: Summary

- We can easily deploy Go web applications by create and put a binary executable file on a VM or a phisical server, and implement Upstart to launch and running it.
- Heroku is a simple PaaS. We only have to modify a few part of scripts, prepare dependency files by godep, create Procfile. By pushing them to Heroku git repository, the application is deployed.
- GAE is a sophisticated and sandboxed PaaS provided by Google. Deploying to GAE is a little complicated, but the app get scalable. Most part of restriction with GAE sandbox is about to use of external services.
- Docker is a good solution to deploy web services or applications, but more complicated. We are required to 1.) create a container of the web service, 2.) deploy it to the local Docker host, 3.) deploy it to the remote Docker host.
