package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Inicializando a rota
	router := gin.Default()

	// Iniciando as rotas da aplicação
	initializeRoutes(router)

	// Startando servidor
	router.Run(":8080")
}
