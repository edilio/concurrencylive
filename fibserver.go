package main

import (
   "net"
   "fmt"
   "strconv"
   "io"
   "bufio"
   "runtime"
   "funct"
)


func fib(n int)  int {
   if n <= 2 {
      return 1
   } else {
      return fib(n-1) + fib(n-2)
   }
}



func launchServer(dnc_host, dnc_port string) {
   msg := fmt.Sprintf("Launching server...on %s:%s", dnc_host, dnc_port)
   fmt.Println(msg)
   // listen on all interfaces
   ln, _ := net.Listen("tcp", dnc_host + ":" + dnc_port)

   // run loop forever (or until ctrl-c)
   for {
      conn, _ := ln.Accept()

      go handleConnection(conn, fib)

   }
}


func handleConnection(conn net.Conn, fib funct.Formula) {
   // will listen for message to process ending in newline (\n)

   for {
      message, err := bufio.NewReader(conn).ReadString('\r')
      if err == io.EOF {
         conn.Close()
         return
      }
      if len(message) > 0 {
         //fmt.Println("======\n")
         //fmt.Print("Message Received:", string(message))
         if message == "q\n" {
            fmt.Println("received a quit message")
            conn.Close()
            return
         }
         i, err := strconv.Atoi(message[0:len(message)-1])
         if err != nil {
            fmt.Println(err)
            fmt.Println(message[0:len(message)-1])
         } else {
            r := fib(i)
            result := fmt.Sprintf("%d\n", r)
            conn.Write([]byte(result))
         }
      }

   }

}


func main() {
   runtime.GOMAXPROCS(runtime.NumCPU())
   launchServer("127.0.0.1", "25000")
}
