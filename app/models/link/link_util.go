package link

import (
	"fmt"
	"gohub/pkg/app"
	"gohub/pkg/cache"
	"gohub/pkg/database"
	"gohub/pkg/helpers"
	"gohub/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

// AllCached 从cache中读取缓存数据，如没有缓存数据则查库读取
func AllCached() (links []Link) {
	cacheKey := "link:all"

	expireTime := 120 * time.Minute

	cache.GetObject(cacheKey, &links)
	fmt.Printf("缓存获取到的连接长度为：%#v", len(links))

	if helpers.Empty(links) {
		links = All()
		if helpers.Empty(links) {
			return
		}

		cache.Set(cacheKey, links, expireTime)
	}
	return
}
