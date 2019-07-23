package main

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//defer conn.Close()
	//c := pb.NewUserCenterServiceClient(conn)
	//
	//// Contact the server and print out its response.
	//name := defaultName
	//if len(os.Args) > 1 {
	//	name = os.Args[1]
	//}
	//r, err := c.RegisterUser(context.Background(), &pb.RegisterUserRequest{NickName: name})
	//if err != nil {
	//	log.Fatalf("could not greet: %v", err)
	//}
	//log.Printf("Greeting: %s", r.UserToken)
	data := make(map[string]map[string]string)
	data["keyword1"] = map[string]string{"value": "1"}
	data["keyword2"] = map[string]string{"value": "1"}

}
