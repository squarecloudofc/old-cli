package rest

import (
	"fmt"
)

// API Version and API URL that squarego will use
var (
	APIVersion = 1
	APIURL     = fmt.Sprintf("https://api.squarecloud.app/v%d", APIVersion)
)

var (
	EndpointUserSelfInfo = "/public/user"
	EndpointUserInfo     = func(userId string) string { return fmt.Sprintf("/public/user/%s", userId) }

	EndpointApplicationStatus = func(appId string) string { return fmt.Sprintf("/public/status/%s", appId) }
	EndpointApplicationBackup = func(appId string) string { return fmt.Sprintf("/public/backup/%s", appId) }
	EndpointApplicationLogs   = func(appId string, fullLogs bool) string {
		var full = ""
		if fullLogs {
			full = "full-"
		}

		return fmt.Sprintf("/public/%slogs/%s", full, appId)
	}
	EndpointApplicationStart   = func(appId string) string { return fmt.Sprintf("/public/start/%s", appId) }
	EndpointApplicationStop    = func(appId string) string { return fmt.Sprintf("/public/stop/%s", appId) }
	EndpointApplicationRestart = func(appId string) string { return fmt.Sprintf("/public/restart/%s", appId) }
	EndpointApplicationUpload  = "/public/upload"
	EndpointApplicationCommit  = func(appId string) string { return fmt.Sprintf("/public/commit/%s", appId) }
	EndpointApplicationDelete  = func(appId string) string { return fmt.Sprintf("/public/delete/%s", appId) }
)
