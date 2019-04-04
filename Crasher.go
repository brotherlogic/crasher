package main

import (
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/brotherlogic/goserver"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pbgs "github.com/brotherlogic/goserver/proto"
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

// Shutdown shutsdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

// Mote promotes this server
func (s *Server) Mote(ctx context.Context, master bool) error {
	//No mote
	return nil
}

//GetState gets the state of the server
func (s *Server) GetState() []*pbgs.State {
	return []*pbgs.State{}

}

func crash(ctx context.Context) error {
	time.Sleep(time.Minute * 5)
	panic("Whoopsie")
	return nil
}

func main() {
	var quiet = flag.Bool("quiet", true, "Show all output")
	flag.Parse()

	if *quiet {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	server := &Server{&goserver.GoServer{}}
	server.Register = server
	server.PrepServer()
	server.RegisterServer("crasher", false)
	server.Log("Beerserver has started!")
	server.RegisterServingTask(crash, "crash")
	server.Serve()
}
