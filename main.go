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
	grpclog.Println("executing pipeline")
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
		grpclog.Printf("response: code: %v, text:%v", chunk.Status, chunk.Text)
	}
	grpclog.Println("end of pipeline")
}

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("error during connecting to grpc", err)
		return
	}
	// cmd := strings.Join(os.Args[1:], " ")
	defer conn.Close()
	client := pb.NewPipelineExecutorClient(conn)

	// execPipeline(client, &pb.ExecutePipeline{Id: 1, Pipeline: []byte(cmd)})
	file, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	execPipeline(client, &pb.ExecutePipeline{Id: 1, Pipeline: file})
}
