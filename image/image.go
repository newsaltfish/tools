package img

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"os"
)

// ImgInterface 图形处理接口
type ImgInterface interface {
	image.Image
	imageRotate(img image.Image, angle float64) image.Image //图片旋转
	imaOverturn(img image.Image, way int) image.Image       //图片翻转
}

// ImgStruct 图片类型
type ImgStruct struct {
	ImgInterface
}

//func NewImg(img image.Image) ImgStruct{
//return
//}

const (
	dx = 1500
	dy = 1300
)

func createImage() {
	file, err := os.Create("test.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//	var ss = TestInterface
	alpha := image.NewAlpha(image.Rect(0, 0, dx, dy))
	ngba := image.NewRGBA(image.Rect(0, 0, dx, dy))
	//	var array []image.Point

	for j := 0; j != 10; j++ {
		var r = 500.0 + float64(j)
		for i := 0.0; i <= math.Pi*2; i += math.Pi / (4 * r) {
			var x = int(r*math.Cos(i)) + dx/2
			var y = int(r*math.Sin(i)) + dy/2
			ngba.Set(x, y, color.Alpha{uint8(255)})
			ngba.SetRGBA(x, 150, color.RGBA{uint8(253), uint8(0), 253, 255})
			ngba.SetRGBA(250, y, color.RGBA{uint8(253), uint8(0), 253, 255})
			//			array = append(array, image.Point{x, y})
			ngba.SetRGBA(x, y, color.RGBA{uint8(253), uint8(0), 253, 255})
		}
	}
	fmt.Println(alpha.At(400, 100))                      //144 在指定位置的像素
	fmt.Println(alpha.Bounds())                          //(0,0)-(500,200)，图片边界
	fmt.Println(alpha.Opaque())                          //false，是否图片完全透明
	fmt.Println(alpha.PixOffset(1, 1))                   //501，指定点相对于第一个点的距离
	fmt.Println(alpha.Stride)                            //500，两个垂直像素之间的距离
	jpeg.Encode(file, ngba, &jpeg.Options{Quality: 100}) //将image信息写入文件中
}

func imageChange() {
	// 创建新图片
	file, err := os.Create("change.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// 资源图片
	srcfile, _ := os.Open("QQ.png")
	pngImg, _ := png.Decode(srcfile)
	defer srcfile.Close()
	//	// 原始图片
	//	before, _ := os.Open("test.jpeg")
	//	beforeImg, _ := jpeg.Decode(before)
	//	defer before.Close()
	jpg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	//	draw.Src
	// 蒙版图片
	draw.Draw(jpg, jpg.Bounds().Add(image.Point{10, 10}), pngImg, pngImg.Bounds().Min, draw.Src) //原始图片转换为灰色图片
	////	draw.Draw()
	jpeg.Encode(file, jpg, &jpeg.Options{100})
}

// imageRotate 图片旋转
func imageRotate(mm image.Image) {
	//	var img = new(ImgStruct)

}
