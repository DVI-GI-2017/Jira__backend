package db

import (
	"time"
	"github.com/DVI-GI-2017/Jira__backend/models"
)

var FakeProjects = models.Projects{
	models.Project{
		Title:       "android-file-chooser",
		Description: "Create a lightweight file/folder chooser",
		Tasks:       models.Tasks{},
		CreatedAt:   time.Date(2010, 12, 3, 23, 00, 00, 00, time.UTC),
		UpdatedAt:   time.Date(2010, 12, 3, 23, 00, 00, 00, time.UTC)},

	models.Project{
		Title:       "krypt-core-c",
		Description: "C implementation of the krypt-core API.",
		Tasks:       models.Tasks{},
		CreatedAt:   time.Date(2007, 2, 3, 13, 00, 00, 00, time.UTC),
		UpdatedAt:   time.Date(2007, 2, 3, 13, 00, 00, 00, time.UTC)},

	models.Project{
		Title:       "crosstool-NG",
		Description: "crosstool-NG with support for Xtensa",
		Tasks:       models.Tasks{},
		CreatedAt:   time.Date(2011, 8, 3, 9, 00, 00, 00, time.UTC),
		UpdatedAt:   time.Date(2011, 8, 3, 9, 00, 00, 00, time.UTC)},
}
