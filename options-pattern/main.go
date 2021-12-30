package main

import "fmt"

/* 0. 定义Gender常量 */
type GenderType int

const (
	Man    GenderType = 1
	Woman  GenderType = 2
	Secret GenderType = 3
)

/* 1. 定义functional options的两个基本要素，一是option function type，二是基本的struct */
type StudentOptions func(*Student) // 注意这里的参数是要更改的结构体指针，因为后面要遍历options列表，调用方式是opt(s)

type Student struct {
	Name   string
	Age    int
	Gender GenderType
}

/* 2. 接下来这个是最关键的部分，注意主要起作用的是返回值 */
func WithName(name string) StudentOptions {
	return func(s *Student) { // opt(s)本质上利用的是此处的返回值函数
		s.Name = name
	}
}

func WithAge(age int) StudentOptions {
	return func(s *Student) {
		s.Age = age
	}
}

func WithGender(gender GenderType) StudentOptions {
	return func(s *Student) {
		s.Gender = gender
	}
}

/* 3. 在这里遍历并调用option function */
func NewStudent(options ...StudentOptions) *Student {
	// 建立空结构体
	s := &Student{}

	// 遍历options函数，并调用函数
	for _, opt := range options {
		opt(s)
	}
	return s
}

func main() {
	student := NewStudent(
		WithName("thinszx"),
		WithAge(20),
		WithGender(Man),
	)
	fmt.Printf("Name: %v, Age: %v, Gender: %v", student.Name, student.Age, student.Gender)
}

// reference: https://www.sohamkamani.com/golang/options-pattern/
