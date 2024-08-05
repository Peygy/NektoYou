package configurations

import (
	"fmt"
	"sort"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	ginServer "github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/services/graphql/graph"
	pb "github.com/peygy/nektoyou/internal/services/graphql/internal/grpc_client/proto"
)

func InitEndpoints(eng *ginServer.GinServer, grpcPull *grpc.GrpcPull) {
	eng.Engine.POST("/query", graphqlHandler())
	eng.Engine.GET("/", playgroundHandler())
	eng.Engine.GET("/hello", helloHandler(grpcPull.Services))
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func helloHandler(grpcServices []grpc.GrpcService) gin.HandlerFunc {
	authConnIdx := sort.Search(len(grpcServices), func(i int) bool { return grpcServices[i].Name == "auth_service" })
	cl := pb.NewGreeterClient(grpcServices[authConnIdx].Conn)

	return func(c *gin.Context) {
		r, err := cl.SayHello(c, &pb.HelloRequest{Word: "HIIIII"})
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Print("MESSAGE: " + r.GetMessage())
	}
}
