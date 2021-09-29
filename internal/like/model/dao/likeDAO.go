package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/like/model"
	"gorm.io/gorm"
	"reflect"
)

type LikeDAO struct{}

var Db = &(globalInit.Db)

func (d LikeDAO) SelectZan(uid int, objType int, objId int64) (zan model.Zan) {
	(*Db).Where(&model.Zan{Uid: uint(uid), ObjType: objType, ObjId: objId}).Find(&zan)
	return
}

func (d LikeDAO) CreatOrUpdate(uid int, objType int, objId int64, cancelLike bool) (err error) {
	tx := globalInit.Transaction()

	err = func(db *gorm.DB) error {
		var zan model.Zan
		tx.Where(&model.Zan{Uid: uint(uid), ObjType: objType, ObjId: objId}).First(&zan)
		if tx.Error != nil {
			return tx.Error
		}

		//是否存在点赞记录
		if reflect.DeepEqual(zan, model.Zan{}) {
			tx.Create(&model.Zan{
				Uid:     uint(uid),
				ObjType: objType,
				ObjId:   objId,
				State:   cpgConst.ZERO,
			})
		} else if zan.State == cpgConst.ONE && !cancelLike{//存在记录则只更新状态
			zan.State = cpgConst.ZERO
			tx.Select("state").Updates(zan)
		} else if zan.State == cpgConst.ZERO && cancelLike{
			zan.State = cpgConst.ONE
			tx.Select("state").Updates(zan)
		}else{
			return common.OK
		}

		if tx.Error != nil {
			tx.Rollback()
			return tx.Error
		}
		return tx.Commit().Error
	}(tx)
	return err
}
