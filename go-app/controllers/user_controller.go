package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-api-project/config"
	"gin-api-project/models"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// OPAリクエスト構造体
type OPARequest struct {
	Input struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	} `json:"input"`
}

// OPAレスポンス構造体
type OPAResponse struct {
	Result bool `json:"result"`
}

// OPAのURLを環境変数から取得
func getOPAURL() string {
	url := os.Getenv("OPA_URL")
	if url == "" {
		return "http://opa:8181/v1/data/example/allow" // デフォルト値
	}
	return url
}

// OPAにポリシー判定を問い合わせる関数
func askOPA(opaURL string, reqBody OPARequest) (bool, error) {
	jsonData, _ := json.Marshal(reqBody)
	resp, err := http.Post(opaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var opaResp OPAResponse
	log.Printf("OPA Response: %s", string(body))
	json.Unmarshal(body, &opaResp)

	return opaResp.Result, nil
}

func GetItems(c *gin.Context) {
	// OPAリクエスト作成
	reqBody := OPARequest{
		Input: struct {
			Method string `json:"method"`
			Path   string `json:"path"`
		}{
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
		},
	}

	log.Printf("OPA Request: %v\n", reqBody)

	// OPAに問い合わせ
	opaURL := getOPAURL()
	allowed, err := askOPA(opaURL, reqBody)
	log.Printf("OPA Allowed %v, OPA Error: %s\n", allowed, err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Policy check failed: %v", err)})
		return
	}

	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied by policy"})
		return
	}

	var users []models.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateItem(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateItemとDeleteItemも同様に実装
