package model

import "time"

type Config struct {
	DataBase struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"DBName"`
		Port     int    `yaml:"port"`
	} `yaml:"DataBase"`
	TestDataBase struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"DBName"`
		Port     int    `yaml:"port"`
	} `yaml:"TestDataBase"`
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"Server"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Expense struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Cost        int        `json:"cost"`
	CreatedData *time.Time `json:"created_data"`
	CreatorId   int        `json:"creator_id"`
}
