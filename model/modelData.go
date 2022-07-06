package model

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Key     string `json:"key"`
	Value   string `json:"value"`
	Context string `json:"context"`
	UserKey string `json:"userKey"`
	Salt    []byte `json:"-"`
}
