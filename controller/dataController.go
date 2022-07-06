package controller

import (
	"log"
	"net/http"

	"github.com/GustavoFreitas22/the_guardian/model"
	"github.com/GustavoFreitas22/the_guardian/service"
	"github.com/gin-gonic/gin"
)

func GetInfoDataById(c *gin.Context) {

	id := c.Params.ByName("id")

	response, err := service.GetInfoById(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Erro ao buscar dados",
		})
		return
	}

	if response.ID == 0 {
		c.JSON(404, gin.H{
			"Error": "Secret não encontrado",
		})
		return
	}

	c.JSON(200, response)
}

func GetDataByContext(c *gin.Context) {
	context := c.Params.ByName("context")

	response := service.GetSecretByContext(context)

	if response[0] == nil {
		c.JSON(404, gin.H{
			"Erro": "Não foram encontrados secrets com esse context",
		})
		return
	}

	c.JSON(200, response)
}

func PostInfoData(c *gin.Context) {
	var newData model.Data

	err := c.ShouldBindJSON(&newData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "Erro ao inserir dados" + err.Error(),
		})
		return
	}

	err = service.EncriptDataAndInsert(&newData)
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(201, gin.H{
		"Message": "Dado inserido com sucesso",
	})
}

func EditInfoData(c *gin.Context) {
	var editData model.Data

	err := c.ShouldBindJSON(&editData)
	if err != nil {
		log.Println(err)
		return
	}

	response := service.DesencryptAndEditData(&editData)

	if response.ID == 0 {
		c.JSON(404, gin.H{
			"Message": "Secret não encontrado",
		})
		return
	}

	c.JSON(200, response)
}

func DeleteInfoData(c *gin.Context) {
	id := c.Params.ByName("id")

	service.DeleteSecretById(id)

}
