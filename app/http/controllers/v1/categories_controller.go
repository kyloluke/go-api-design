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
	// 1. 验证分页参数的有效性
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	// 2. 返回数据和分页信息
	data, pager := category.Paginate(c, 10)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"data":  data,
			"pager": pager,
		},
	})
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
		return
	}
	//response.Abort500(c)
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "分类编辑失败",
	})
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {
	categoryModel := category.Get(c.Param("id"))

	if categoryModel.ID == 0 {
		//response.Abort404()
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "数据不存在，请确认请求信息是否正确",
		})
		return
	}

	rowsAffected := categoryModel.Delete()

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "删除失败，请稍后再试",
	})
}
