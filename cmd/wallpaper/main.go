package main

import (
	"encoding/json"
	"fmt"
	"github/bigmangos/bing-wallpaper/internal/bing"
	"github/bigmangos/bing-wallpaper/internal/html"
	"github/bigmangos/bing-wallpaper/internal/model"
	"strings"
	"time"
)

const (
	bingApi = "https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=10&nc=1612409408851&pid=hp&FORM=BEHPTB&uhd=1&uhdwidth=3840&uhdheight=2160"
	bingUrl = "https://cn.bing.com"
)

func main() {
	body := bing.Get(bingApi)
	var allImg model.Images
	err := json.Unmarshal(body, &allImg)
	if err != nil {
		fmt.Println("Unmarshal err: ", err)
		return
	}
	if len(allImg.Images) == 0 {
		fmt.Println("no images")
		return
	}
	image := allImg.Images[0]

	picUrl := bingUrl + image.Url[0:strings.Index(image.Url, "&")]
	localDate, _ := time.ParseInLocation("20060102", image.EndDate, time.UTC)
	desc := image.Copyright

	images, err := bing.ReadBingWallpaperMd(html.BingPath)
	if err != nil {
		fmt.Println("ReadBing", err)
		return
	}

	if images[0].Desc() == desc {
		return
	}

	newImages := make([]*bing.Image, len(images)+1, len(images)+1)
	newImages[0] = bing.NewImage(desc, localDate.Format("2006-01-02"), picUrl)
	copy(newImages[1:], images)

	if err = bing.WriteBingWallpaperMd(html.BingPath, newImages); err != nil {
		fmt.Println("WriteBing", err)
		return
	}
	if err = bing.WriteReadme(html.ReadmePath, newImages); err != nil {
		fmt.Println("WriteReadme", err)
		return
	}
	if err = bing.WriteMonthInfo(html.MonthPath, newImages); err != nil {
		fmt.Println("WriteMonthInfo", err)
		return
	}
	html.GenerateAll(newImages)
}
