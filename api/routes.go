package api

import (
	"fmt"
	"net/url"
)

const rootPath = "/"
const apiVersion = "v1"
const basePath = "risks"

const GetRisksList = ""
const CreateRisk = ""
const GetRisksById = "{Id}"

const Health = "/_health"

func GetPreparedURLPath(route string) string {
	path, _ := url.JoinPath(rootPath, apiVersion, basePath)
	if route == "" {
		return path
	}
	return fmt.Sprintf("%s/%s", path, route)
}
