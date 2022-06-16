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

func (s SystemNotificationVO) Len() int{
	return len(s.NotificationList)
}

func (s SystemNotificationVO) Less(i, j int) bool{
	return s.NotificationList[i].EndTime.Unix() < s.NotificationList[j].EndTime.Unix()
}

func (s SystemNotificationVO) Swap(i, j int) {
	s.NotificationList[i], s.NotificationList[j] = s.NotificationList[j], s.NotificationList[i]
}