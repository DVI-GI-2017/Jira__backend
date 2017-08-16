package db

import "github.com/DVI-GI-2017/Jira__backend/models"

var FakeLabels = models.Labels{
	models.Label{Name: "to-do"},
	models.Label{Name: "doing"},
	models.Label{Name: "done"},
	models.Label{Name: "bug"},
}