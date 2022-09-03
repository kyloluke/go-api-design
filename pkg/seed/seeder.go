package seed

import (
	"gohub/pkg/console"
	"gohub/pkg/database"
	"gorm.io/gorm"
)

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

// SetRunOrder 设置需要优先执行的 seeder
func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if sdr.Name == name {
			return sdr
		}
	}

	return Seeder{}
}

func RunAll() {
	excuted := make(map[string]string)
	// 先运行 ordered 的
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)

		if len(sdr.Name) > 0 {
			console.Warning("Running Ordered Seeder: " + sdr.Name)
			sdr.Func(database.DB) // 注意这里不能是 *gorm.DB 因为
			excuted[name] = name
		}
	}
	// 在运行剩下的
	for _, sdr := range seeders {
		// 过滤已运行
		if _, ok := excuted[sdr.Name]; !ok {
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder 运行单个 Seeder
func RunSeeder(sdr Seeder) {
	sdr.Func(database.DB)
}
