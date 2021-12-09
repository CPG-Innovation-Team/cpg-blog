package impl

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/notify/model"
	"cpg-blog/internal/notify/model/dao"
	"cpg-blog/internal/notify/qo"
	"cpg-blog/internal/notify/vo"
	"cpg-blog/pkg/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"time"
)

/**
  @author: chenxi@cpgroup.cn
  @date:2021/12/6
  @description: 通知服务接口实现
**/

type Notify struct{}

func (Notify) AddNotification(ctx *gin.Context, query *qo.AddNotificationQO) {
	//转换请求参数
	util.JsonConvert(ctx, query)

	//判断通知类型，若类型为4则uid默认为0;其余情况需将UID转换进行校验
	if query.Type == cpgConst.FOUR && len(query.Uid) > 1 && query.Uid[0] != cpgConst.ZERO {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//将uid列表、content、state、notify_date写入model
	notifyDao := new(dao.NotifyDao)
	uidByte, _ := json.Marshal(query.Uid)
	_ = copier.Copy(notifyDao, query)
	notifyDao.Uid = string(uidByte)

	//model写入数据库
	id, err := notifyDao.Creat(ctx)

	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, vo.AddNotificationVO{Id: id})
}

func (Notify) UpdateNotification(ctx *gin.Context, query *qo.UpdateNotificationQO) {
	//转换请求参数
	util.JsonConvert(ctx, query)

	//判断通知类型，若类型为4则uid默认为0;其余情况需将UID转换进行校验
	if query.Type == cpgConst.FOUR && len(query.Uid) > 1 && query.Uid[0] != cpgConst.ZERO {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	notifyDao := new(dao.NotifyDao)

	//查询id是否存在记录
	result := globalInit.Db.Model(&model.Notify{}).Where("id", query.Id).Find(notifyDao)
	if result.RowsAffected == int64(cpgConst.ZERO) {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//将uid列表、content、state、notify_date写入model
	uidByte, _ := json.Marshal(query.Uid)
	_ = copier.Copy(notifyDao, query)
	notifyDao.Uid = string(uidByte)

	//更新
	err := notifyDao.Update(ctx)

	if err != nil {
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, "")
}

func (Notify) SystemNotify(ctx *gin.Context) {
	result := new(vo.SystemNotificationVO)
	//查询通知时间最近且为系统通知且处于开启状态的一条通知详情
	globalInit.Db.Model(&model.Notify{}).
		Where("type", cpgConst.FOUR).
		Where("state", cpgConst.ONE).
		Where("end_date > ?", time.Now()).
		Find(result.NotificationList)
	if len(result.NotificationList) == cpgConst.ZERO {
		err := common.OK
		err.Message = "当前时间段暂无通知"
		common.SendResponse(ctx, err, "")
		return
	}
	common.SendResponse(ctx, common.OK, result)
}
