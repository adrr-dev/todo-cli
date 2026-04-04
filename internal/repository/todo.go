// Package repository contains the model and repo
package repository

import "gorm.io/gorm"

type Todo struct {
	gorm.Model // This adds ID, CreatedAt, etc.
	Content    string
	Completed  bool
}
