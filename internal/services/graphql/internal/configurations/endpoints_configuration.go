package configurations

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/peygy/nektoyou/internal/pkg/grpc"
	ginServer "github.com/peygy/nektoyou/internal/pkg/gin"
	"github.com/peygy/nektoyou/internal/services/graphql/graph"
)

func InitEndpoints(eng *ginServer.GinServer, grpcPull *grpc.GrpcPull) {
	eng.Engine.Use(cors.Default())
	routeGroup1 := eng.Engine.Group("/graphql") 
	{
		routeGroup1.POST("/hello", graphqlHandler(grpcPull.Services))
	}
}

// Defining the Graphql handler
func graphqlHandler(grpcServices []grpc.GrpcService) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{GrpcServices: grpcServices}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
