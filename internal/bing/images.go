package bing

type Image struct {
	desc string
	date string
	url  string
}

func (im *Image) Desc() string {
	if im != nil {
		return im.desc
	}
	return ""
}

func (im *Image) SetDesc(desc string) {
	if im != nil {
		im.desc = desc
	}
}

func (im *Image) Date() string {
	if im != nil {
		return im.date
	}
	return ""
}

func (im *Image) SetDate(date string) {
	if im != nil {
		im.date = date
	}
}

func (im *Image) Url() string {
	if im != nil {
		return im.url
	}
	return ""
}

func (im *Image) SetUrl(url string) {
	if im != nil {
		im.url = url
	}
}

func NewImage(desc, date, url string) *Image {
	return &Image{desc: desc, date: date, url: url}
}

func (im *Image) String() string {
	smallUrl := im.url + "&pid=hp&w=384&h=216&rs=1&c=4"
	return "![](" + smallUrl + ")" + im.date + " [download 4K](" + im.url + ")"
}

func (im *Image) formatMarkdown() string {
	return im.date + " | [" + im.desc + "](" + im.url + ") "
}

func (im *Image) toLarge() string {
	smallUrl := im.url + "&w=1000"
	return "![](" + smallUrl + ")Today: [" + im.desc + "](" + im.url + ")"
}
