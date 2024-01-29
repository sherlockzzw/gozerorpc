package job

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-template/internal/svc"
	"context"
	"github.com/spf13/cobra"
	"os"
)

const (
	codeFailure = 1
)

var (
	svcCtx *svc.ServiceContext

	rootCmd = &cobra.Command{
		Use:   "cron",
		Short: "exec rpc template cron job",
		Long:  "exec rpc template cron job",
	}

	detailCount = &cobra.Command{
		Use:   "detailCount",
		Short: "输出用户详细总数",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			return registerDetailCountJob(svcCtx).run(ctx)
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(codeFailure)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(detailCount)
}

func initConfig() {
	c := config.Init()
	svcCtx = svc.NewJobContext(*c)
}
