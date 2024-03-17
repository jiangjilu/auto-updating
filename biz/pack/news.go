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

package pack

import (
	"github.com/jiangjilu/auto-updating/biz/hertz_gen/news_gorm"
	"github.com/jiangjilu/auto-updating/biz/model"
)

// Newses Convert model.News list to news_gorm.News list
func Newses(models []*model.News) []*news_gorm.News {
	users := make([]*news_gorm.News, 0, len(models))
	for _, m := range models {
		if u := News(m); u != nil {
			users = append(users, u)
		}
	}
	return users
}

// News Convert model.News to news_gorm.News
func News(model *model.News) *news_gorm.News {
	if model == nil {
		return nil
	}
	return &news_gorm.News{
		ID:      int64(model.ID),
		Title:   model.Title,
		State:   news_gorm.State(model.State),
		Cid:     model.Cid,
		Content: model.Content,
	}
}
