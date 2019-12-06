package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/peterbourgon/ff"
	"github.com/tchaudhry91/bottle/bottle"
	"github.com/tchaudhry91/bottle/crate"
	"github.com/tchaudhry91/bottle/crate/pb"
	"google.golang.org/grpc"
	"os"
)

func main() {
	fs := flag.NewFlagSet("bottle", flag.ExitOnError)
	var (
		serverAddr = fs.String("server-addr", "127.0.0.1:14999", "crate service to connect to")
		id         = fs.String("id", "", "bottle id")
		mode       = fs.String("mode", "store", "Mode to operate under:drain/pour/store")
	)
	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarNoPrefix())
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	// Build Client
	client := pb.NewCrateClient(conn)

	switch *mode {
	case "store":
		b := bottle.NewBottle(*id)
		err = b.Fill(os.Stdin)
		if err != nil {
			fmt.Printf("Failed to store bottle: %s", err)
			os.Exit(1)
		}
		resp, err := client.Store(context.Background(), &pb.StoreRequest{Bottle: crate.BottleToPB(b)})
		if err != nil {
			fmt.Printf("Failed to store bottle: %s", err)
			os.Exit(1)
		}
		if resp.Err != "" {
			fmt.Printf("Failed to store bottle: %s", resp.Err)
			os.Exit(1)
		}
	case "pour":
		var b *bottle.Bottle
		resp, err := client.Pour(context.Background(), &pb.GetRequest{Id: *id})
		if err != nil {
			fmt.Printf("Failed to fetch bottle: %s", err)
			os.Exit(1)
		}
		if resp.Err != "" {
			fmt.Printf("Failed to fetch bottle: %s", resp.Err)
			os.Exit(1)
		}
		b = crate.PBToBottle(resp.Bottle)
		err = b.Pour(os.Stdout)
		if err != nil {
			fmt.Printf("Failed to pour bottle: %s", err)
			os.Exit(1)
		}
	case "drain":
		var b *bottle.Bottle
		resp, err := client.Drain(context.Background(), &pb.GetRequest{Id: *id})
		if err != nil {
			fmt.Printf("Failed to fetch bottle to drain: %s", err)
			os.Exit(1)
		}
		if resp.Err != "" {
			fmt.Printf("Failed to fetch bottle to drain: %s", resp.Err)
			os.Exit(1)
		}
		b = crate.PBToBottle(resp.Bottle)
		err = b.Drain(os.Stdout)
		if err != nil {
			fmt.Printf("Failed to drain bottle: %s", err)
			os.Exit(1)
		}
	}
}
