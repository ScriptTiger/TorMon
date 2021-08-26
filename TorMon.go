package main

import "net"
import "fmt"
import "bufio"
import "os"
import "strings"
import "flag"

func main() {

	// Catch flags
	hostPtr := flag.String("host", "127.0.0.1", "host")
	portPtr := flag.String("port", "9051", "port")
	secretPtr := flag.String("secret", "", "secret")
	cmdPtr := flag.String("cmd", "", "cmd")

	// Parse flags
	flag.Parse()

	// Declare/initialize variables
	var text string
	console := bufio.NewReader(os.Stdin)
	conn, err := net.Dial("tcp", *hostPtr+":"+*portPtr)
	if err != nil {os.Exit(1)}

	// Authenticate connection to Tor
	fmt.Fprintf(conn, "authenticate \""+*secretPtr+"\"\n")
	text, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {os.Exit(2)}
	if  strings.TrimRight(text, "\r\n") != "250 OK" {os.Exit(3)}

	// Initialize function for receiving and handling command output
	receive := func() {
		// Handle responses
		text = ""
		tor := bufio.NewScanner(conn)
		for tor.Scan() {
			text = tor.Text()
			if text == "250 OK" {
				break
			} else if text == "." {
				break
			} else if text == "250 closing connection" {
				os.Exit(0)
			} else if strings.HasPrefix(text, "250 ") {
				fmt.Println(strings.TrimPrefix(text, "250 "))
				break
			} else if strings.HasPrefix(text, "251 ") {
				fmt.Println(strings.TrimPrefix(text, "251 "))
				break
			} else if strings.HasPrefix(text, "252 ") {
				fmt.Println(strings.TrimPrefix(text, "252 "))
				break
			} else if strings.HasPrefix(text, "451 ") {
				fmt.Println(strings.TrimPrefix(text, "451 "))
				break
			} else if strings.HasPrefix(text, "500 ") {
				fmt.Println(strings.TrimPrefix(text, "500 "))
				break
			} else if strings.HasPrefix(text, "510 ") {
				fmt.Println(strings.TrimPrefix(text, "510 "))
				break
			} else if strings.HasPrefix(text, "511 ") {
				fmt.Println(strings.TrimPrefix(text, "511 "))
				break
			} else if strings.HasPrefix(text, "512 ") {
				fmt.Println(strings.TrimPrefix(text, "512 "))
				break
			} else if strings.HasPrefix(text, "513 ") {
				fmt.Println(strings.TrimPrefix(text, "513 "))
				break
			} else if strings.HasPrefix(text, "514 ") {
				fmt.Println(strings.TrimPrefix(text, "514 "))
				break
			} else if strings.HasPrefix(text, "515 ") {
				fmt.Println(strings.TrimPrefix(text, "515 "))
				break
			} else if strings.HasPrefix(text, "550 ") {
				fmt.Println(strings.TrimPrefix(text, "550 "))
				break
			} else if strings.HasPrefix(text, "551 ") {
				fmt.Println(strings.TrimPrefix(text, "551 "))
				break
			} else if strings.HasPrefix(text, "552 ") {
				fmt.Println(strings.TrimPrefix(text, "552 "))
				break
			} else if strings.HasPrefix(text, "553 ") {
				fmt.Println(strings.TrimPrefix(text, "553 "))
				break
			} else if strings.HasPrefix(text, "554 ") {
				fmt.Println(strings.TrimPrefix(text, "554 "))
				break
			} else if strings.HasPrefix(text, "555 ") {
				fmt.Println(strings.TrimPrefix(text, "555 "))
				break
			} else if strings.HasPrefix(text, "650 ") {
				fmt.Println(strings.TrimPrefix(text, "650 "))
				break
			} else {
				fmt.Println(text)
			}
		}
	}

	// Send non-interactive command and exit
	if *cmdPtr != "" {
		fmt.Fprintf(conn, *cmdPtr+"\n")
		receive()	
		fmt.Fprintf(conn, "quit\n")
		os.Exit(0)
	}


	// Announce Tor version as banner
	fmt.Fprintf(conn, "getinfo version\n")
	text, err = bufio.NewReader(conn).ReadString('\n')
	if err != nil {os.Exit(4)}
	fmt.Print("Tor "+strings.TrimPrefix(text, "250-version="))
	
	for {
		text = ""

		// Command prompt
		fmt.Print("> ")
		text, err = console.ReadString('\n')
		if err != nil {os.Exit(5)}

		// Internal TorMon commands
		switch strings.ToLower(strings.TrimRight(text, "\r\n")) {
			case "help":
				fmt.Print("getinfo config/names\ngetinfo signal/names\ngetinfo info/names\nquit\n")
				continue
		}

		// Send Tor command
		fmt.Fprintf(conn, text)

		// Handle special Tor commands
		switch strings.ToLower(strings.TrimRight(text, "\r\n")) {
			case "signal shutdown":
				os.Exit(0)
			case "signal halt":
				os.Exit(0)
			case "signal int":
				os.Exit(0)
			case "signal term":
				os.Exit(0)
		}

		// Receive command output
		receive()

	}
}