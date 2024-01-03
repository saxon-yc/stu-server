package apiproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	errorcode "student-server/pkg/errors"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type SkuFilterItem struct {
	AttrID    string `json:"attr_id"`
	Operator  string `json:"operator"`
	AttrValue string `json:"attr_value"`
	Name      string `json:"name"`
}

type SkuItem struct {
	SkuID   string          `json:"sku_id"`
	SpecID  string          `json:"spec_id"`
	Filters []SkuFilterItem `json:"filters"`
}

type SkusResponse struct {
	RectCode int         `json:"ret_code"`
	Spec     interface{} `json:"spec"`
	Skus     []SkuItem   `json:"skus"`
	Total    string      `json:"total"`
	TraceID  string      `json:"trace_id"`
}

type NodeTemplate struct {
	Id            uint   `json:"id,omitempty"`
	InstanceClass int32  `json:"instance_class"`
	InstanceType  string `json:"instance_type"`
	Cpu           int32  `json:"cpu"`
	CpuModel      string `json:"cpu_model"`
	Memory        int32  `json:"memory"`
	VolumeClass   int32  `json:"volume_class"`
	VolumeType    int32  `json:"volume_type"`
	VolumeSize    int32  `json:"volume_size"`
	Gpu           int32  `json:"gpu"`
	GpuClass      int32  `json:"gpu_class"`
	OsVolumeClass int32  `json:"os_volume_class"`
	OsVolumeSize  int32  `json:"os_volume_size"`
}
type NodeSpec struct {
	SpecName string `json:"spec_name"`
	Role     string `json:"role"`
	Default  bool   `json:"default"`
	NodeTemplate
}

type NodeSpecsInput struct {
	Role string `json:"role,omitempty"`
}

type NodeSpecsOutput struct {
	NodeSpecs []NodeSpec `json:"node_specs,omitempty"`
}

func describeOtherService(url string, input interface{}) (output []byte, err error) {
	jsonData, err := json.Marshal(input)
	if err != nil {
		return output, errorcode.New(errorcode.QUERY_OTHER_SERVICE_CODE, "QueryOtherService", err.Error())
	}
	r := bytes.NewReader(jsonData)
	// header := c.Request.Header
	// httpHeader := make(http.Header)
	// for key, values := range header {
	// 	for _, value := range values {
	// 		httpHeader.Add(key, value)
	// 	}
	// }

	resp, err := http.Post(url, "application/json", r)

	if err != nil {
		// 如果请求失败，返回错误信息给前端
		return output, errorcode.New(errorcode.QUERY_OTHER_SERVICE_CODE, "QueryOtherService", err.Error())
	}
	defer resp.Body.Close()

	// 读取响应数据
	output, err = io.ReadAll(resp.Body)
	return output, err
}

func (h *proxyHandle) QuerySkus() gin.HandlerFunc {
	// url := viper.GetString("prod_center.addr") + viper.GetString("prod_center.skus_url")
	prodCenter := viper.GetString("prod_center")
	url := fmt.Sprintf("%s%s", viper.GetString(prodCenter+".addr"), viper.GetString(prodCenter+".skus_url"))
	input := map[string]interface{}{
		"prod_id":   "app-o6lvbkhm",
		"region_id": []string{"testing"},
		"spec_id":   "instance",
		"status":    []string{"sale"},
		"version":   "latest",
		"offset":    0,
		"limit":     10240,
	}

	return func(c *gin.Context) {
		resp, err := describeOtherService(url, input)
		if err != nil {
			// 如果读取响应数据失败，返回错误信息给前端
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		var skuData SkusResponse
		// var nodeSpecs []map[string]interface{}
		var nodeSpecs []NodeSpec
		json.Unmarshal(resp, &skuData)

		for k1, v1 := range skuData.Skus {
			var nodeSkusMap NodeSpec
			nodeSkusMap.Id = uint(k1)
			for _, v2 := range v1.Filters {
				switch v2.AttrID {
				case "role":
					nodeSkusMap.Role = v2.AttrValue
				case "instance_type_id":
					nodeSkusMap.InstanceType = v2.AttrValue
				case "instance_class":
					value, _ := strconv.ParseInt(v2.AttrValue, 10, 32)
					nodeSkusMap.InstanceClass = int32(value)
					// value, _ := strconv.Atoi(v2.AttrValue)
					// nodeSkusMap.InstanceClass = int32(value)
				case "cpu":
					value, _ := strconv.Atoi(v2.AttrValue)
					nodeSkusMap.Cpu = int32(value)
				case "cpu_model":
					nodeSkusMap.CpuModel = v2.AttrValue
				case "memory":
					value, _ := strconv.Atoi(v2.AttrValue)
					nodeSkusMap.Memory = int32(value)
				case "gpu":
					value, _ := strconv.Atoi(v2.AttrValue)
					nodeSkusMap.Gpu = int32(value)
				case "gpu_class":
					value, _ := strconv.Atoi(v2.AttrValue)
					nodeSkusMap.GpuClass = int32(value)
				}
			}
			nodeSpecs = append(nodeSpecs, nodeSkusMap)
		}
		sort.Slice(nodeSpecs, func(i, j int) bool {
			if nodeSpecs[i].Cpu > nodeSpecs[j].Cpu {
				return false
			} else if nodeSpecs[i].Cpu < nodeSpecs[j].Cpu {
				return true
			} else {
				if nodeSpecs[i].Memory > nodeSpecs[j].Memory {
					return false
				} else {
					return true
				}
			}
		})

		c.JSON(http.StatusOK, NodeSpecsOutput{
			NodeSpecs: nodeSpecs,
		})
		// ctx.Data(http.StatusOK, resp.Header.Get("Content-Type"), nodeSpecReponse)
	}
}
