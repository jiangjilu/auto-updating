/*
 * 实现一个简单的CMS系统
 * 已完成功能：用户开放式访问首页、列表页、详情页
 * 作者：Bill Jiang
 * 创建时间：2024-04-01
 */

package mw

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/jiangjilu/auto-updating/biz/dal/mysql"
	"github.com/jiangjilu/auto-updating/biz/model"
	"github.com/jiangjilu/auto-updating/biz/myutils"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 自定义模板函数
func wordsToUpper(words string) string {
	return strings.ToUpper(words)
}
func wordsToLower(words string) string {
	return strings.ToLower(words)
}
func keepHtml(s string) template.HTML {
	return template.HTML(s)
}
func cutString(s string, n int) string {
	s = stripTags(s)
	if len(s) > n {
		s = string([]rune(s)[:n])
	}
	return s + "..."
}

func stripTags(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	//去除html注释 <!-- Card Body -->
	re, _ = regexp.Compile("\\<\\!\\-\\-[\\S\\s]+?\\-\\-\\>")
	src = re.ReplaceAllString(src, "")
	return strings.TrimSpace(src)
}

func time2timestamp(time time.Time) int64 {
	return time.Unix()
}

func GetTerms() interface{} {
	var rowset []model.News

	return rowset
}

func Cms(h *server.Hertz) app.HandlerFunc {
	h.Delims("{{", "}}")
	h.StaticFS("/templates", &app.FS{})
	h.SetFuncMap(template.FuncMap{
		"wordsToUpper":   wordsToUpper,
		"wordsToLower":   wordsToLower,
		"keepHtml":       keepHtml,
		"stripTags":      stripTags,
		"cutString":      cutString,
		"time2timestamp": time2timestamp,
		"getConfig":      myutils.GetConfig,
		"getTerms":       GetTerms,
	})
	h.LoadHTMLGlob("./templates/*.html")

	RegisterGroupRoute(h)

	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		fmt.Println("cms pre-handle")

		c.Next(ctx) // call the next middleware(handler)

		// post-handle
		fmt.Println("cms post-handle")
	}
}

func RegisterGroupRoute(h *server.Hertz) {
	// 公开访问
	publicGroup := h.Group("")
	{
		// 首页
		publicGroup.GET("/", func(ctx context.Context, c *app.RequestContext) {
			c.HTML(http.StatusOK, "index.html", utils.H{
				"title": "首页",
			})
		})

		// 列表
		publicGroup.GET("/list/*term", func(ctx context.Context, c *app.RequestContext) {
			var rowset []model.News
			mysql.DB.Scopes(Paginate(c)).Order("id desc").Find(&rowset)
			c.HTML(http.StatusOK, "list.html", utils.H{
				"title":  "列表",
				"url":    c.FullPath(),
				"rowset": rowset,
			})
		})

		// 详情
		publicGroup.GET("/detail/:id", func(ctx context.Context, c *app.RequestContext) {
			id := c.Param("id")
			sql := fmt.Sprintf("SELECT * FROM news WHERE ID IN (%s)", id)
			var row model.News
			mysql.DB.Raw(sql).Scan(&row)
			c.HTML(http.StatusOK, "detail.html", utils.H{
				"id":    id,
				"title": "详情",
				"url":   c.FullPath(),
				"row":   row,
			})
		})
	}
}

func Paginate(r *app.RequestContext) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r
		page, _ := strconv.Atoi(q.Query("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
