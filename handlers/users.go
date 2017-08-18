package handlers

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"net/http"
)

var AllUsers = func(w http.ResponseWriter, _ *http.Request) {
	tools.JsonResponse(db.FakeUsers, w)
}
