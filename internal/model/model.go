package model

import "time"

type Config struct {
	DataBase struct {
		DBName string `yaml:"DBName"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
	}
	TestDataBase struct {
		DBName string `yaml:"DBName"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
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
