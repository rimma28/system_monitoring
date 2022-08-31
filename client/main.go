package main

import (
	"context"
	"fmt"
	"io"

	"protobuf"

	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8088", grpc.WithInsecure())
	if err != nil {
		fmt.Println("cannot connect to gRPC server:", err.Error())
	}
	defer conn.Close()
	c := protobuf.NewMonitorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := c.GetStats(ctx, &protobuf.Settings{})
	if err != nil {
		fmt.Println("cannot open stream:", err.Error())
	}

	for {
		r, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("stream is closed, finishing session")
			return
		}
		if err != nil {
			fmt.Println("error during streaming:", err.Error())
			return
		}

		fmt.Println("CPU statistics:")
		fmt.Println("load average:", r.CPUstats.La)
		fmt.Printf("Загрузка CPU usr: %.2f%%, sys: %.2f%%, idle: %.2f%%\n", r.CPUstats.Usr, r.CPUstats.Sys, r.CPUstats.Idle)

		fmt.Println("Загрузка дисков:")
		for _, dev := range r.DevStats {
			fmt.Printf("Name: %s, Transfers per sec: %.3f, Read: %.3f Kbps, Write: %.3f Kbps\n",
				dev.Name,
				dev.Tps,
				dev.Read,
				dev.Write,
			)
		}

		fmt.Println("Информация о дисках по каждой файловой системе:")
		for _, fs := range r.FsStats {
			fmt.Printf("Name: %s, Used storage: %d Gb, Used storage persentage: %.2f%%, Used inodes: %d, Used inodes persentage: %.2f%%\n",
				fs.Name,
				fs.Bytes,
				fs.BytesPercent,
				fs.Inode,
				fs.InodePercent,
			)
		}
	}
}
