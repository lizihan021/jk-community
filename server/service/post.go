package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// @title    CreatePost
// @description   create a post, 新增post
// @auth                     （2020/04/05  20:22）
// @param     api             model.Post
// @return                    error

func CreatePost(api model.SysApi) (err error) {
	findOne := global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).Find(&model.SysApi{}).Error
	if findOne == nil {
		return errors.New("存在相同api")
	} else {
		err = global.GVA_DB.Create(&api).Error
	}
	return err
}
