package repository

import "github.com/tonouchi510/Jeeek/domain"

type ActivityRepository interface {
	InsertAll(uid string, activities []*domain.Activity) (success int, err error)
	Insert(uid string, activity domain.Activity) error
}
