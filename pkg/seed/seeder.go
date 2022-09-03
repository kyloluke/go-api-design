package seed

import "gorm.io/gorm"

type SeederFunc func(*gorm.DB)

type Seeder struct {
	Func SeederFunc
	Name string // 文件名称
}

// 按顺序执行的 Seeder 数组
// 支持一些必须按顺序执行的 seeder，例如 topic 创建的
// 时必须依赖于 user, 所以 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

// 存放所有 Seeder
var seeders []Seeder

func Add(name string, fn SeederFunc) {

	seeders = append(seeders, Seeder{
		Func: fn,
		Name: name,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}
