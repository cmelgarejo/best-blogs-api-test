package bootstrap

import (
	"go-rest-blog/service"
)

func Init(port int) error {
	api := service.NewRestApiService()
	return api.ServeContent(port)
}
