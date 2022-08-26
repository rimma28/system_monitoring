package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"os/linux"
	"protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	protobuf.UnimplementedMonitorServer
	rootFsPrefix string
}

func main() {
	rootFsPrefix := flag.String("r", "", "path where host root FS is mounted")
	lis, err := net.Listen("tcp", "localhost:8088")
	if err != nil {
		fmt.Println("cannot listen TCP:", err.Error())
	}
	s := grpc.NewServer()
	protobuf.RegisterMonitorServer(s, &server{
		rootFsPrefix: *rootFsPrefix,
	})
	fmt.Println("server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Println("cannot handle reguest:", err.Error())
	}
}

// GetStats implements protobuf.MonitorServer
func (s *server) GetStats(_ *protobuf.Settings, srv protobuf.Monitor_GetStatsServer) error {
	fmt.Println("got new client")
	for {
		cpu, err := linux.CalcCPUUsage(s.rootFsPrefix)
		if err != nil {
			return status.Errorf(codes.Aborted, "server error: %s", err.Error())
		}

		devs, err := linux.CalcDevStats(s.rootFsPrefix)
		if err != nil {
			return status.Errorf(codes.Aborted, "server error: %s", err.Error())
		}

		fsystems, err := linux.CalcFsUtilization(s.rootFsPrefix)
		if err != nil {
			return status.Errorf(codes.Aborted, "server error: %s", err.Error())
		}

		cpuStats := protobuf.CPUstats{
			La:   cpu.LA,
			Sys:  cpu.SysUsagePercent,
			Usr:  cpu.UsrUsagePercent,
			Idle: cpu.IdlePercent,
		}

		stats := protobuf.Stats{
			CPUstats: &cpuStats,
		}

		for _, dev := range devs {
			devStats := protobuf.DevStats{
				Name:  dev.Name,
				Tps:   dev.TransPS,
				Read:  dev.ReadPS,
				Write: dev.WritePS,
			}

			stats.DevStats = append(stats.DevStats, &devStats)
		}

		for _, fs := range fsystems {
			fsStats := protobuf.FsStats{
				Name:         fs.Name,
				Bytes:        fs.UsedGBytes,
				BytesPercent: fs.UsedStoragePercent,
				Inode:        fs.UsedInodes,
				InodePercent: fs.UsedInodesPercent,
			}

			stats.FsStats = append(stats.FsStats, &fsStats)
		}

		if err := srv.Send(&stats); err != nil {
			return status.Errorf(codes.Aborted, "sending message error: %s", err.Error())
		}

		time.Sleep(time.Second)
	}
}
