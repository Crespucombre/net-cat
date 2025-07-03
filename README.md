# NetCat

Netcat is a chat create with Golang language.
You can connect by typing the command below.

## Command Line

Before beginning interactions with others, please make sure that the server is on.



```go 
go run .
```
You have to launch it before typing the second command : 

```bash
nc {YourIP} 8989
```
The port used in my project is 8989. 

## Usage

If the connection is successful you will have the interface follows;

```cmd
Welcome to the TCP-Tchat !
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: 
```

After you have typed your name, you have access to the chat.

```cmd
John has joined our chat..
John has left our chat..
John has joined our chat..
Doe has joined our chat..
[2025-07-03 14:43:48][Doe]:Hello John ! 
[2025-07-03 14:42:55][John]:Hello Doe
[2025-07-03 14:43:48][Doe]:How's going ?
[2025-07-03 14:42:55][John]:Great ! 
[2025-07-03 14:42:55][John]:I love this chat !
```

## Enjoy !

## Credits

[Netcat](https://fr.wikipedia.org/wiki/Netcat)

[TCP](https://fr.wikipedia.org/wiki/Transmission_Control_Protocol)

[GoRoutine](https://go.dev/tour/concurrency/1)
