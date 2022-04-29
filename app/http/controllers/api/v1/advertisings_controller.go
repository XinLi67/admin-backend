package v1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gohub/app/http/assemblies"
	"gohub/app/models/advertising"
	"gohub/app/models/advertising_plan"
	"gohub/app/policies"
	"gohub/app/requests"
	"gohub/pkg/paginator"
	"gohub/pkg/response"
	"gohub/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdvertisingsController struct {
	BaseAPIController
}

func (ctrl *AdvertisingsController) Index(c *gin.Context) {

	title := c.Query("title")
	status := c.Query("status")
	adtype := c.Query("type")
	advertising_position_id := c.Query("advertising_position_id")
	start_date := c.Query("start_date")
	end_date := c.Query("end_date")
	creator_name := c.Query("creator_name")

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	var data []advertising.Advertising
	var pager paginator.Paging

	if len(title) > 0 || len(status) > 0 || len(adtype) > 0 || len(advertising_position_id) > 0 || len(start_date) > 0 || len(end_date) > 0 || len(creator_name) > 0 {
		data, pager = advertising.Paginate2(c, 0)
	} else {
		data, pager = advertising.Paginate(c, 0)
	}

	advertisings := assemblies.AdvertisingAssemblyFromModelList(data, len(data))
	response.JSON(c, gin.H{
		"data":  advertisings,
		"pager": pager,
	})
}

func (ctrl *AdvertisingsController) Show(c *gin.Context) {
	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	advertisingAssembly := assemblies.AdvertisingAssemblyFromModel(advertisingModel)
	response.Data(c, advertisingAssembly)
}

func (ctrl *AdvertisingsController) Store(c *gin.Context) {

	host := utils.UploadConfig.UploadConfOss.HostRepostry

	request := requests.AdvertisingRequest{}
	if ok := requests.Validate(c, &request, requests.AdvertisingSave); !ok {
		return
	}

	var upErr, upErr2, upErr3 error
	var pathRelative, pathRelative2, pathRelative3 string

	advertisingModel := advertising.Advertising{
		AdvertisingPositionId: request.AdvertisingPositionId,
		CreatorId:             request.CreatorId,
		DepartmentId:          request.DepartmentId,
		Title:                 request.Title,
		Type:                  request.Type,
		RedirectTo:            request.RedirectTo,
		MaterialId:            request.MaterialId,
		MaterialType:          request.MaterialType,
		Size:                  request.Size,
		RedirectParams:        request.RedirectParams,
		Description:           request.Description,
		Status:                request.Status,
		PushContent:           request.PushContent,
		PushTitle:             request.PushTitle,
		AdvertisingCreativity: request.AdvertisingCreativity,
		Url:                   "",
		Url2:                  "",
		Url3:                  "",
	}

	//upErr, _, pathRelative = utils.ActionUplaodFile(c, request.Url)
	//upErr2, _, pathRelative2 = utils.ActionUplaodFile(c, request.Url2)
	//upErr3, _, pathRelative3 = utils.ActionUplaodFile(c, request.Url3)

	if len(request.Url.Filename) > 0 {
		upErr, _, pathRelative = utils.ActionUplaodFile(c, request.Url)
		advertisingModel.Url = host + "/" + pathRelative
	}

	if len(request.Url2.Filename) > 0 {
		upErr, _, pathRelative = utils.ActionUplaodFile(c, request.Url2)
		advertisingModel.Url2 = host + "/" + pathRelative2
	}

	if len(request.Url3.Filename) > 0 {
		upErr, _, pathRelative = utils.ActionUplaodFile(c, request.Url3)
		advertisingModel.Url3 = host + "/" + pathRelative3
	}

	if upErr != nil && upErr2 != nil && upErr3 != nil {
		return
	}

	advertisingModel.Create()
	if advertisingModel.ID > 0 {
		response.Created(c, advertisingModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingsController) Update(c *gin.Context) {

	advertisingModel := advertising.Get(c.Param("id"))
	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.AdvertisingRequest{}
	bindOk := requests.Validate(c, &request, requests.AdvertisingSave)
	if !bindOk {
		return
	}

	advertisingModel.AdvertisingNo=request.AdvertisingNo
	advertisingModel.AdvertisingPositionId = request.AdvertisingPositionId
	advertisingModel.CreatorId = request.CreatorId
	advertisingModel.DepartmentId = request.DepartmentId
	advertisingModel.Title = request.Title
	advertisingModel.Type = request.Type
	advertisingModel.RedirectTo = request.RedirectTo
	advertisingModel.MaterialId = request.MaterialId
	advertisingModel.MaterialType = request.MaterialType
	advertisingModel.Size = request.Size
	advertisingModel.RedirectParams = request.RedirectParams
	advertisingModel.Description = request.Description
	advertisingModel.Status = request.Status
	advertisingModel.PushContent = request.PushContent
	advertisingModel.PushTitle = request.PushTitle
	advertisingModel.AdvertisingCreativity = request.AdvertisingCreativity

	rowsAffected := advertisingModel.Save()
	if rowsAffected > 0 {
		response.Data(c, advertisingModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AdvertisingsController) Delete(c *gin.Context) {

	advertisingModel := advertising.Get(c.Param("id"))
	advertisingPlanList := make([]advertising_plan.AdvertisingPlan, 10)
	//advertisingPlanList[] := advertising_plan.Get(string(advertisingModel.ID))
	for _, plan := range advertisingPlanList {
		if plan.ID > 0 {
			plan.Delete()
		}
	}

	if advertisingModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	if advertisingModel.AdvertisingPositionId > 0 {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "广告已分配对应广告位，无法删除",
		})
		return
	}

	rowsAffected := advertisingModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *AdvertisingsController) BatchDelete(c *gin.Context) {

	request := requests.BatchDeleteRequest{}
	bindOk := requests.Validate(c, &request, requests.BatchDelete)
	if !bindOk {
		return
	}

	advertisingModel := advertising.Advertising{}
	if ok := policies.CanModifyAdvertising(c, advertisingModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := advertisingModel.BatchDelete(request.Ids)

	response.Data(c, map[string]int64{
		"rowsAffected": rowsAffected,
	})
}

//数据导出
func (ctrl *AdvertisingsController) Export(c *gin.Context) {

	listData := advertising.All2()
	f := excelize.NewFile() // 设置单元格的值
	// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "ID")
	f.SetCellValue("Sheet1", "B1", "广告编号")
	f.SetCellValue("Sheet1", "C1", "广告位编号")
	f.SetCellValue("Sheet1", "D1", "创建者编号")
	f.SetCellValue("Sheet1", "E1", "部门ID")
	f.SetCellValue("Sheet1", "F1", "标题")
	f.SetCellValue("Sheet1", "G1", "类型")
	f.SetCellValue("Sheet1", "H1", "跳转")
	f.SetCellValue("Sheet1", "I1", "素材ID")
	f.SetCellValue("Sheet1", "J1", "素材类型")
	f.SetCellValue("Sheet1", "K1", "尺寸")
	f.SetCellValue("Sheet1", "L1", "跳转参数")
	f.SetCellValue("Sheet1", "M1", "描述")
	f.SetCellValue("Sheet1", "N1", "状态")

	line := 1

	// 循环写入数据
	for _, v := range listData {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		//fmt.Println("AdvertisingNo:",v.AdvertisingNo)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.AdvertisingNo)
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", line), v.AdvertisingPositionId)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", line), v.CreatorId)
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", line), v.DepartmentId)
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", line), v.Title)
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", line), v.Type)
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", line), v.RedirectTo)
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", line), v.MaterialId)
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", line), v.MaterialType)
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", line), v.Size)
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", line), v.RedirectParams)
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", line), v.Description)
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", line), v.Status)
	}

	var fileName = utils.RandFileName()
	var fullPath = "G:/" + fileName + ".xlsx"

	// 保存文件
	if err := f.SaveAs(fullPath); err != nil {
		fmt.Println(err)
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;fileName=%s", fileName))
	c.Writer.Header().Add("Content-Type", "application/octet-stream;charset=utf-8")

	c.File(fullPath)
}

func (ctrl *AdvertisingsController) Upload(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件至指定目录
		// c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))

}
