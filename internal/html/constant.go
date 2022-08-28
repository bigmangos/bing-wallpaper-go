package html

import "strings"

// 侧边栏目录的归档菜单
const (
	Sidebar         = "${sidebar}"
	SidebarNowColor = "w3-green"
	SidebarColor    = "w3-hover-green"
	sidebarMenu     = "<a href=\"${sidebar_href_url}\" onclick=\"w3_close()\" class=\"w3-bar-item w3-button w3-hover-green w3-large\">${sidebar_href_title}</a>"
)

func getSidebarMenuList(hrefUrl, hrefTitle string) string {
	result := strings.Replace(sidebarMenu, "${sidebar_href_url}", hrefUrl, -1)
	return strings.Replace(result, "${sidebar_href_title}", hrefTitle, -1)

}

// 头部图片
const (
	HeadImgUrl  = "${head_img_url}"
	HeadImgDesc = "${head_img_desc}"
	HeadTitle   = "${head_title}"
)

// 图片列表
const (
	ImgCardList = "${img_card_list}"
	ImgCardUrl  = "${img_card_url}"
	ImgCardDate = "${img_card_date}"
	ImaCard     = "<div class=\"w3-third \" style=\"position: relative;\">\n" +
		"  <img class=\"smallImg\" src=\"${img_card_url}&pid=hp&w=50\"  style=\"width:95%;\" />" +
		" <img class=\"bigImg\" src=\"${img_card_url}&pid=hp&w=384&h=216&rs=1&c=4\" style=\"width:95%\" onload=\"imgloading(this)\">\n" +
		" <p>${img_card_date} <a href=\"${img_card_url}\" target=\"_blank\">Download 4k</a> </p>\n" +
		"</div>"
)

func getImgCard(imgUrl, date string) string {
	result := strings.Replace(ImaCard, ImgCardUrl, imgUrl, -1)
	return strings.Replace(result, ImgCardDate, date, -1)
}

// 底部归档
const (
	MonthHistory              = "${month_history}"
	MonthHistoryNowMonthColor = "w3-green"
	MonthHistoryMonthColor    = "w3-light-grey"
	MonthHistoryHrefUrl       = "${month_href_url}"
	MonthHistoryHrefTitle     = "${month_href_title}"
	MonthHistoryHref          = "<a class=\"w3-tag w3-button w3-hover-green w3-light-grey w3-margin-bottom\" href=\"${month_href_url}\">${month_href_title}</a>"
)

func getMonthHistory(url, title string) string {
	result := strings.Replace(MonthHistoryHref, MonthHistoryHrefUrl, url, -1)
	return strings.Replace(result, MonthHistoryHrefTitle, title, -1)
}
