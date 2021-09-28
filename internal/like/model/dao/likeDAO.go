package dao

import (
	"cpg-blog/global/common"
	"cpg-blog/global/cpgConst"
	"cpg-blog/global/globalInit"
	"cpg-blog/internal/like/model"
	"gorm.io/gorm"
	"log"
	"reflect"
)

type LikeDAO struct{}

var Db = &(globalInit.Db)

func (d LikeDAO) Creat(uid int, objType int, objId int64) (err error) {
	tx := globalInit.Transaction()

	err = func(db *gorm.DB) error {
		var zan model.Zan
		tx.Where(&model.Zan{Uid: uint(uid), ObjId: objId}).First(&zan)
		if tx.Error != nil {
			return tx.Error
		}

		//是否存在点赞记录
		if reflect.DeepEqual(zan, model.Zan{}) {
			log.Println("dddd", zan)
			tx.Create(&model.Zan{
				Uid:     uint(uid),
				ObjType: objType,
				ObjId:   objId,
				State:   cpgConst.ZERO,
			})
		} else if zan.State == cpgConst.ONE {
			//存在记录则只更新状态
			zan.State = cpgConst.ZERO
			tx.Select("state").Updates(zan)
		} else if zan.State == cpgConst.ZERO {
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
