package html

import (
	"fmt"
	"github/bigmangos/bing-wallpaper/internal/bing"
	"strings"
)

const (
	BingPath   = "bing-wallpaper.md"
	ReadmePath = "README.md"
	MonthPath  = "picture/"
)

func GenerateAll(images []*bing.Image) {
	monthMap := bing.ConvertImgListToMonthMap(images)
	generateIndex(images, monthMap)
	generateMonth(monthMap)
}

func generateIndex(images []*bing.Image, monthMap map[string][]*bing.Image) {
	templateFile, err := readIndexTemplateFile()
	if err != nil {
		fmt.Println("ReadIndexTemplateFile", err)
		return
	}
	// 替换头部图片和描述
	html := replaceHead(templateFile, "", images[0])
	// 替换侧边目录
	html = replaceSidebar(html, "", monthMap)
	// 替换图片列表
	html = replaceImgList(html, images[0:30])
	// 替换底部月度历史
	html = replaceMonthHistory(html, "", monthMap)
	// 写入文件
	if err = writeIndexHtml(html); err != nil {
		fmt.Println("WriteIndexHtml", err)
	}

}

func generateMonth(monthMap map[string][]*bing.Image) {
	for month, images := range monthMap {
		templateFile, err := readIndexTemplateFile()
		if err != nil {
			fmt.Println("ReadIndexTemplateFile", err)
			continue
		}
		// 替换头部图片和描述
		html := replaceHead(templateFile, month, images[0])
		// 替换侧边目录
		html = replaceSidebar(html, month, monthMap)
		// 替换图片列表
		html = replaceImgList(html, images)
		// 替换底部月度历史
		html = replaceMonthHistory(html, month, monthMap)
		// 写入文件
		if err = writeMonthHtml(month, html); err != nil {
			fmt.Println("WriteMonthHtml", err)
			continue
		}
	}
}

func replaceSidebar(html, nowMonth string, monthMap map[string][]*bing.Image) string {
	var buf strings.Builder
	for month, _ := range monthMap {
		sidebarMenu := getSidebarMenuList(month+".html", month)
		if month != "" && month == nowMonth {
			sidebarMenu = strings.Replace(sidebarMenu, SidebarColor, SidebarNowColor, -1)
		}
		buf.WriteString(sidebarMenu)
	}
	return strings.Replace(html, Sidebar, buf.String(), -1)
}

// replaceHead 更新头部大图和描述
func replaceHead(html, month string, images *bing.Image) string {
	html = strings.Replace(html, HeadImgUrl, images.Url(), -1)
	html = strings.Replace(html, HeadImgDesc, images.Desc(), -1)
	if month != "" {
		return strings.Replace(html, HeadTitle, "Bing Wallpaper("+month+")", -1)
	}
	return strings.Replace(html, HeadTitle, "Bing Wallpaper", -1)
}

func replaceImgList(html string, bingImages []*bing.Image) string {
	var imgList strings.Builder
	for _, image := range bingImages {
		imgList.WriteString(getImgCard(image.Url(), image.Date()))
	}
	return strings.Replace(html, ImgCardList, imgList.String(), -1)
}

// replaceMonthHistory 替换底部的月度历史链接
func replaceMonthHistory(html, nowMonth string, monthMap map[string][]*bing.Image) string {
	var buf strings.Builder
	for month, _ := range monthMap {
		history := getMonthHistory(month+".html", month)
		if month != "" && month == nowMonth {
			history = strings.Replace(history, MonthHistoryMonthColor, MonthHistoryNowMonthColor, -1)
		}
		buf.WriteString(history + " ")
	}
	return strings.Replace(html, MonthHistory, buf.String(), -1)
}
