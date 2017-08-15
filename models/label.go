package models

import "image/color"

type Label struct {
	Name  string `json:"name"`
	Color color.RGBA `json:"color"`
}

type Labels []Label
