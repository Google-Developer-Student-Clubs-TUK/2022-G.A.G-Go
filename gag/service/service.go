package service

import (
	"gag.com/model"
)

type testService struct {
	UserRepository   model.UserRepository
	DeviceRepository model.DeviceRepository
	EclassRepository model.EclassRepository
	PostRepository   model.PostRepository
}

type userService struct {
	UserRepository   model.UserRepository
	DeviceRepository model.DeviceRepository
	EclassRepository model.EclassRepository
}

type subjectService struct {
	PostRepository    model.PostRepository
	CommentRepository model.CommentRepository
}

type USConfig struct {
	UserRepository   model.UserRepository
	DeviceRepository model.DeviceRepository
	EclassRepository model.EclassRepository
}

type SSConfig struct {
	PostRepository    model.PostRepository
	CommentRepository model.CommentRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &userService{
		UserRepository:   c.UserRepository,
		DeviceRepository: c.DeviceRepository,
		EclassRepository: c.EclassRepository,
	}
}

func NewSubjectService(c *SSConfig) model.SubjectService {
	return &subjectService{
		PostRepository:    c.PostRepository,
		CommentRepository: c.CommentRepository,
	}
}
