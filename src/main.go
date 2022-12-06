package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Body struct {
	//obrigatório, numerico, número > 0
	Id int `json:"id" binding:"required,numeric,gt=0"`
	//obrigatório
	Name string `json:"name" binding:"required"`
	//obrigatório, >=10 e <=1000
	Price uint `json:"price" binding:"required,numeric,gte=10,lte=1000"`
}

func main() {
	engine := gin.New()
	engine.POST("/test", func(context *gin.Context) {

		body := Body{}

		if err := context.ShouldBindJSON(&body); err != nil {

			context.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					//retorna a mensagem dos erros da validação
					"message": err.Error(),
				})
			return
		}
		context.JSON(http.StatusAccepted, &body)
	})
	engine.Run(":3000")
}

//Exeplo do body request POST
// {
// 	"id":1,
// 	"name":"teste",
// 	"price":300
// }
