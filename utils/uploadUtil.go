package utils

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
)

type UploadConf struct {
	UpdateVilidateExtString string           `json:"updateVilidateExtString"` //从配置里获取,空为不限制  示例 "jpg,jpeg,png,gif,bmp"
	FileMaxSize             int64            `json:"fileMaxSize"`             //单位Mb 0:不限制大小 / 5
	Drive                   string           `json:"drive"`
	UploadConfLocal         *UploadConfLocal `json:"uploadConfLocal"`
	UploadConfOss           *UploadConfOss   `json:"uploadConfOss"`
	UploadConfQiniu         *UploadConfQiniu `json:"uploadConfQiniu"`
}

// 本地配置
type UploadConfLocal struct {
	Host string `json:"updateVilidateExtString"` //根域名
}

// oss配置
// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
type UploadConfOss struct {
	Host            string `json:"host"`     //根域名
	Endpoint        string `json:"endpoint"`
	HostRepostry    string `json:"hostRepostry"` //文件显示使用前缀
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BucketName      string `json:"bucketName"`
}

type UploadConfQiniu struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	ImgPath   string `json:"imgPath"`
}

func InitUploadConf() *UploadConf {
	return &UploadConf{
		UpdateVilidateExtString: "",      //从配置里获取,空为不限制  示例 "jpg,jpeg,png,gif,bmp"
		FileMaxSize:             5,       //单位Mb 0:不限制大小 / 5
		Drive:                   "oss", // local / oss / qiniu
		UploadConfLocal: &UploadConfLocal{
			Host: "http://localhost:9000", //根域名
		},
		//以下oss账号仅限测试。非法操作者。请自重 感谢 壹cup清茶 贡献的测试账号
		UploadConfOss: &UploadConfOss{
			//公司的账号
			Host:            "https://oss-cn-beijing.aliyuncs.com", //返回的自己的阿里云域名
			HostRepostry:     "https://adcenter02.oss-cn-beijing.aliyuncs.com", //文件显示使用前缀
			AccessKeyId:     "LTAI5t5Yo4DSh3VWA3ZydMiV",
			AccessKeySecret: "89gCTM5juPyXZeuiKevCEsRahm2KUO",
			BucketName:      "adcenter02",
		},
		//请自行七牛申请对应的 公钥 私钥 bucket 和 域名地址
		//UploadConfQiniu: &UploadConfQiniu{
		//	AccessKey: "",
		//	SecretKey: "",
		//	Bucket:    "",
		//	ImgPath:   "",
		//},
	}
}

// 初使化配置
var UploadConfig = new(UploadConf)

func init() {
	UploadConfig = InitUploadConf()
}

//单个文件 上传文件到本地目录里
func UplaodFile(c *gin.Context) (err error, host, pathRelative string) {
	// 单个文件
	//fileName := c.MultipartForm("avatar")
	//fmt.Println(fileName)
	file, err := c.FormFile("file")
	if err != nil {
		err = errors.New(fmt.Sprintf("获取数据失败，%v", err))
		return
	}
	return ActionUplaodFile(c, file)
}

//多图片返回参
type ResFiles struct {
	HostListStr string     `json:"hostListStr"`
	UriListStr  string     `json:"uriListStr"`
	UrlListStr  string     `json:"urlListStr"`
	FileList    []FileItem `json:"fileList"`
}
type FileItem struct {
	Host string `json:"host"`
	Uri  string `json:"uri"`
	Url  string `json:"url"`
}

//多文件上传
func UpLoadMultipartFile(c *gin.Context) (resFiles ResFiles, err error) {
	form, _ := c.MultipartForm()
	files := form.File["files[]"]
	if len(files) < 1 {
		err = errors.New("文件不能为空")
		return
	}
	for _, file := range files {
		err, host, pathRelative := ActionUplaodFile(c, file)
		if err != nil {
			continue
		}
		var rItem FileItem
		rItem.Host = host
		rItem.Uri = pathRelative
		rItem.Url = path.Join(host, pathRelative)
		resFiles.FileList = append(resFiles.FileList, rItem)
		if resFiles.HostListStr == "" {
			resFiles.HostListStr = rItem.Host
			resFiles.UriListStr = rItem.Uri
			resFiles.UrlListStr = rItem.Url
		} else {
			resFiles.HostListStr += fmt.Sprintf(",%s", rItem.Host)
			resFiles.UriListStr += fmt.Sprintf(",%s", rItem.Uri)
			resFiles.UrlListStr += fmt.Sprintf(",%s", rItem.Url)
		}
	}
	return resFiles, nil
}

/**
@ 执行单个文件 上传文件到本地目录里
@ 调用方法 上传图片流,文件名为file   调用tools.UplaodFileToLocal(c)
@ return host 根域名
@ return path 上传后的相对文件路径
@ author haima
*/
func ActionUplaodFile(c *gin.Context, file *multipart.FileHeader) (err error, host, pathRelative string) {

	// 判断上传文件的大小
	if UploadConfig.FileMaxSize != 0 {
		fsize := file.Size //上传文件的大小
		if fsize > UploadConfig.FileMaxSize*1024*1024 {
			err = errors.New(fmt.Sprintf("只能上传小于 %dMb 的文件 ", UploadConfig.FileMaxSize))

			return
		}
	}

	//获取上传文件后缀
	extString := strings.ToUpper(Ext(file.Filename))
	if extString == "" {
		err = errors.New(fmt.Sprintf("上传失败，文件类型不支持，只能上传 %s 格式的。", UploadConfig.UpdateVilidateExtString))
		return
	}

	//验证文件类型
	if len(UploadConfig.UpdateVilidateExtString) > 0 {
		VilidateExtStrSliceTmp := strings.Split(strings.ToUpper(UploadConfig.UpdateVilidateExtString), ",")
		VilidateExtStrSlice := make([]string, 0)
		for _, v := range VilidateExtStrSliceTmp {
			VilidateExtStrSlice = append(VilidateExtStrSlice, fmt.Sprintf(".%s", v))
		}

		//验证文件类型
		//extString .JPG
		//VilidateExtStrSlice  []string{".JPG",".JPEG",".PNG",".GIF",".BMP"}
		if !ContainArray(extString, VilidateExtStrSlice) {
			err = errors.New(fmt.Sprintf("上传失败，文件类型不支持，只能上传 %s 格式的。", UploadConfig.UpdateVilidateExtString))
			return
		}
	}
	//上传的文件名
	file.Filename = fmt.Sprintf("%s%s", time.Now().Format("20060102150405"), file.Filename) // 文件名格式 自己可以改 建议保证唯一性 20060102150405test.xlsx
	switch UploadConfig.Drive {
	case "local":
		return uploadfileToLocal(c, file, file.Filename)
	case "oss":
		return uploadfileToOss(c, file, file.Filename)
	case "qiniu":
		return nil, UploadConfig.UploadConfQiniu.ImgPath, "qiniu 开发中...."
	default:
		err = errors.New("只支持上传到本地,oss,七牛")
		return
	}
}

func uploadfileToLocal(c *gin.Context, file *multipart.FileHeader, filename string) (err error, host, pathRelative string) {
	host = UploadConfig.UploadConfLocal.Host
	filepath := path.Join("static/uploadfile", time.Now().Format("20060102"))
	//上传到的路径
	pathRelative = path.Join(filepath, filename) //路径+文件名上传
	//如果没有filepath文件目录就创建一个
	if _, err := os.Stat(filepath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(filepath, os.ModePerm)
		}
	}
	// 上传文件到指定的目录
	err = c.SaveUploadedFile(file, pathRelative)
	if err != nil {
		err = errors.New(fmt.Sprintf("上传失败，%v", err))
		return
	}
	return
}

func uploadfileToOss(c *gin.Context, file *multipart.FileHeader, filename string) (err error, host, ossPathFileName string) {
	host = UploadConfig.UploadConfOss.Host
	err, _, localPathFileName := uploadfileToLocal(c, file, filename) //上传到本地
	if err != nil {
		err = errors.New(fmt.Sprintf("上传失败，%v", err))
	}

	ossPath := path.Join("upload", time.Now().Format("20060102"))
	ossPathFileName = path.Join(ossPath, file.Filename)
	// 创建OSSClient实例。
	client, err := oss.New(UploadConfig.UploadConfOss.Host, UploadConfig.UploadConfOss.AccessKeyId, UploadConfig.UploadConfOss.AccessKeySecret)
	if err != nil {
		os.Remove(localPathFileName)
		err = errors.New(fmt.Sprintf("文件上传服务器失败. err:%s", err.Error()))
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(UploadConfig.UploadConfOss.BucketName)
	if err != nil {
		os.Remove(localPathFileName)
		err = errors.New(fmt.Sprintf("文件上传云端失败. err:%s", err.Error()))
		return
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(ossPathFileName, localPathFileName)
	if err != nil {
		os.Remove(localPathFileName)
		err = errors.New(fmt.Sprintf("文件上传云端失败. err:%s", err.Error()))
		return
	}
	err = os.Remove(localPathFileName) //如果本地不想删除,可以注释了
	if err != nil {
		fmt.Println(err)
	}
	return
}

//获取文件的扩展名
func Ext(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

//Contain 判断obj是否在target中，target支持的类型array,slice,map   false:不在 true:在
func ContainArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}
