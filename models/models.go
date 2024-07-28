package models

type User struct {
	ID      int64  `json : "id" example:"12345"`
	Name    string `json: "name example: "John"`
	Age     int    `json:"age" example: "30"`
	Address string `json:"address" example "Jl.YYYY No.14 , MMM , Inida"`
}

type UserCreateParam struct {
	Name    string `json: "name" example:"JOHN"`
	Age     int    `json:"age" example: "30"`
	Address string `json:"address" exmaple:"Jl.YYYY No.14 , MMM , India"`
}

type UserUpdateParam struct{
	Name    *string `json:"name",omitempty`
	Age     *int 	`json:"age",omitempty`
	Address *string `json:"address",omitempty`
}