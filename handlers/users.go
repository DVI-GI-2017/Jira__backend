package handlers

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

func AllUsers(w http.ResponseWriter, _ *http.Request) {
	tools.JsonResponse(db.FakeUsers, w)
}
