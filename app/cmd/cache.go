package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gohub/pkg/cache"
	"gohub/pkg/console"
	"gohub/pkg/logger"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runCacheClear,
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "forget",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	CmdCache.AddCommand(CmdCacheClear, CmdCacheForget)

	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "key of the cache")
	err := CmdCacheForget.MarkFlagRequired("key")
	logger.LogIf(err)
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}
