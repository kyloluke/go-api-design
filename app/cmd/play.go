package cmd

import (
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	// 存进redis中
	//redis.Redis.Set("hello", "hi from redis", 10*time.Second)
	//
	//console.Success(redis.Redis.Get("hello"))
}
