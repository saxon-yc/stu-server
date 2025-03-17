package excelizelib

import (
	"fmt"
	"mime/multipart"
	"net/url"
	"strconv"
	errorcode "student-server/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

var (
	defaultHeight = 25.0 //默认行高度
)

type lkExcelExport struct {
	file      *excelize.File
	fileName  string
	sheetName string // 可定义默认sheet名称
}

func NewMyExcel(fileName, sheetName string) *lkExcelExport {
	return &lkExcelExport{file: createFile(sheetName), fileName: fmt.Sprintf("%s.xlsx", fileName), sheetName: sheetName}
}

// ExportToPath 导出基本的表格
func (l *lkExcelExport) ExportToPath(params []map[string]string, data []map[string]interface{}, path string) (string, error) {
	l.export(params, data)
	filePath := path + "/" + l.fileName
	err := l.file.SaveAs(filePath)
	return filePath, err
}

// ExportToWeb 导出到浏览器。此处使用的gin框架 其他框架可自行修改ctx
func (l *lkExcelExport) ExportToWeb(params []map[string]string, data []map[string]interface{}, ctx *gin.Context) {
	l.export(params, data)
	buffer, _ := l.file.WriteToBuffer()
	//设置文件类型
	ctx.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	//设置文件名称
	ctx.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(l.fileName))
	_, _ = ctx.Writer.Write(buffer.Bytes())
}

// 设置首行
func (l *lkExcelExport) writeTop(params []map[string]string) {
	topStyle, _ := l.file.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		}})

	var word = 'A'
	//首行写入
	for _, conf := range params {
		title := conf["title"]
		width, _ := strconv.ParseFloat(conf["width"], 64)
		line := fmt.Sprintf("%c1", word)
		//设置标题
		_ = l.file.SetCellValue(l.sheetName, line, title)
		//列宽
		_ = l.file.SetColWidth(l.sheetName, fmt.Sprintf("%c", word), fmt.Sprintf("%c", word), width)
		//设置样式
		_ = l.file.SetCellStyle(l.sheetName, line, line, topStyle)
		word++
	}
}

// 写入数据
func (l *lkExcelExport) writeData(params []map[string]string, data []map[string]interface{}) {
	lineStyle, _ := l.file.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		}})
	//数据写入
	var j = 2 //数据开始行数
	for i, val := range data {
		//设置行高
		_ = l.file.SetRowHeight(l.sheetName, i+1, defaultHeight)
		//逐列写入
		var word = 'A'
		for _, conf := range params {
			valKey := conf["key"]
			line := fmt.Sprintf("%c%v", word, j)
			isNum := conf["is_num"]

			//设置值
			if isNum != "0" {
				valNum := fmt.Sprintf("'%v", val[valKey])
				_ = l.file.SetCellValue(l.sheetName, line, valNum)
			} else {
				_ = l.file.SetCellValue(l.sheetName, line, val[valKey])
			}

			//设置样式
			_ = l.file.SetCellStyle(l.sheetName, line, line, lineStyle)
			word++
		}
		j++
	}
	//设置行高 尾行
	_ = l.file.SetRowHeight(l.sheetName, len(data)+1, defaultHeight)
}

func (l *lkExcelExport) export(params []map[string]string, data []map[string]interface{}) {
	l.writeTop(params)
	l.writeData(params, data)
}

func createFile(sheetName string) *excelize.File {
	f := excelize.NewFile()
	// 创建一个默认工作表
	index, _ := f.NewSheet(sheetName)
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	return f
}

func ReadExcel(file multipart.File) (rows *excelize.Rows, err error) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return nil, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", "Excel解析失败")
	}

	sheetName := f.GetSheetName(0)
	if sheetName == "" {
		return nil, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", "未找到有效工作表")
	}
	fmt.Printf("sheetName: %v\n", sheetName)
	rows, err = f.Rows(sheetName)
	if err != nil {
		return nil, errorcode.New(errorcode.INVALID_FILE_UNCORRENT_CODE, "ImportStudentsAPI", "工作表读取失败")
	}
	return rows, nil
}
