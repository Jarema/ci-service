package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	pb "github.com/jarema/ci-service/executor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func execPipeline(client pb.PipelineExecutorClient, pipeline *pb.ExecutePipeline) {
	fmt.Println("executing pipeline")
	stream, err := client.Execute(context.Background(), pipeline)
	if err != nil {
		grpclog.Fatalf("error: %v", err)
	}
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatalf("error! %v\n", err)
		}
		fmt.Printf("%s\n", chunk.Text)
	}
	fmt.Println("end of pipeline")
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("error during connecting to grpc", err)
		return
	}
	defer conn.Close()
	client := pb.NewPipelineExecutorClient(conn)

	file, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	execPipeline(client, &pb.ExecutePipeline{Id: 1, Pipeline: file})
}
