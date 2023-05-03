package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	id           string    `json:"id" gorm:"text;not null;default:null`
	name         string    `json:"name"gorm:"text;not null;default:null`
	salary       string    `json:"salary" gorm:"text;not null;default:null`
	projects     *Projects `json:"projects" gorm:"text;not null;default:null`
	technologies string    `json:"technologies" gorm:"text;not null;default:null`
	manager      *Manager  `json:"manager" gorm:"text;not null;default:null`
}

type Projects struct {
	gorm.Model
	id                  string `json:"id" gorm:"text;not null;default:null`
	projectname         string `json:"projectname" gorm:"text;not null;default:null`
	billing             string `json:"billing" gorm:"text;not null;default:null`
	duration            string `json:"duration" gorm:"text;not null;default:null`
	noofemployeeworking int    `json:"noofemployeeworking" gorm:"text;not null;default:null`
}

type Manager struct {
	gorm.Model
	id   string `json:"id" gorm:"text;not null;default:null`
	Name string `json:"name" gorm:"text;not null;default:null`
}
