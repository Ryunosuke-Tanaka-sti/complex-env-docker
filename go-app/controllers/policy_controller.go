package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// OPAからポリシーを取得する関数
func getPoliciesFromOPA() ([]string, error) {
	resp, err := http.Get("http://opa:8181/v1/policies")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch policies: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var policies map[string]interface{}
	if err := json.Unmarshal(body, &policies); err != nil {
		return nil, err
	}

	var policyList []string
	for key := range policies {
		policyList = append(policyList, key)
	}

	return policyList, nil
}

// ポリシー確認APIエンドポイント
func ListPolicies(c *gin.Context) {
	policies, err := getPoliciesFromOPA()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"policies": policies})
}
