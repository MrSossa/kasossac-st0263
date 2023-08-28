package entities

type (
	AppType  string
	AppTypes []AppType
)

const (
	GRPCServerApp AppType = "gRPC-server"
	MOMServerApp  AppType = "mom-server"
	ApiGWApp      AppType = "api-gw"
)
