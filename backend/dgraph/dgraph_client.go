package dgraph

import (
	"google.golang.org/grpc"
	"log"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/dgraph-io/dgo"
)

func CreateClient() (dgo.Dgraph, *grpc.ClientConn) {
	conn, err := grpc.Dial("127.0.0.1:9080", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	dc := api.NewDgraphClient(conn)
	dg := dgo.NewDgraphClient(dc)

	return *dg, conn
}