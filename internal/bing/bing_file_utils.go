package bing

import (
	"bytes"
	"os"
	"path"
	"strings"
)

// ReadBingWallpaperMd 读取 bing-wallpaper.md 文件
func ReadBingWallpaperMd(bingPath string) ([]*Image, error) {
	content, err := os.ReadFile(bingPath)
	if err != nil {
		return nil, err
	}
	allLines := strings.Split(string(content), "\n")
	imgList := make([]*Image, 0, len(allLines))
	for _, line := range allLines {
		if !strings.Contains(line, "|") {
			continue
		}

		descEnd := strings.Index(line, "]")
		urlsStart := strings.LastIndex(line, "(") + 1
		imgList = append(imgList, NewImage(line[14:descEnd], line[0:10], line[urlsStart:len(line)-2]))
	}
	return imgList, nil
}

// WriteBingWallpaperMd 写入 bing-wallpaper.md 文件
func WriteBingWallpaperMd(bingPath string, imgList []*Image) error {
	var buf bytes.Buffer
	buf.WriteString("## Bing Wallpaper\n")
	buf.WriteString("\n")
	for _, img := range imgList {
		buf.WriteString(img.formatMarkdown())
		buf.WriteString("\n\n")
	}

	if err := os.WriteFile(bingPath, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

// ReadReadme 读取 readme.md 文件
func ReadReadme(readmePath string) ([]*Image, error) {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return nil, err
	}
	allLines := strings.Split(string(content), "\n")
	imgList := make([]*Image, 0, len(allLines))
	for _, line := range allLines {
		if !strings.Contains(line, "|") {
			continue
		}
		s := strings.TrimSpace(line)
		descEnd := strings.Index(s, "]")
		urlsStart := strings.LastIndex(s, "(") + 1
		imgList = append(imgList, NewImage(line[0:10], line[14:descEnd], line[urlsStart:len(line)-1]))
	}
	return imgList, nil
}

// WriteReadme 写入 readme.md 文件
func WriteReadme(readmePath string, imgList []*Image) error {
	var buf bytes.Buffer
	images := imgList[0:30]
	writePicToBuf(&buf, "", images)

	buf.WriteString("\n")
	buf.WriteString("### 历史归档：")
	buf.WriteString("\n")

	// 月份归档链接
	dateList := make([]string, 0, 60)
	dateSet := make(map[string]struct{}, 60)
	for _, img := range imgList {
		if _, ok := dateSet[img.Date()[0:7]]; !ok {
			dateList = append(dateList, img.Date()[0:7])
			dateSet[img.Date()[0:7]] = struct{}{}
		}
	}

	var i int
	for _, date := range dateList {
		link := "[" + date + "](https://github.com/bigmangos/bing-wallpaper-go/tree/master/picture/" + date + "/) | "
		buf.WriteString(link)
		i++
		if i%8 == 0 {
			buf.WriteString("\n")
		}
	}

	if err := os.WriteFile(readmePath, buf.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

// WriteMonthInfo 按月份写入图片信息
func WriteMonthInfo(monthPicPath string, imgList []*Image) error {
	monthMap := ConvertImgListToMonthMap(imgList)
	for mouth, images := range monthMap {
		mouthPath := path.Join(monthPicPath, mouth)
		_, err := os.Stat(mouthPath)
		if os.IsNotExist(err) {
			if err = os.Mkdir(mouthPath, 0755); err != nil {
				return err
			}
		}
		var buf bytes.Buffer
		writePicToBuf(&buf, mouth, images)
		if err = os.WriteFile(path.Join(mouthPath, "README.md"), buf.Bytes(), 0644); err != nil {
			return err
		}
	}
	return nil
}

func ConvertImgListToMonthMap(imgList []*Image) map[string][]*Image {
	monthMap := make(map[string][]*Image)
	for _, img := range imgList {
		if img == nil || img.url == "" {
			continue
		}
		key := img.date[0:7]
		if _, ok := monthMap[key]; !ok {
			monthMap[key] = []*Image{img}
		} else {
			monthMap[key] = append(monthMap[key], img)
		}
	}
	return monthMap
}

// writePicToBuf 写入图片列表到指定位置
func writePicToBuf(buf *bytes.Buffer, mouth string, imgList []*Image) {
	title := "## Bing Wallpaper\n"
	if mouth != "" {
		title = "## Bing Wallpaper (" + mouth + ")\n"
	}
	buf.WriteString(title)
	buf.WriteString("本项目是 [niumoo/bing-wallpaper](https://github.com/niumoo/bing-wallpaper) 项目的Go实现，感谢原作者提供优质项目。\n\n")

	buf.WriteString(imgList[0].toLarge() + "\n")
	buf.WriteString("|      |      |      |\n")
	buf.WriteString("| :----: | :----: | :----: |\n")

	i := 1
	for _, img := range imgList {
		buf.WriteString("|" + img.String())
		if i%3 == 0 {
			buf.WriteString("|\n")
		}
		i++

	}
	if i%3 != 1 {
		buf.WriteString("|")
	}
}
