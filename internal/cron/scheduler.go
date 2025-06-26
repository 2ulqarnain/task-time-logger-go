package cron

import (
	"task-time-logger-go/internal/cron/jobs"
	"task-time-logger-go/internal/logger"

	"github.com/robfig/cron/v3"
)

type CronManager struct {
	cron *cron.Cron
}

func NewCronManager() *CronManager {
	cm := &CronManager{
		cron: cron.New(),
	}
	cm.registerJobs()
	count := cm.getJobsCount()
	logger.AppLogger.Printf("%d Cron Jobs Registered !", count)
	return cm
}

func (cm *CronManager) registerJobs() {
	cm.cron.AddFunc("30 18 * * *", jobs.EodDurationJob)
}

func (cm *CronManager) getJobsCount() int {
	return len(cm.cron.Entries())
}

func (cm *CronManager) Start() {
	cm.cron.Start()
}

func (cm *CronManager) Stop() {
	cm.cron.Stop()
}
