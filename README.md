<p align="center"> <img src="docs/assets/projectavatar.png" width="360"></p> 
<p align="center"> <a href="https://travis-ci.org/oleg-balunenko/simple-chat"> 
        <img src="https://travis-ci.org/oleg-balunenko/simple-chat.svg?branch=master" alt="Build Status"></img>
    </a>
    <a href="https://goreportcard.com/report/github.com/oleg-balunenko/simple-chat">
        <img src="https://goreportcard.com/badge/github.com/oleg-balunenko/simple-chat" alt="Go Report Card"></img>
    </a>
    <a href='https://coveralls.io/github/oleg-balunenko/simple-chat?branch=master'>
        <img src='https://coveralls.io/repos/github/oleg-balunenko/simple-chat/badge.svg?branch=master' alt='Coverage Status' />
     </a>
    <a href="https://codecov.io/gh/oleg-balunenko/simple-chat">
      <img src="https://codecov.io/gh/oleg-balunenko/simple-chat/branch/master/graph/badge.svg" />
    </a>
    <a href="https://codebeat.co/projects/github-com-oleg-balunenko-simple-chat-master">
        <img alt="codebeat badge" src="https://codebeat.co/badges/2413b790-8465-42a2-aace-3e7a51750556" />
    </a>
    <a href="https://sonarcloud.io/dashboard?id=simple-chat">
        <img src="https://sonarcloud.io/api/project_badges/measure?project=simple-chat&metric=alert_status" alt="Quality Gate Status"></img>
    </a>
    <a href="https://app.codacy.com/app/oleg.balunenko/simple-chat?utm_source=github.com&utm_medium=referral&utm_content=oleg-balunenko/simple-chat&utm_campaign=Badge_Grade_Dashboard">
        <img src="https://api.codacy.com/project/badge/Grade/af78d928544e4f2b97e992dbed309b07" alt="Codacity code quality" />
    </a>
    <a href="https://github.com/oleg-balunenko/simple-chat/releases/latest">
        <img src="https://img.shields.io/badge/artifacts-download-blue.svg" alt ="Latest release artifacts"></img>
    </a>
</p>

# simple-chat



Chat application that allows to send messages between host and guest users

## Usage

### Flags

```text
 -ip string
    	server machine ip
 -listen
    	Listens on the specified ip address
 -port string
    	server port (default "8080")

```

#### Run host

Open console and run executable `simple-chat` file with flag `-listen` and pass the `ip` of your machine as argument

##### Example:

```bash
./simple-chat -listen -ip=192.168.02.11
```

#### Run guest

Open console and run executable `simple-chat` file and pass as the `ip` of `host` as argument

##### Example:

```bash
./simple-chat -ip=192.168.02.11

```

======================================

Now you can send messages via guest and host.
