package objj

import (
	"golangAPI/database"
	"log"
)

type User struct { //DB : Users
	UUID int    `gorm:"column:UUID;primaryKey;AUTO_INCREMENT;not null"`  // 編號
	Name string `json:"Name" form:"Name" gorm:"column:Name;uniqueIndex"` // 資料庫名稱
	Rank string `json:"Rank" form:"Rank" gorm:"column:Rank"`             // 稀有階級
	Job  string `json:"Job" form:"Job" gorm:"column:Job"`                // 類別
	Camp string `json:"Camp" form:"Camp" gorm:"column:Camp"`             // 陣營
	Base string `json:"Base" form:"Base" gorm:"column:Base"`             // 基本屬性
	Grow string `json:"Grow" form:"Grow" gorm:"column:Grow"`             // 成長屬性
	//Resistance int    `json:"UserResistance" binding:"required" ` // 基本抗性
}

// get User
func FindAllUser() []User {
	var users []User
	database.DBConnect.Find(&users)
	return users
}

func FindByUserId(UserUUID string) User {
	var user User
	database.DBConnect.Where("UUID = ?", UserUUID).First(&user) //"UUID" in the bracket is database primary key's UUID, and fisrt parameter is find user table
	return user
}

// Post
func CreateUser(user User) (err error) {
	d := database.DBConnect.Create(&user) //using d to catch the create's info
	// database.DBConnect.Where("Name = ?", user).Find(&user)
	// if user.Name !=nil{
	// }
	if err = d.Error; err != nil {
		return err
	}
	return err
}

// Delete
func DeleteUser(UserUUID string) bool { //Use the boolean to check whether the delete was done
	user := User{}
	result := database.DBConnect.Where("UUID = ?", UserUUID).Delete(&user)
	log.Println(result)
	return result.RowsAffected > 0 // RowAffected store the numbers of executions
}

// update
func UpdateUser(UserUUID string, user User) User { // return the updated data (user)
	database.DBConnect.Model(&user).Where("UUID = ?", UserUUID).Updates(user)
	return user

}
