```console

$ cat main.go
package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Database: "123",
		Username: "456",
		Password: "789",
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(session)
}

$ godep version
godep v56 (darwin/amd64/go1.6)

$ godep save

$ git status
On branch master

Initial commit

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	Godeps/
	main.go
	vendor/

nothing added to commit but untracked files present (use "git add" to track)

$ git add -A .
$ git status
On branch master

Initial commit

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)

	new file:   Godeps/Godeps.json
	new file:   Godeps/Readme
	new file:   main.go
	new file:   vendor/gopkg.in/mgo.v2/LICENSE
	new file:   vendor/gopkg.in/mgo.v2/Makefile
	new file:   vendor/gopkg.in/mgo.v2/README.md
	new file:   vendor/gopkg.in/mgo.v2/auth.go
	new file:   vendor/gopkg.in/mgo.v2/bson/LICENSE
	new file:   vendor/gopkg.in/mgo.v2/bson/bson.go
	new file:   vendor/gopkg.in/mgo.v2/bson/decode.go
	new file:   vendor/gopkg.in/mgo.v2/bson/encode.go
	new file:   vendor/gopkg.in/mgo.v2/bulk.go
	new file:   vendor/gopkg.in/mgo.v2/cluster.go
	new file:   vendor/gopkg.in/mgo.v2/doc.go
	new file:   vendor/gopkg.in/mgo.v2/gridfs.go
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sasl.c
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sasl.go
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sasl_windows.c
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sasl_windows.go
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sasl_windows.h
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sspi_windows.c
	new file:   vendor/gopkg.in/mgo.v2/internal/sasl/sspi_windows.h
	new file:   vendor/gopkg.in/mgo.v2/internal/scram/scram.go
	new file:   vendor/gopkg.in/mgo.v2/log.go
	new file:   vendor/gopkg.in/mgo.v2/queue.go
	new file:   vendor/gopkg.in/mgo.v2/raceoff.go
	new file:   vendor/gopkg.in/mgo.v2/raceon.go
	new file:   vendor/gopkg.in/mgo.v2/saslimpl.go
	new file:   vendor/gopkg.in/mgo.v2/saslstub.go
	new file:   vendor/gopkg.in/mgo.v2/server.go
	new file:   vendor/gopkg.in/mgo.v2/session.go
	new file:   vendor/gopkg.in/mgo.v2/socket.go
	new file:   vendor/gopkg.in/mgo.v2/stats.go

$ git commit -am "Test"
[master (root-commit) 8478bc9] Test
 33 files changed, 11671 insertions(+)
 create mode 100644 Godeps/Godeps.json
 create mode 100644 Godeps/Readme
 create mode 100644 main.go
 create mode 100644 vendor/gopkg.in/mgo.v2/LICENSE
 create mode 100644 vendor/gopkg.in/mgo.v2/Makefile
 create mode 100644 vendor/gopkg.in/mgo.v2/README.md
 create mode 100644 vendor/gopkg.in/mgo.v2/auth.go
 create mode 100644 vendor/gopkg.in/mgo.v2/bson/LICENSE
 create mode 100644 vendor/gopkg.in/mgo.v2/bson/bson.go
 create mode 100644 vendor/gopkg.in/mgo.v2/bson/decode.go
 create mode 100644 vendor/gopkg.in/mgo.v2/bson/encode.go
 create mode 100644 vendor/gopkg.in/mgo.v2/bulk.go
 create mode 100644 vendor/gopkg.in/mgo.v2/cluster.go
 create mode 100644 vendor/gopkg.in/mgo.v2/doc.go
 create mode 100644 vendor/gopkg.in/mgo.v2/gridfs.go
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sasl.c
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sasl.go
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sasl_windows.c
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sasl_windows.go
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sasl_windows.h
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sspi_windows.c
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/sasl/sspi_windows.h
 create mode 100644 vendor/gopkg.in/mgo.v2/internal/scram/scram.go
 create mode 100644 vendor/gopkg.in/mgo.v2/log.go
 create mode 100644 vendor/gopkg.in/mgo.v2/queue.go
 create mode 100644 vendor/gopkg.in/mgo.v2/raceoff.go
 create mode 100644 vendor/gopkg.in/mgo.v2/raceon.go
 create mode 100644 vendor/gopkg.in/mgo.v2/saslimpl.go
 create mode 100644 vendor/gopkg.in/mgo.v2/saslstub.go
 create mode 100644 vendor/gopkg.in/mgo.v2/server.go
 create mode 100644 vendor/gopkg.in/mgo.v2/session.go
 create mode 100644 vendor/gopkg.in/mgo.v2/socket.go
 create mode 100644 vendor/gopkg.in/mgo.v2/stats.go

$ heroku create
Creating app... done, stack is cedar-14
https://fast-garden-78997.herokuapp.com/ | https://git.heroku.com/fast-garden-78997.git

$ git push heroku master
Counting objects: 43, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (40/40), done.
Writing objects: 100% (43/43), 101.19 KiB | 0 bytes/s, done.
Total 43 (delta 2), reused 0 (delta 0)
remote: Compressing source files... done.
remote: Building source:
remote:
remote: -----> Go app detected
remote: -----> Checking Godeps/Godeps.json file.
remote: -----> Installing go1.6... done
remote: -----> Running: go install -v -tags heroku .
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/bson
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/scram
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2
remote: github.com/freeformz/tmgo
remote: -----> Discovering process types
remote:        Procfile declares types -> (none)
remote:
remote: -----> Compressing...
remote:        Done: 1.7M
remote: -----> Launching...
remote:        Released v3
remote:        https://fast-garden-78997.herokuapp.com/ deployed to Heroku
remote:
remote: Verifying deploy.... done.
To https://git.heroku.com/fast-garden-78997.git
 * [new branch]      master -> master
 ```