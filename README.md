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
 
 ## And now the error
 
 ```console
$ godep save ./...

$ git status
On branch master
Your branch is up-to-date with 'origin/master'.
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   Godeps/Godeps.json

no changes added to commit (use "git add" and/or "git commit -a")

$ git diff
diff --git a/Godeps/Godeps.json b/Godeps/Godeps.json
index 432aa37..b90e0ce 100644
--- a/Godeps/Godeps.json
+++ b/Godeps/Godeps.json
@@ -1,6 +1,9 @@
 {
        "ImportPath": "github.com/freeformz/tmgo",
        "GoVersion": "go1.6",
+       "Packages": [
+               "./..."
+       ],
        "Deps": [
                {
                        "ImportPath": "gopkg.in/mgo.v2",

$ git add Godeps/Godeps.json

$ git commit -am "Maybe"
[master f69012b] Maybe
 1 file changed, 3 insertions(+)

$ git push heroku master
Counting objects: 7, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (7/7), done.
Writing objects: 100% (7/7), 2.06 KiB | 0 bytes/s, done.
Total 7 (delta 1), reused 0 (delta 0)
remote: Compressing source files... done.
remote: Building source:
remote:
remote: -----> Using set buildpack heroku/go
remote: -----> Go app detected
remote: -----> Checking Godeps/Godeps.json file.
remote: -----> Using go1.6
remote: -----> Running: go install -v -tags heroku ./...
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/bson
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/scram
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/sasl
remote: # github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/sasl
remote: vendor/gopkg.in/mgo.v2/internal/sasl/sasl.go:15:24: fatal error: sasl/sasl.h: No such file or directory
remote:  // #include <sasl/sasl.h>
remote:                         ^
remote: compilation terminated.
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2
remote: github.com/freeformz/tmgo
remote:
remote:  !     Push rejected, failed to compile Go app
remote:
remote: Verifying deploy...
remote:
remote: !	Push rejected to fast-garden-78997.
remote:
To https://git.heroku.com/fast-garden-78997.git
 ! [remote rejected] master -> master (pre-receive hook declined)
error: failed to push some refs to 'https://git.heroku.com/fast-garden-78997.git'
 ```
 
 ## Still Broken
 
 ```console
$ heroku buildpacks:set https://github.com/ddollar/heroku-buildpack-apt.git
Buildpack set. Next release on fast-garden-78997 will use https://github.com/ddollar/heroku-buildpack-apt.git.
Run git push heroku master to create a new release using this buildpack.

$ heroku buildpacks:add heroku/go
Buildpack added. Next release on fast-garden-78997 will use:
  1. https://github.com/ddollar/heroku-buildpack-apt.git
  2. heroku/go
Run git push heroku master to create a new release using these buildpacks.

$ echo "libsasl2-dev" > Aptfile

$ git add Aptfile; git commit -am "Aptfile"
[master 8e9a337] Aptfile
 1 file changed, 1 insertion(+)
 create mode 100644 Aptfile

$ git push heroku master
Counting objects: 13, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (12/12), done.
Writing objects: 100% (13/13), 3.13 KiB | 0 bytes/s, done.
Total 13 (delta 3), reused 0 (delta 0)
remote: Compressing source files... done.
remote: Building source:
remote:
remote: -----> Fetching set buildpack https://github.com/ddollar/heroku-buildpack-apt.git... done
remote: -----> Apt app detected
remote: -----> Updating apt caches
remote:        Get:1 http://apt.postgresql.org trusty-pgdg InRelease [30.1 kB]
remote:        Ign http://archive.ubuntu.com trusty InRelease
remote:        Get:2 http://archive.ubuntu.com trusty-security InRelease [65.9 kB]
remote:        Get:3 http://archive.ubuntu.com trusty-updates InRelease [65.9 kB]
remote:        Get:4 http://archive.ubuntu.com trusty Release.gpg [933 B]
remote:        Get:5 http://archive.ubuntu.com trusty Release [58.5 kB]
remote:        Get:6 http://apt.postgresql.org trusty-pgdg/main amd64 Packages [57.5 kB]
remote:        Get:7 http://archive.ubuntu.com trusty-security/main amd64 Packages [430 kB]
remote:        Ign http://apt.postgresql.org trusty-pgdg/main Translation-en_US
remote:        Ign http://apt.postgresql.org trusty-pgdg/main Translation-en
remote:        Get:8 http://archive.ubuntu.com trusty-security/main Translation-en [235 kB]
remote:        Get:9 http://archive.ubuntu.com trusty-updates/main amd64 Packages [712 kB]
remote:        Get:10 http://archive.ubuntu.com trusty-updates/main Translation-en [360 kB]
remote:        Get:11 http://archive.ubuntu.com trusty/main amd64 Packages [1,350 kB]
remote:        Get:12 http://archive.ubuntu.com trusty/universe amd64 Packages [5,859 kB]
remote:        Get:13 http://archive.ubuntu.com trusty/main Translation-en [762 kB]
remote:        Get:14 http://archive.ubuntu.com trusty/universe Translation-en [4,089 kB]
remote:        Ign http://archive.ubuntu.com trusty/main Translation-en_US
remote:        Ign http://archive.ubuntu.com trusty/universe Translation-en_US
remote:        Fetched 14.1 MB in 8s (1,699 kB/s)
remote:        Reading package lists...
remote: -----> Fetching .debs for libsasl2-dev
remote:        Reading package lists...
remote:        Building dependency tree...
remote:        The following NEW packages will be installed:
remote:          libsasl2-dev
remote:        0 upgraded, 1 newly installed, 0 to remove and 8 not upgraded.
remote:        Need to get 311 kB of archives.
remote:        After this operation, 827 kB of additional disk space will be used.
remote:        Get:1 http://archive.ubuntu.com/ubuntu/ trusty/main libsasl2-dev amd64 2.1.25.dfsg1-17build1 [311 kB]
remote:        Fetched 311 kB in 0s (674 kB/s)
remote:        Download complete and in download only mode
remote: -----> Installing libsasl2-dev_2.1.25.dfsg1-17build1_amd64.deb
remote: -----> Writing profile script
remote: -----> Using set buildpack heroku/go
remote: -----> Go app detected
remote: -----> Checking Godeps/Godeps.json file.
remote: -----> Using go1.6
remote: -----> Running: go install -v -tags heroku ./...
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/bson
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/scram
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/sasl
remote: github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2
remote: # github.com/freeformz/tmgo/vendor/gopkg.in/mgo.v2/internal/sasl
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(dlopen.o): In function `_sasl_locate_entry':
remote: (.text+0x1b): undefined reference to `dlsym'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(dlopen.o): In function `_sasl_get_plugin':
remote: (.text+0xd0): undefined reference to `dlopen'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(dlopen.o): In function `_sasl_get_plugin':
remote: (.text+0x107): undefined reference to `dlerror'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(dlopen.o): In function `_sasl_done_with_plugins':
remote: (.text+0x310): undefined reference to `dlclose'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(db_berkeley.o): In function `berkeleydb_open':
remote: (.text+0x56): undefined reference to `db_create'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(db_berkeley.o): In function `berkeleydb_open':
remote: (.text+0x13a): undefined reference to `db_strerror'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(db_berkeley.o): In function `berkeleydb_close.isra.0':
remote: (.text+0x1b6): undefined reference to `db_strerror'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(db_berkeley.o): In function `_sasldb_getdata':
remote: (.text+0x440): undefined reference to `db_strerror'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(db_berkeley.o): In function `_sasldb_putdata':
remote: (.text+0x682): undefined reference to `db_strerror'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(db_berkeley.o): In function `_sasldb_putdata':
remote: (.text+0x70a): undefined reference to `db_strerror'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `init_3des':
remote: (.text+0x682): undefined reference to `DES_key_sched'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `init_3des':
remote: (.text+0x6a6): undefined reference to `DES_key_sched'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `init_3des':
remote: (.text+0x6de): undefined reference to `DES_key_sched'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `init_3des':
remote: (.text+0x6fd): undefined reference to `DES_key_sched'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `init_des':
remote: (.text+0x7ae): undefined reference to `DES_key_sched'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o):(.text+0x7de): more undefined references to `DES_key_sched' follow
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `dec_3des':
remote: (.text+0x860): undefined reference to `DES_ede3_cbc_encrypt'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `enc_3des':
remote: (.text+0x95d): undefined reference to `DES_ede3_cbc_encrypt'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `dec_des':
remote: (.text+0x9b0): undefined reference to `DES_cbc_encrypt'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(digestmd5.o): In function `enc_des':
remote: (.text+0xab2): undefined reference to `DES_cbc_encrypt'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_hash':
remote: (.text+0xaa): undefined reference to `EVP_DigestInit'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_hash':
remote: (.text+0xbd): undefined reference to `EVP_DigestUpdate'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_hash':
remote: (.text+0xcd): undefined reference to `EVP_DigestFinal'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `generate_otp.isra.3':
remote: (.text+0x44b): undefined reference to `EVP_get_digestbyname'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `word2bin.isra.5':
remote: (.text+0xa22): undefined reference to `EVP_DigestInit'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `word2bin.isra.5':
remote: (.text+0xa42): undefined reference to `EVP_DigestUpdate'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `word2bin.isra.5':
remote: (.text+0xa54): undefined reference to `EVP_DigestFinal'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_server_mech_step':
remote: (.text+0x1ea1): undefined reference to `EVP_get_digestbyname'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_server_mech_step':
remote: (.text+0x28e3): undefined reference to `EVP_get_digestbyname'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_server_plug_init':
remote: (.text+0x2c31): undefined reference to `OpenSSL_add_all_digests'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_client_plug_init':
remote: (.text+0x2c81): undefined reference to `OpenSSL_add_all_digests'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(otp.o): In function `otp_common_mech_free':
remote: (.text+0x1a1): undefined reference to `EVP_cleanup'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_free_context_contents':
remote: (.text+0x6b): undefined reference to `gss_delete_sec_context'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_free_context_contents':
remote: (.text+0x88): undefined reference to `gss_release_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_free_context_contents':
remote: (.text+0xa5): undefined reference to `gss_release_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_free_context_contents':
remote: (.text+0xc2): undefined reference to `gss_release_cred'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_free_context_contents':
remote: (.text+0xdf): undefined reference to `gss_release_cred'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_seterror_':
remote: (.text+0x35e): undefined reference to `gss_display_status'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_seterror_':
remote: (.text+0x3cb): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_seterror_':
remote: (.text+0x4e6): undefined reference to `gss_display_status'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_seterror_':
remote: (.text+0x557): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x7d6): undefined reference to `gss_unwrap'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x93b): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xa6f): undefined reference to `gss_accept_sec_context'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xb20): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xbaa): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xc39): undefined reference to `gss_display_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xcfb): undefined reference to `GSS_C_NT_USER_NAME'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xd18): undefined reference to `gss_import_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xd80): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xde9): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0xe81): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x106e): undefined reference to `GSS_C_NT_HOSTBASED_SERVICE'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x1076): undefined reference to `gss_import_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x10fb): undefined reference to `gss_release_cred'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x1179): undefined reference to `gss_acquire_cred'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x1238): undefined reference to `gss_wrap_size_limit'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x12fd): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x1404): undefined reference to `gss_wrap'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x148e): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x14db): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x1635): undefined reference to `gss_compare_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x169d): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x16e4): undefined reference to `gss_release_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x17f2): undefined reference to `gss_release_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_server_mech_step':
remote: (.text+0x189f): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_encode':
remote: (.text+0x1a08): undefined reference to `gss_wrap'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_encode':
remote: (.text+0x1b31): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_encode':
remote: (.text+0x1b93): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `sasl_gss_encode':
remote: (.text+0x1bb7): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_decode_packet':
remote: (.text+0x1c84): undefined reference to `gss_unwrap'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_decode_packet':
remote: (.text+0x1d0a): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_decode_packet':
remote: (.text+0x1d70): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_decode_packet':
remote: (.text+0x1dd3): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x1fda): undefined reference to `gss_unwrap'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x205e): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x21a7): undefined reference to `gss_init_sec_context'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2259): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x22ab): undefined reference to `gss_delete_sec_context'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x23cb): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x248b): undefined reference to `gss_wrap'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x252f): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x25e9): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x264c): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x26b5): undefined reference to `gss_inquire_context'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2717): undefined reference to `gss_display_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2769): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2947): undefined reference to `GSS_C_NT_HOSTBASED_SERVICE'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2961): undefined reference to `gss_import_name'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2be3): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2c28): undefined reference to `gss_wrap_size_limit'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(gssapi.o): In function `gssapi_client_mech_step':
remote: (.text+0x2ca3): undefined reference to `gss_release_buffer'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `P16_nt':
remote: (.text+0x83a): undefined reference to `MD4'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0x9c9): undefined reference to `EVP_md5'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0x9ee): undefined reference to `HMAC'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0x9f3): undefined reference to `EVP_md5'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0xa05): undefined reference to `HMAC_Init'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0xa17): undefined reference to `HMAC_Update'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0xa2b): undefined reference to `HMAC_Update'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0xa39): undefined reference to `HMAC_Final'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `V2':
remote: (.text+0xa41): undefined reference to `HMAC_CTX_cleanup'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `E.constprop.8':
remote: (.text+0x11af): undefined reference to `DES_set_odd_parity'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `E.constprop.8':
remote: (.text+0x11ba): undefined reference to `DES_set_key'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ntlm.o): In function `E.constprop.8':
remote: (.text+0x11d3): undefined reference to `DES_ecb_encrypt'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_exec':
remote: (.text+0x1049): undefined reference to `sqlite3_exec'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_exec':
remote: (.text+0x107c): undefined reference to `sqlite3_free'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_open':
remote: (.text+0x11a5): undefined reference to `sqlite3_open'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_open':
remote: (.text+0x11c2): undefined reference to `sqlite3_exec'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_open':
remote: (.text+0x11f6): undefined reference to `sqlite3_errmsg'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_open':
remote: (.text+0x1214): undefined reference to `sqlite3_close'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_open':
remote: (.text+0x1253): undefined reference to `sqlite3_free'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x12de): undefined reference to `PQexec'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x12e9): undefined reference to `PQresultStatus'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x12ff): undefined reference to `PQntuples'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x1337): undefined reference to `PQgetvalue'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x1364): undefined reference to `PQclear'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x138a): undefined reference to `PQresStatus'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x13a7): undefined reference to `PQclear'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_exec':
remote: (.text+0x13d1): undefined reference to `PQclear'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_escape_str':
remote: (.text+0x145e): undefined reference to `PQescapeString'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_open':
remote: (.text+0x1586): undefined reference to `PQconnectdb'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_open':
remote: (.text+0x1599): undefined reference to `PQstatus'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_open':
remote: (.text+0x17dd): undefined reference to `PQerrorMessage'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x187d): undefined reference to `mysql_real_query'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x1885): undefined reference to `mysql_errno'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x1895): undefined reference to `mysql_field_count'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x18b4): undefined reference to `mysql_store_result'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x18c8): undefined reference to `mysql_num_rows'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x18f7): undefined reference to `mysql_fetch_row'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x193b): undefined reference to `mysql_free_result'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x196b): undefined reference to `mysql_error'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x1992): undefined reference to `mysql_free_result'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_exec':
remote: (.text+0x19d2): undefined reference to `mysql_free_result'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_escape_str':
remote: (.text+0x1a5e): undefined reference to `mysql_escape_string'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_open':
remote: (.text+0x1ae4): undefined reference to `mysql_init'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_open':
remote: (.text+0x1b38): undefined reference to `mysql_real_connect'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_sqlite3_close':
remote: (.text+0x1001): undefined reference to `sqlite3_close'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_pgsql_close':
remote: (.text+0x12b1): undefined reference to `PQfinish'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(sql.o): In function `_mysql_close':
remote: (.text+0x1831): undefined reference to `mysql_close'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_interact':
remote: (.text+0xdc): undefined reference to `ldap_get_option'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_connect.isra.0':
remote: (.text+0x1e5): undefined reference to `ldap_initialize'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_connect.isra.0':
remote: (.text+0x263): undefined reference to `ldap_set_option'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_connect.isra.0':
remote: (.text+0x277): undefined reference to `ldap_start_tls_s'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_connect.isra.0':
remote: (.text+0x2b5): undefined reference to `ldap_sasl_interactive_bind_s'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_connect.isra.0':
remote: (.text+0x2df): undefined reference to `ldap_whoami_s'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_connect.isra.0':
remote: (.text+0x311): undefined reference to `ber_bvfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_store':
remote: (.text+0x464): undefined reference to `ldap_err2string'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_store':
remote: (.text+0x498): undefined reference to `ldap_unbind_ext'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_store':
remote: (.text+0x523): undefined reference to `ldap_modify_ext_s'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_store':
remote: (.text+0x530): undefined reference to `ber_bvfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x748): undefined reference to `ldap_unbind_ext'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x825): undefined reference to `ldap_search_ext_s'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x836): undefined reference to `ber_bvfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x856): undefined reference to `ldap_first_message'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x883): undefined reference to `ldap_next_message'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x89b): undefined reference to `ldap_msgtype'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x8cc): undefined reference to `ldap_get_values_len'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x946): undefined reference to `ber_bvecfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_auxprop_lookup':
remote: (.text+0x98f): undefined reference to `ldap_msgfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xa71): undefined reference to `ldap_unbind_ext'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xb40): undefined reference to `ldap_search_ext_s'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xb4f): undefined reference to `ber_bvfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xb66): undefined reference to `ldap_first_message'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xb89): undefined reference to `ldap_next_message'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xb9d): undefined reference to `ldap_msgtype'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xbb4): undefined reference to `ldap_get_values_len'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xc08): undefined reference to `ber_bvecfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xc27): undefined reference to `ldap_err2string'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xc53): undefined reference to `ldap_msgfree'
remote: /tmp/build_831e7eeb26e7caa03cd1e46a26542418/.apt/usr/lib/x86_64-linux-gnu/libsasl2.a(ldapdb.o): In function `ldapdb_canon_server':
remote: (.text+0xcc1): undefined reference to `ber_bvfree'
remote: collect2: error: ld returned 1 exit status
remote: github.com/freeformz/tmgo
remote:
remote:  !     Push rejected, failed to compile Go app
remote:
remote: Verifying deploy...
remote:
remote: !	Push rejected to fast-garden-78997.
remote:
To https://git.heroku.com/fast-garden-78997.git
 ! [remote rejected] master -> master (pre-receive hook declined)
error: failed to push some refs to 'https://git.heroku.com/fast-garden-78997.git'
 ```