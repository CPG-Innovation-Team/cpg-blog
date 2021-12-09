package vo

import "cpg-blog/internal/notify/model"

/**
  @author: chenxi@cpgroup.cn
  @date:2021/12/8
  @description: 通知返回参数
**/

// AddNotificationVO 添加通知返回参数
type AddNotificationVO struct {
	Id int
}

type SystemNotificationVO struct {
	NotificationList []model.Notify
}