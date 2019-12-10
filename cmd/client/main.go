package main

import (
	"context"
	"fmt"
	"github.com/tchaudhry91/bottle/bottle"
	"github.com/tchaudhry91/bottle/crate"
	"github.com/tchaudhry91/bottle/crate/pb"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"os"
	"strings"
)

func main() {
	var serverAddr string
	var mode string
	var id string

	app := &cli.App{
		Name:  "bottle",
		Usage: "Store data from pipes",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "server",
				Aliases:     []string{"s"},
				Value:       "",
				EnvVars:     []string{"BOTTLE_SERVER"},
				Usage:       "Address of the remote server",
				Destination: &serverAddr,
			},
			&cli.StringFlag{
				Name:        "mode",
				Aliases:     []string{"m"},
				Value:       "list",
				Usage:       "Mode to operate in, fill/drain/pour",
				EnvVars:     []string{"BOTTLE_MODE"},
				Destination: &mode,
			},
			&cli.StringFlag{
				Name:        "id",
				Aliases:     []string{"i"},
				Value:       "clip",
				EnvVars:     []string{"BOTTLE_ID"},
				Usage:       "ID of the bottle to operate on",
				Destination: &id,
			},
		},
		Action: func(c *cli.Context) error {
			conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
			if err != nil {
				panic(err)
			}
			// Build Client
			client := pb.NewCrateClient(conn)

			switch strings.ToLower(mode) {
			case "fill":
				return fill(id, client)
			case "pour":
				return pour(id, client)
			case "drain":
				return drain(id, client)
			case "list":
				return list(client)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func fill(id string, client pb.CrateClient) error {
	b := bottle.NewBottle(id)
	err := b.Fill(os.Stdin)
	if err != nil {
		fmt.Printf("Failed to store bottle: %s", err)
		return err
	}
	resp, err := client.Store(context.Background(), &pb.StoreRequest{Bottle: crate.BottleToPB(b)})
	if err != nil {
		fmt.Printf("Failed to store bottle: %s", err)
		return err
	}
	if resp.Err != "" {
		fmt.Printf("Failed to store bottle: %s", resp.Err)
		return fmt.Errorf(resp.Err)
	}
	return nil
}

func pour(id string, client pb.CrateClient) error {
	var b *bottle.Bottle
	resp, err := client.Pour(context.Background(), &pb.GetRequest{Id: id})
	if err != nil {
		fmt.Printf("Failed to fetch bottle: %s", err)
		return err
	}
	if resp.Err != "" {
		fmt.Printf("Failed to fetch bottle: %s", resp.Err)
		return fmt.Errorf(resp.Err)
	}
	b = crate.PBToBottle(resp.Bottle)
	err = b.Pour(os.Stdout)
	if err != nil {
		fmt.Printf("Failed to pour bottle: %s", err)
		return err
	}
	return nil
}

func drain(id string, client pb.CrateClient) error {
	var b *bottle.Bottle
	resp, err := client.Drain(context.Background(), &pb.GetRequest{Id: id})
	if err != nil {
		fmt.Printf("Failed to fetch bottle to drain: %s", err)
		return err
	}
	if resp.Err != "" {
		fmt.Printf("Failed to fetch bottle to drain: %s", resp.Err)
		return fmt.Errorf(resp.Err)
	}
	b = crate.PBToBottle(resp.Bottle)
	err = b.Drain(os.Stdout)
	if err != nil {
		fmt.Printf("Failed to drain bottle: %s", err)
		return fmt.Errorf(resp.Err)
	}
	return nil
}

func list(client pb.CrateClient) error {
	resp, err := client.List(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Printf("Failed to fetch bottle list: %s", err)
		return err
	}
	if resp.Err != "" {
		fmt.Printf("Failed to fetch bottle list: %s", resp.Err)
		return fmt.Errorf(resp.Err)
	}
	for _, b := range resp.Bottles {
		fmt.Println(b)
	}
	return nil
}
