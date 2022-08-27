package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/models"
	"github.com/silverswords/sand/services"
)

func RegisterProduct(r gin.IRouter) {
	r.GET("insert", insertProduct)
}

func insertProduct(ctx *gin.Context) {
	product := models.Product{
		ProId:       "123",
		VersionId:   "123",
		MainTitle:   "123",
		Subtitle:    "123",
		MainPicture: "123",
		Price:       12.2,
		Status:      12,
		Active:      1,
	}

	err := services.InsertProduct(product)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": http.StatusBadRequest, "error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
