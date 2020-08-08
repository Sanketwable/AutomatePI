package filecontent

import (
	"AutomatePI/getauto"
	"fmt"
	"strings"
)
var userStruct string

func UserModels() {
	Models := getauto.GetModelData()
	//fmt.Println(Models)
	for _, str := range Models {
		str[2] = str[2][:strings.IndexByte(str[2], ';')]
		userStruct = userStruct + "`" + str[0] + " " + str[1] + " " + str[2] + "`" + "\n"
	}
	//fmt.Println(userStruct)
	pp := []byte(
		userStruct,
	)
	fmt.Println(string(pp))
}

var UserModelContent = []byte(
	`package models

	import (
		"automatepi/security"
		"errors"
		"html"
		"log"
		"strings"
		"time"
	
		"github.com/badoux/checkmail"
		//"github.com/sqs/goreturns/returns"
	)
	
	//User is a struct
	type User struct {
		ID        uint32    *replace*gorm:"primary_key;auto_increment" json:"id"*replace*
		UserName  string    *replace*gorm:"size:20;not null;unique" json:"username"*replace*
		Email     string    *replace*gorm:"size:50;not null;unique" json:"email"*replace*
		Password  string    *replace*gorm:"size:60;not null" json:"password"*replace*
		CreatedAt time.Time *replace*gorm:"" json:"created_at"*replace*
		UpdatedAt time.Time *replace*gorm:"" json:"updated_at"*replace*
		
	}
	
	//BeforeSave is a func
	func (u *User) BeforeSave() error {
		hashedPassword, err  := security.Hash(u.Password)
		if err != nil {
			log.Fatal("error comes : ", err)
		}
	
		u.Password = string(hashedPassword)
		return nil
	}
	
	//Prepare is a struct
	func (u *User) Prepare() {
		u.ID = 0
		u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
		u.Email = html.EscapeString(strings.TrimSpace(u.Email))
		u.CreatedAt = time.Now()
		u.UpdatedAt = time.Now()
	
	}
	
	//Validate is a func
	func (u *User) Validate(action string) error {
		switch action {
		case "update":
			if u.UserName == "" {
				return errors.New("Required UserName")
			}
			if u.Email == "" {
				return errors.New("Required email")
			}
		
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			return nil
		case "login":
			if u.Email == "" {
				return errors.New("Required email")
			}
		
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			if u.Password == "" {
				return errors.New("Required password")
			}
			return nil
		case "signup" :
			if u.Email == "" {
				return errors.New("Required email")
			}
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			if u.Password == "" {
				return errors.New("Required password")
			}
			return nil
		default:
			if u.UserName == "" {
				return errors.New("Required username")
			}
			if u.Password == "" {
				return errors.New("Required password")
			}
			if u.Email == "" {
				return errors.New("Required email")
			}
		
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				return errors.New("Invalid email")
			}
			return nil
		}
	}`)