package main

import (
	"github.com/gliderlabs/ssh"
	"io"
	"log"
	"log/syslog"
	"os"
)

// this is the function called when ssh client requests authentication
// here we log attacker credentials on purpose
func authHandler(ctx ssh.Context, password string) bool {
	log.Printf("User: %s connecting from %s with password: %s\n", ctx.User(), ctx.RemoteAddr(), password)
	return true
}

// his session will end very quickly...
// to free up resources we only send a message and quit
func sessionHandler(s ssh.Session) {
	io.WriteString(s, "Welcome!\n")
}

func main() {

	// Configure logger to write to the syslog
	logwriter, e := syslog.New(syslog.LOG_INFO, os.Args[0])
	if e == nil {
		log.SetOutput(logwriter)
	}

	s := &ssh.Server{
		Addr:            ":2222",
		Handler:         sessionHandler,
		PasswordHandler: authHandler,
	}
	log.Println("starting ssh server on port 2222...")
	log.Fatal(s.ListenAndServe())
}
