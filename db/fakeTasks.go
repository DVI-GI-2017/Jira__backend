package db

import (
	"github.com/DVI-GI-2017/Jira__backend/models"
	"time"
)

var FakeTasks = models.Tasks{
	models.Task{
		Title:       "Meaningful commit message",
		Description: "TODO: write meaningful commit message",
		Initiator:   &models.User{},
		Assignee:    &models.User{},
		Labels:      models.Labels{},
		CreatedAt:   time.Date(2015, 12, 3, 23, 00, 00, 00, time.UTC),
		UpdatedAt:   time.Date(2015, 12, 3, 23, 00, 00, 00, time.UTC)},

	models.Task{
		Title:       "I tried",
		Description: "this doesn't really make things faster, but I tried",
		Initiator:   &models.User{},
		Assignee:    &models.User{},
		Labels:      models.Labels{},
		CreatedAt:   time.Date(2016, 2, 3, 23, 00, 00, 00, time.UTC),
		UpdatedAt:   time.Date(2016, 2, 3, 23, 00, 00, 00, time.UTC)},

	models.Task{
		Title:       "Engineer",
		Description: "Trust me, I'm an engineer!... What the f*ck did just happened here?",
		Initiator:   &models.User{},
		Assignee:    &models.User{},
		Labels:      models.Labels{},
		CreatedAt:   time.Date(2014, 1, 21, 00, 00, 00, 00, time.UTC),
		UpdatedAt:   time.Date(2014, 1, 21, 00, 00, 00, 00, time.UTC)},
}
