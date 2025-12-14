package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "9090"
	}

	logFile, err := os.OpenFile("/var/log/goservice.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error escuchando puerto %s: %v", port, err)
	}
	log.Printf("Servicio escuchando en el puerto %s", port)

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("Error de conexión:", err)
				continue
			}
			log.Printf("Nueva conexión desde %s", conn.RemoteAddr())
			fmt.Fprintln(conn, "Servicio activo")
			conn.Close()
		}
	}()

	<-ctx.Done()
	log.Println("Recibida señal de apagado, cerrando servicio...")
	listener.Close()
	time.Sleep(2 * time.Second)
}
