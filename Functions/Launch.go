package Functions

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"time"
)

func Launch() {
	args := os.Args
	Name := ""
	i := 0
	NameList := []string{}
	log := []string{}
	port := ":"
	client := make(chan net.Conn)
	nbrOfC := 0
	disconnect := make(chan net.Conn)
	message := make(chan string)
	connexion := make(map[string]net.Conn)

	content, _ := ioutil.ReadFile("linux.txt")

	if len(args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	if len(args) < 2 {
		port += "8989"
	} else {
		port += args[1]
	}

	connect, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on the port", port)

	go client_Connexion(connect, client)

	for {
		select {
		case lastClient := <-client:

			nbrOfC++
			if nbrOfC > 10 {
				fmt.Fprintf(lastClient, "Only 10 connexions are authorized")
				os.Exit(0)
			}
			fmt.Fprint(lastClient, string(content)+"\n")
			fmt.Fprint(lastClient, "[ENTER YOUR NAME]: ")
			fmt.Fscanf(lastClient, "%s", &Name)
			for {
				if Name == "" {
					fmt.Fprint(lastClient, "\x1bc")
					i++
					fmt.Fprint(lastClient, i, "Empty name....", "..", "Cannot accept empty name"+"\n")
					fmt.Fprint(lastClient, string(content)+"\n")
					fmt.Fprint(lastClient, "[ENTER YOUR NAME]: ")
					fmt.Fscanf(lastClient, "%s", &Name)

				} else {
					if len(log) > 0 {
						for _, ligne := range log {
							fmt.Fprint(lastClient, ligne)
						}

					}
					connexion[Name] = lastClient

					break
				}
			}
			NameList = append(NameList, Name)
			fmt.Println(NameList)
			log = append(log, Name, " has joined our chat.."+"\n")
			for _, joined := range connexion {
				if joined == lastClient {
					continue
				}
				fmt.Fprint(joined, Name, " has joined our chat.."+"\n")
			}

			go HandleMess(lastClient, disconnect, Name, message)

		case MSG := <-message:
			log = append(log, MSG+"\n")
			for _, client := range connexion {
				fmt.Fprint(client, MSG+"\n")

				fmt.Println(connexion)
			}

		case disconnected := <-disconnect:

			nbrOfC--
			delete(connexion, disconnected.RemoteAddr().String())
			fmt.Println(connexion)
			for key, dis := range connexion {
				if dis == disconnected {
					Name = key
					continue
				}
				fmt.Fprint(dis, Name, " has left our chat..\n")
			}
			log = append(log, Name, " has left our chat..\n")
		}
	}
}

func client_Connexion(connect net.Listener, client chan net.Conn) {
	for {
		client_Connected, err := connect.Accept()
		if err != nil {
			fmt.Println("Cannot connecting client")
			return
		}
		client <- client_Connected
	}
}

func HandleMess(client net.Conn, disconnect chan net.Conn, Name string, message chan string) {
	date := time.Now().String()

	for {
		Mess, err := bufio.NewReader(client).ReadString('\n')
		Mess = strings.Trim(Mess, "\n")
		if err != nil {
			fmt.Fprint(client, "Cannot accept empty name")
		}
		if len(Mess) > 1 {
			RealDate := Convert(date)
			Send := RealDate + "[" + Name + "]:" + Mess
			message <- Send
		}
		if err == io.EOF {
			break
		}
	}
	disconnect <- client

}

func Convert(t string) string {
	Date := ""

	for i := 0; i < len(t); i++ {
		if i == 19 {
			Date += "[" + t[:i] + "]"
			break
		}
	}
	return Date

}
