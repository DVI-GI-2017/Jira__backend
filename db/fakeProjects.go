package db

import "github.com/DVI-GI-2017/Jira__backend/models"

var FakeProjects = models.Projects{
	models.Project{
		Title:       "android-file-chooser",
		Description: "Create a lightweight file/folder chooser",
		Tasks:       models.Tasks{}},

	models.Project{
		Title:       "krypt-core-c",
		Description: "C implementation of the krypt-core API.",
		Tasks:       models.Tasks{}},

	models.Project{
		Title:       "crosstool-NG",
		Description: "crosstool-NG with support for Xtensa",
		Tasks:       models.Tasks{}},
}
