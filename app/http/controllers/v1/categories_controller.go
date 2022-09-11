package v1

import (
	"github.com/gin-gonic/gin"
	"gohub/app/models/category"
	"gohub/app/requests"
	"gohub/pkg/response"
	"net/http"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {

}

func (ctrl *CategoriesController) Show(c *gin.Context) {

}

func (ctrl *CategoriesController) Store(c *gin.Context) {
	request := requests.CategoryRequest{}

	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}

	categoryModel.Create()

	if categoryModel.ID > 0 {
		//response.Created(c, categoryModel)
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"data":    categoryModel,
		})
	} else {
		//response.Abort500(c, "创建失败，请稍后尝试~")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "创建失败，请稍后尝试~",
		})
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffect := categoryModel.Save()

	if rowsAffect > 0 {
		//response.Data(c, categoryModel)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    categoryModel,
		})
	} else {
		//response.Abort500(c)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "分类编辑失败",
		})
	}
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {

}
