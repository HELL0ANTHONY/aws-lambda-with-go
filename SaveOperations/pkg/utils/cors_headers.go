package utils

const (
	allowHeaders = "Access-Control-Allow-Headers"
	allowMethods = "Access-Control-Allow-Methods"
	allowOrigin  = "Access-Control-Allow-Origin"
	contentType  = "Content-Type"
)

func CORSHeaders(origin string) map[string]string {
	return map[string]string{
		allowHeaders: "Access-Control-Allow-Origin, Access-Control-Allow-Methods, Content-Type",
		allowMethods: "OPTIONS, POST",
		allowOrigin:  origin,
		contentType:  "application/json",
	}
}
