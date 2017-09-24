package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/brotherlogic/goserver"
	"google.golang.org/grpc"
)

//Server main server type
type Server struct {
	*goserver.GoServer
}

// DoRegister Registers this server
func (s *Server) DoRegister(server *grpc.Server) {
	//WE do nothing
}

// ReportHealth Determines if the server is healthy
func (s *Server) ReportHealth() bool {
	return true
}

// Mote promotes this server
func (s *Server) Mote(master bool) error {
	//No mote
	return nil
}

func crash() {
	//time.Sleep(time.Hour)
	panic("Whoopsie")
}

func main() {
	var quiet = flag.Bool("quiet", false, "Show all output")
	flag.Parse()

	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	server := &Server{&goserver.GoServer{}}
	server.Register = server
	server.RegisterServer("crasher", false)
	server.Log("Beerserver has started!")
	server.RegisterServingTask(crash)
	server.Serve()
}
