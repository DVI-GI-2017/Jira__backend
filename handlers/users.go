package handlers

import (
	"Jira__backend/tools"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"net/http"
)

var AllUsers = GetOnly(
	func(w http.ResponseWriter, _ *http.Request) {
		tools.JsonResponse(db.FakeUsers, w)
	})
