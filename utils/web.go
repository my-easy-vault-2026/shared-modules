package utils

import (
	"encoding/json"
	"strconv"

	"shared-modules/common"
	"shared-modules/logger"

	"github.com/gin-gonic/gin"
)

type Page struct {
	Current  int `json:"current" validate:"min=1,required"`
	PageSize int `json:"pageSize" validate:"max=3000,required"`
}

type PageData[T interface{}] struct {
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Records  T   `json:"records"`
}

type Response[T interface{}] struct {
	Code int    `json:"code,de"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func ReError(c *gin.Context, err error, headers ...map[string]string) {

	if len(headers) != 0 {
		for k, v := range headers[0] {
			c.Writer.Header().Set(k, v)
		}
	}

	reqID, ok := GetMDCValue("reqId")
	if ok {
		c.Writer.Header().Set("Request-Id", reqID)
	}

	businessError, ok := err.(*common.BusinessError)

	if ok {
		c.Set("errCode", businessError.Code())
		c.JSON(200, gin.H{
			"code": businessError.Code(),
			"msg":  businessError.Error(),
			"data": nil,
		})
		logger.Infof("user=%v,ip=%v,host=%v,path=%v,errcode=%v,errmsg=%v", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI(), businessError.Code(), businessError.Error())
		return
	}

	c.JSON(200, gin.H{
		"code": common.CODE_SYSTEM_ERROR,
		"msg":  err.Error(),
		"data": nil,
	})

	logger.Infof("user=%v,ip=%v,host=%v,path=%v,errcode=%v,errmsg=%v", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI(), common.CODE_SYSTEM_ERROR, err.Error())

}

func ReErrorWithData(c *gin.Context, err error, data interface{}, headers ...map[string]string) {

	businessError, ok := err.(*common.BusinessError)

	if ok {
		c.Set("errCode", businessError.Code())
		c.JSON(200, gin.H{
			"code": businessError.Code(),
			"msg":  businessError.Error(),
			"data": data,
		})
		logger.Infof("user=%v,ip=%v,host=%v,path=%v,errcode=%v,errmsg=%v", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI(), businessError.Code(), businessError.Error())
		return
	}

	if len(headers) != 0 {
		for k, v := range headers[0] {
			c.Writer.Header().Set(k, v)
		}
	}

	reqID, ok := GetMDCValue("reqId")
	if ok {
		c.Writer.Header().Set("Request-Id", reqID)
	}

	c.JSON(200, gin.H{
		"code": common.CODE_SYSTEM_ERROR,
		"msg":  err.Error(),
		"data": data,
	})
	logger.Infof("user=%v,ip=%v,host=%v,path=%v,errcode=%v,errmsg=%v", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI(), common.CODE_SYSTEM_ERROR, err.Error())
}

func MReData(c *gin.Context, data interface{}, headers ...map[string]string) {

	c.Set("errCode", 0)
	if len(headers) != 0 {
		for k, v := range headers[0] {
			c.Writer.Header().Set(k, v)
		}
	}

	reqID, ok := GetMDCValue("reqId")
	if ok {
		c.Writer.Header().Set("Request-Id", reqID)
	}

	c.JSON(200, gin.H{
		"code": common.CODE_OK,
		"msg":  "ok",
		"data": data,
	})

	logger.Infof("user=%v,ip=%v,host=%v,path=%v,errcode=1,errmsg=ok", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI())
}

func ReData(c *gin.Context, data interface{}, headers ...map[string]string) {

	c.Set("errCode", 0)
	if len(headers) != 0 {
		for k, v := range headers[0] {
			c.Writer.Header().Set(k, v)
		}
	}

	reqID, ok := GetMDCValue("reqId")
	if ok {
		c.Writer.Header().Set("Request-Id", reqID)
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})

	logger.Infof("user=%v,app_version=%v,ip=%v,host=%v,path=%v,errcode=0,errmsg=ok", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_APP_VERSION), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI())
}

func ReRaw(c *gin.Context, code int, raw interface{}, headers ...map[string]string) error {

	if len(headers) != 0 {
		for k, v := range headers[0] {
			c.Writer.Header().Set(k, v)
		}
	}

	reqID, ok := GetMDCValue("reqId")
	if ok {
		c.Writer.Header().Set("Request-Id", reqID)
	}

	if d, ok := raw.([]byte); ok {
		c.Data(code, "application/json", d)
		logger.Infof("user=%v,ip=%v,host=%v,path=%v,status=%d", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI(), code)
		return nil
	}

	data, err := json.Marshal(raw)
	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}

	c.JSON(code, result)
	logger.Infof("user=%v,ip=%v,host=%v,path=%v,status=%d", c.GetHeader("X-Uid"), c.GetHeader(common.HEADER_X_REAL_IP), c.Request.Host, c.Request.URL.RequestURI(), code)
	return nil
}

func GetUserId(c *gin.Context) (uint64, error) {
	uidStr := c.GetHeader("X-Uid")
	return strconv.ParseUint(uidStr, 10, 64)
}
