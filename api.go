package basego

import "github.com/gin-gonic/gin"



type Api struct {
	engine *gin.Engine
	Port   string
	repo   Repository
	prefix string
}

func NewAPI(repo Repository, prefix, port string) *Api {
	return &Api{
		engine: gin.Default(),
		Port:   port,
		prefix: prefix,
		repo:   repo,
	}
}

func (api *Api) registerEndpoints() {

	r := api.engine.Group(api.prefix)
	api.registerUserEndpoints(r)

}

func (api *Api) Launch() error {
	api.registerEndpoints()
	return api.engine.Run(api.Port)
}



