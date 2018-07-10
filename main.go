package main

import(
    "net"
    "fmt"
    "bufio"
    "log"
    "strings"
)

const (
    HOST = "localhost"
    PORT = "8081"
    TYPE = "tcp"
)

func main() {

  fmt.Printf("Starting server on %s...\n", PORT)
  ln, err := net.Listen(TYPE, HOST + ":" + PORT)

  if err != nil {
    panic(err)
  }
  defer ln.Close()

  for {
    conn, err := ln.Accept()
    if err != nil {
        log.Println(err.Error())
        continue
    }
    go httpHandler(conn)
  }
}

func httpHandler(conn net.Conn) {
    defer conn.Close()
    scanner := bufio.NewScanner(conn)

    fmt.Println("-------------------------------")

    i := 0
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)

        if line == "" {
            break
        }

        if i == 0 {
            requestHandler(conn, line)
        }
        i++
    }
}

func requestHandler(conn net.Conn, line string) {
    method := strings.Fields(line)[0]
    url := strings.Fields(line)[1]
    body := "Hello From Server.."

    fmt.Println("[Server] Responding to " + method + " " + url)


    fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\r\n")
    fmt.Fprint(conn, body)
}
