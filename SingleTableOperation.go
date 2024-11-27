package main

import (
	"gorm.io/gorm"
)

// 单表操作
type Student struct {
	ID     uint   `gorm:"size:10"`
	Name   string `gorm:"size:16";json:"name"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:128"`
}

func main() {
	////	// 自动迁移并创建表
	//if err := DB.AutoMigrate(&Student{}); err != nil {
	//	panic("自动迁移（创建表）失败，error=" + err.Error())
	//}

	// 添加记录
	//email := "123456789@qq.com"
	//s1 := Student{
	//	Name:   "abc",
	//	Age:    21,
	//	Gender: true,
	//	Email:  &email,
	//}
	//err := DB.Create(&s1).Error
	//fmt.Print(err)

	//批量插入数据
	//abd_email := "123456780@qq.com"
	//acd_email := "123456783@qq.com"
	//students := []*Student{
	//	{Name: "abd", Age: 18, Email: &abd_email, Gender: true},
	//	{Name: "acd", Age: 18, Email: &acd_email, Gender: false},
	//}
	//
	//err := DB.Create(&students).Error
	//fmt.Println(err)

	//var students = []Student{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	//DB.Create(&students)

	//for _, student := range students {
	//	student.ID
	//}

	// 查询
	//var student Student
	// 打印查询日志
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	//DB.First(&student)
	//fmt.Println(student)
	//DB.Take(&student, "name = ?", "acd")
	//fmt.Print(student)

	//student.ID = 3
	//DB.Take(&student)
	//fmt.Print(student)

	// 查询多条记录，并将其转换为json
	//var studentList []Student
	////count := DB.Find(&studentList).RowsAffected //获取信息的条数
	////fmt.Println(count)
	////data, _ := json.Marshal(studentList)
	//DB.Find(&studentList, "name in ?", []string{"abc", "acd"})
	////fmt.Println(string(data)？
	//fmt.Print(studentList)

	//var student Student
	//DB.Take(&student, 2)
	//student.Name = "张三"
	//student.Age = 0
	//// 全字段更新
	//DB.Select("name").Save(&student) // 指定更新name字段，其他字段不更新

	//email := "xxx@qq.com"
	//DB.Model(&Student{}).Where("age = ?", 21).Updates(Student{
	//	Email:  &email,
	//	Gender: true,
	//})
}
