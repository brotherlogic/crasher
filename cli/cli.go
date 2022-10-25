package main

import (
	"log"
	"strconv"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/crasher/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	host, port, err := utils.Resolve("crasher", "crasher-cli")
	if err != nil {
		log.Fatalf("Unable to reach crasher: %v", err)
	}
	conn, err := grpc.Dial(host+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCrashServiceClient(conn)
	ctx, cancel := utils.BuildContext("recordader-cli", "recordadder")
	defer cancel()

	client.Crash(ctx, &pb.CrashRequest{})

}
