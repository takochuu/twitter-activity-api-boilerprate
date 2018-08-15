package response

type Response interface {
	JSON(data interface{})
	GRPC(data interface{})
}
