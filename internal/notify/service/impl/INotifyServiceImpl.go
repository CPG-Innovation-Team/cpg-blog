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
	"log"
	"sort"
	"strconv"
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
	if query.Type == cpgConst.FOUR && len(query.Uid) > cpgConst.ONE {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	} else if query.Type == cpgConst.FOUR {
		query.Uid = []int{cpgConst.ZERO}
	}

	//校验时间
	nowTime := time.Now()
	beginTimeInt, _ := strconv.ParseInt(query.BeginTime, 10, 64)
	endTimeInt, _ := strconv.ParseInt(query.EndTime, 10, 64)
	begin := time.Unix(beginTimeInt, 0)
	end := time.Unix(endTimeInt, 0)


	//校验参数
	if end.Before(begin) || end.Before(nowTime) {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//校验参数时间段内是否存在通知
	notifyList, err := dao.Notify{}.GetNotifyList(begin,end)

	if err != nil {
		log.Println(err)
	}

	if len(notifyList) > cpgConst.ZERO {
		e := common.ErrParam
		e.Message = "该时间段内已存在通知!"
		common.SendResponse(ctx, e, "")
		return
	}

	//将uid列表、content、state、beginTime、endTime写入model
	notifyDao := new(dao.Notify)
	uidByte, _ := json.Marshal(query.Uid)
	content, _ := json.Marshal(query.Content)

	_ = copier.Copy(notifyDao, query)

	notifyDao.Uid = string(uidByte)
	notifyDao.Content = string(content)
	notifyDao.BeginTime = begin
	notifyDao.EndTime = end

	//model写入数据库
	id, err := notifyDao.Creat()

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
	if query.Type == cpgConst.FOUR && len(query.Uid) > 1 {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	notifyDao := new(dao.Notify)

	//查询id是否存在记录
	result := globalInit.Db.Model(&model.Notify{}).Where("id", query.Id).Find(notifyDao)
	if result.RowsAffected == int64(cpgConst.ZERO) {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//校验时间
	nowTime := time.Now()
	beginTimeInt, _ := strconv.ParseInt(query.BeginTime, 10, 64)
	endTimeInt, _ := strconv.ParseInt(query.EndTime, 10, 64)
	begin := time.Unix(beginTimeInt, 0)
	end := time.Unix(endTimeInt, 0)
	if end.Before(begin) || end.Before(nowTime) {
		common.SendResponse(ctx, common.ErrParam, "")
		return
	}

	//校验参数时间段内是否存在通知
	notifyList, err := dao.Notify{}.GetNotifyList(begin,end)

	if err != nil {
		log.Println(err)
	}

	if len(notifyList) > cpgConst.ZERO {
		e := common.ErrParam
		e.Message = "该时间段内已存在通知!"
		common.SendResponse(ctx, e, "")
		return
	}

	//将uid列表、content、state、beginTime、endTime写入model
	uidByte, _ := json.Marshal(query.Uid)
	content, _ := json.Marshal(query.Content)

	_ = copier.Copy(notifyDao, query)

	notifyDao.Uid = string(uidByte)
	notifyDao.Content = string(content)
	notifyDao.BeginTime = begin
	notifyDao.EndTime = end

	//更新
	err = notifyDao.Update(ctx)

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
		Where("end_time > ?", time.Now()).
		Find(&result.NotificationList)
	if len(result.NotificationList) == cpgConst.ZERO {
		err := common.OK
		err.Message = "当前时间段暂无通知"
		common.SendResponse(ctx, err, "")
		return
	}

	//如果存在多条，则返回离当前时间最近的一条通知
	if len(result.NotificationList) > cpgConst.ONE {
		sort.Sort(result)
		log.Println("所有符合条件的通知：", result.NotificationList)
		common.SendResponse(ctx, common.OK, vo.SystemNotificationVO{NotificationList: result.NotificationList[:cpgConst.ONE]})
		return
	}

	common.SendResponse(ctx, common.OK, result)
	return
}
