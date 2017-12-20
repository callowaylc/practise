package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
)

func handleConnection(conn net.Conn) {
  defer conn.Close()

  reader := bufio.NewReaderSize(conn, 2048)
  tmpChan := make(chan []byte, 20140)

  if tcp, ok := conn.(*net.TCPConn); ok {
    tcp.SetNoDelay(true)
  }

  go func() {
    for {
      data := <-tmpChan
      remain := len(tmpChan)

      for i := 0; i < remain; i++ {
        left := <-tmpChan
        data = append(data, left...)
      }

      size, err1 := conn.Write(data)
      if err1 != nil || size != len(data) {
        fmt.Println("write error: ", err1, size)
        return
      }
    }

  }()

  buf := make([]byte, 8192)
  for {
    n, err := reader.Read(buf)
    if err != nil {
      fmt.Println("read error: ", err)
      return
    }

    for pos := 0; pos < n; pos += 64 {
      end := pos + 64
      if end > n {
        end = n
      }
      send := []byte{}
      send = append(send, buf[pos:end]...)
      tmpChan <- send
    }
  }
}

func main() {
  ln, err := net.Listen("tcp", ":3050")
  if err != nil {
    panic(err)
  }

  fmt.Println("listen 3050 ok")

  for {
    conn, err := ln.Accept()
    if err != nil {
      log.Fatal("get client connection error: ", err)
    }

    go handleConnection(conn)
  }
}