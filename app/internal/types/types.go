// Code generated by goctl. DO NOT EDIT.
package types

import ()

type CaptchaRequest struct {
	Email string `json:"email" validate:"email"`
}

type EmailPwdLoginRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"alphanum,min=8,max=16"`
}

type EmailCaptchaLoginRequest struct {
	Email   string `json:"email" validate:"email"`
	Captcha string `json:"captcha" validate:"number,len=6"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"email"`
	Captcha  string `json:"captcha" validate:"number,len=6"`
	Name     string `json:"name" validate:"required,excludesall=;#<>"`
	Password string `json:"password" validate:"alphanum,min=8,max=16"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
	Name         string `json:"name"`
}

type ChangePwdRequest struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"alphanum,min=8,max=16"`
}

type UserNameRequest struct {
	Name string `json:"name" validate:"required,excludesall=;#<>"`
}

type UserNameResponse struct {
	Name string `json:"name"`
}

type Todo struct {
	Id          int64  `json:"id"`
	IsCompleted bool   `json:"isCompleted"`
	Content     string `json:"content"`
	Deadline    int64  `json:"deadline"`
	Repeat      string `json:"repeat"`
}

type AddTodoRequest struct {
	Content    string `json:"content"`
	DueDate    int64  `json:"dueDate"`
	Recurrence int64  `json:"recurrence,range=[0:4]"`
}

type UpdateTodoRequest struct {
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	DueDate    int64  `json:"dueDate"`
	Recurrence int64  `json:"recurrence,range=[0:4]"`
}

type CompleteTodoRequest struct {
	Id int64 `json:"id"`
}

type DelTodoRequest struct {
	Id int64 `json:"id"`
}

type QueryTodoRequest struct {
	Cursor      int64 `form:"cursor"`
	PageSize    int64 `form:"pageSize"`
	IsCompleted bool  `form:"isCompleted"`
}

type QueryTodoResponse struct {
	Todos []Todo `json:"todos"`
}
