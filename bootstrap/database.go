package bootstrap

import (
	"fmt"
	"ownapihub/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initdatabase() {
	username := pkg.Get("database.username") //账号
    password := pkg.Get("database.password") //密码
    host := pkg.Get("database.host") //数据库地址，可以是Ip或者域名
    port := pkg.Get("database.port") //数据库端口
    Dbname := pkg.Get("database.Dbname") //数据库名
    timeout := pkg.Get("database.timeout") //连接超时，10秒
    //拼接dsn参数，这里使用Sprintf动态拼接
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	fmt.Println(dsn)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	 if err != nil {
		 panic("连接数据库失败, error=" + err.Error())
	 }
 
	 sqlDB, _ := db.DB()
 
	 //设置数据库连接池参数
	 sqlDB.SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	 sqlDB.SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

