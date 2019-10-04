# Aurora setup

To build aurora you must have go and dep in your centOS enivironment. To install go you need to perform  below steps and go version must be >=1.9


# Install go

- sudo curl -O https://storage.googleapis.com/golang/go1.9.1.linux-amd64.tar.gz
- sudo tar -xvf go1.9.1.linux-amd64.tar.gz
- Sudo vi .bash_profile to set GOPATH, GOBIN,GOROOT 
- source ~/.bash_profile
- Check go version
  $ go version

# Install dep
- sudo -s
- echo  $PATH copy it and update path 
- Set path via export PATH=”{previous path}:/home/centos/go/bin”
- Set GOPATH export GOPATH=”/home/centos/go”
- Set GOROOT export GOROOT=”/home/centos/go”
- curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
- Check dep version 
```ssh
$ dep version

```

# Install aurora

- mkdir github.com/hcnet in /home/centos/go/src
- go to /home/centos/go/src/github.com/hcnet and execute following command:
- git clone https://github.com/HashCash-Consultants/go.git
- go to /home/centos/go/src/github.com/hcnet /go and execute following command:
```
	$ sudo yum install mercurial
	$ dep ensure –v
```
- go to /home/centos/go/src and execute following command:
```ssh
$ go install -ldflags "-X github.com/hcnet/go/support/app.version=aurora-0.16.0" github.com/hcnet/go/services/aurora/
```
- after running above command you check aurora build in <Your_dir>/go/bin folder and you can check aurora version by  following command :
```ssh
$ ./aurora version
```

# Aurora database setup
- Create a user for HC Net aurora database.
```
$ sudo -s
$ su – postgres
$ createuser <username> --pwprompt
$ Enter password for new role: <Enter password>
$ Enter it again: <Enter the pwd again>
```
- You need to add Aurora user. Exit from postgres and login as root user and execute following command.
```
$ exit
$ adduser <username>;
```
- To verify if user is created, execute following commands
```
$ su - postgres
$ psql
$ \du
```
- Create a blank database using following command.
```
 $ CREATE DATABASE <DB_NAME> OWNER <user created>;
```
 # Initialize aurora
 - Initialize aurora with database login as root user and Go   “go/bin” using following command
```
 $ export DATABASE_URL="postgresql://define aurora db username:define aurora db user password@localhost/define aurora database name?sslmode=disable"
```
- After that Go “go/bin” and execute following command
```
$ ./aurora db init
```
# Aurora up command
- Go “go/bin” execute following command
```
$ sudo nohup ./aurora --db-url="postgresql://define aurora db username:define aurora db user password @localhost/define aurora database name?sslmode=disable" --hcnet-core-db-url="postgresql://define Node1 database username:define Node1 database user password@localhost/define Node1 database name?sslmode=disable" --hcnet-core-url="http://localhost:11626" --network-passphrase="define Network password" --ingest="true" --per-hour-rate-limit=540000000 &
```
