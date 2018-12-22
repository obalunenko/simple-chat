# simple-chat

Chat application that allows to send messages between host and guest users


## Usage:

Flags:
 
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
