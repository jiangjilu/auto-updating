/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mysql

import "github.com/jiangjilu/auto-updating/biz/model"

func CreateNews(list []*model.News) error {
	return DB.Create(list).Error
}

func DeleteNews(id int64) error {
	return DB.Where("id = ?", id).Delete(&model.News{}).Error
}

func UpdateNews(news *model.News) error {
	return DB.Updates(news).Error
}

func QueryNews(keyword *string, page, pageSize int64) ([]*model.News, int64, error) {
	db := DB.Model(model.News{})
	if keyword != nil && len(*keyword) != 0 {
		db = db.Where(DB.Or("title like ?", "%"+*keyword+"%").
			Or("content like ?", "%"+*keyword+"%"))
	}
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var res []*model.News
	if err := db.Limit(int(pageSize)).Offset(int(pageSize * (page - 1))).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, total, nil
}
