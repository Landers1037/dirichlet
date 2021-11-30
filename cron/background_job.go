/*
Project: dirichlet background_job.go
Created: 2021/11/30 by Landers
*/

package cron

import (
	"fmt"

	"github.com/landers1037/dirichlet/app/app_manager"
	"github.com/landers1037/dirichlet/logger"
)

// 启动时执行的轮询任务 用于随时刷新持久化数据
// 持久化数据用于恢复
const (
	Duration_DBSaver = 3600
	Duration_APPSync = 60
)

func InitBackgroundJobs() {
	AddJobDBSaver()
	AddJobAPPSync()
}

// AddJobDBSaver 数据库持久化刷新
func AddJobDBSaver() {
	logger.Logger.Info("job: database sync start")
	AddTicker(Duration_DBSaver, func() {

	})
}

// AddJobAPPSync app配置文件同步
// 同步顺序
func AddJobAPPSync() {
	logger.Logger.Info("job: app config sync start")
	AddTicker(Duration_APPSync, func() {
		app_manager.APPManager.APPManagerMap.Range(func(key, value interface{}) bool {
			app := value.(app_manager.App)
			_, err := app.Sync()
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("job app config sync failed: %s", err.Error()))
			}
			return true
		})
	})
}