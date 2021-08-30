package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fogleman/gg"
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var (
	bgImg     image.Image
	qrCodeImg image.Image
	offset    image.Point
)

type QrCode struct {
	Debug        bool   `json:"debug"`         // 调试模式本地生成海报 不走OSS
	Id           string `json:"id"`            // 海报的唯一标识
	Size         int    `json:"size"`          // 二维码大小 150 == x150 y150
	Type         int    `json:"type"`          // 生成图片类型，默认1.jpg  2.png
	Content      string `json:"content"`       // 二维码识别出的内容
	BackendImage string `json:"backend_image"` // 背景图片名称 3.png
	MidX         bool   `json:"mid_x"`         // 二维码X坐标是否居中
	MidY         bool   `json:"mid_y"`         // 二维码y坐标是否居中
	X            int    `json:"x"`             // 二维码相对图片坐标
	Y            int    `json:"y"`
}
// content: 二维码识别信息内容输入
func (q *QrCode) createQrCode() (img image.Image, err error) {
	qrCode, err := qrcode.New(q.Content, qrcode.Highest)
	if err != nil {
		err = errors.New("创建二维码失败")
		return
	}

	qrCode.DisableBorder = true
	img = qrCode.Image(q.Size)
	return
}

func readImgData(url string) (pix []byte, file io.ReadCloser, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	// defer resp.Body.Close()
	file = resp.Body
	return
}

func (q *QrCode) QrCode4ImageDebug() (err error) {
	nameList := strings.Split(q.BackendImage, ".")
	imageType := nameList[len(nameList)-1]
	qrCodeImg, err = q.createQrCode()
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}

	i, err := os.Open(path.Base("./" + q.BackendImage))
	if err != nil {
		return
	}
	defer i.Close()
	switch imageType {
	case "png":
		bgImg, err = png.Decode(i)
		if err != nil {
			return
		}
	case "jpg", "jpeg":
		bgImg, err = jpeg.Decode(i)
		if err != nil {
			return
		}
	default:
		err = errors.New("图片格式只支持png/jpg/jpeg")
		return
	}

	b := bgImg.Bounds()
	offset = image.Pt(q.X, q.Y)
	if q.MidX {
		offset = image.Pt(b.Max.X/2-q.Size/2, q.Y)
	}

	if q.MidY {
		offset = image.Pt(q.X, b.Max.Y/2-q.Size/2)
	}

	if q.MidX && q.MidY {
		offset = image.Pt(b.Max.X/2-q.Size/2, b.Max.Y/2-q.Size/2)
	}
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, qrCodeImg.Bounds().Add(offset), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Over)

	// 本地生成海报图
	nowName := fmt.Sprintf("%s_backend_%s.%s", nameList[0], q.Id, imageType)
	j, err := os.Create(path.Base(nowName))
	if err != nil {
		return
	}
	defer j.Close()
	if nameList[1] == "png" {
		_ = png.Encode(j, m)
	} else {
		_ = jpeg.Encode(j, m, nil)
	}
	return
}

func (q *QrCode) QrCode4Image() (addr string, err error) {
	if q.Debug {
		err = q.QrCode4ImageDebug()
		if err != nil {
			return
		}
		return
	}

	nameList := strings.Split(q.BackendImage, ".")
	imageType := nameList[len(nameList)-1]
	imageHostList := strings.Split(nameList[len(nameList)-2], "/")
	imageHost := imageHostList[len(imageHostList)-1]
	nowName := fmt.Sprintf("%s_backend_%s.%s", imageHost, q.Id, imageType)
	//if !oss.IsFilePostfix(nowName) { // 云OSS格式限制判断
	//	err = errors.New("上传文件格式不符合规范，请重新上传~")
	//	return
	//}

	//ossClient := oss.GetClient()             // 初始化你的oss
	//exit, _ := ossClient.FileIsExit(nowName) // 判断OSS上该文件是否已存在
	addr = nowName                           // 返回文件名

	_, file, err := readImgData(q.BackendImage) // 读取背景图 URL
	if err != nil {
		return
	}
	defer file.Close()

	qrCodeImg, err = q.createQrCode()
	if err != nil {
		fmt.Println("生成二维码失败:", err)
		return
	}
	switch imageType {
	case "png":
		bgImg, err = png.Decode(file)
		if err != nil {
			return
		}
	case "jpg", "jpeg":
		bgImg, err = jpeg.Decode(file)
		if err != nil {
			return
		}
	default:
		err = errors.New("图片格式只支持png/jpg/jpeg")
		return
	}

	b := bgImg.Bounds()
	offset = image.Pt(q.X, q.Y)
	if q.MidX {
		offset = image.Pt(b.Max.X/2-q.Size/2, q.Y)
	}

	if q.MidY {
		offset = image.Pt(q.X, b.Max.Y/2-q.Size/2)
	}

	if q.MidX && q.MidY {
		offset = image.Pt(b.Max.X/2-q.Size/2, b.Max.Y/2-q.Size/2)
	}
	m := image.NewRGBA(b)
	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0}, draw.Src)
	draw.Draw(m, qrCodeImg.Bounds().Add(offset), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Over)

	// 上传至 oss
	imgBuff := bytes.NewBuffer(nil)
	if nameList[1] == "png" {
		_ = png.Encode(imgBuff, m)
	} else {
		_ = jpeg.Encode(imgBuff, m, nil)
	}

	//ossName, err := ossClient.Upload(nowName, imgBuff.Bytes())
	//log.Println(ossName)
	return
}

func main(){
	//background := image.NewRGBA(image.Rect(0,0,500,500))

	url:= "https://v2data.cdn.pulizu.com/20210829/20210829085412665825.jpg?x-oss-process=style/watermark_style"
	r,_:=http.Get(url)
	im,s,_:=image.Decode(r.Body)
	fmt.Println(s)
	//
	head:="https://cdn2.jianshu.io/assets/default_avatar/7-0993d41a595d6ab6ef17b19496eb2f21.jpg"
	r2,_:=http.Get(head)
	im2,s2,_:=image.Decode(r2.Body)
	fmt.Println(s2)
	f,_:=os.Create("1.jpg")
	defer f.Close()

	alpha := image.NewAlpha(image.Rect(0, 0, im.Bounds().Dx(), im.Bounds().Dy()))
	for x := 0; x < im.Bounds().Dx(); x++ {
		for y := 0; y < im.Bounds().Dy(); y++ {
			alpha.Set(x, y, im.At(x,y))   //设定alpha图片的透明度
		}
	}

	//for x := 0; x < im2.Bounds().Dx(); x++ {
	//	for y := 0; y < im2.Bounds().Dy(); y++ {
	//		alpha.Set(x, y, im.At(x,y))   //设定alpha图片的透明度
	//	}
	//}

	fmt.Println(im2.Bounds())
	jpeg.Encode(f, alpha, nil)      //将image信息写入文件中

	fmt.Println("OK")

}

// DrawCircleImg 头像圆图
func DrawCircleImg(im image.Image) (image.Image, error) {
	loadStart := time.Now()
	defer func() {
		log.Println("生成圆头像耗时: %v\n", time.Now().Sub(loadStart).Milliseconds())
	}()
	b := im.Bounds()
	w := float64(b.Dx())
	h := float64(b.Dy())
	dc := gg.NewContext(int(w), int(h))

	r := float64(w / 2) // 半径

	dc.DrawRoundedRectangle(0, 0, w, h, r)
	dc.Clip()
	dc.DrawImage(im, 0, 0)
	return dc.Image(), nil
}


// ScaleImage 缩放图片
func ScaleImage(image image.Image, x, y int) image.Image {
	loadStart := time.Now()
	defer func() {
		log.Println("缩放耗时: %v\n", time.Now().Sub(loadStart).Milliseconds())
	}()
	w := image.Bounds().Size().X
	h := image.Bounds().Size().Y
	dc := gg.NewContext(x, y)
	var ax float64 = float64(x) / float64(w)
	var ay float64 = float64(y) / float64(h)
	dc.Scale(ax, ay)
	dc.DrawImage(image, 0, 0)
	return dc.Image()
}