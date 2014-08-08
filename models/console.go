package models

import (
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    "strconv"
    "time"
)

const (
    _DBAlias  = "default"
    _DBDriver = "mysql"
    _DBConn   = "doolp:PASSWORD@tcp(127.0.0.1:3306)/doolpweb?charset=utf8"
)

type Category struct {
    Id              int64
    Title           string
    Created         time.Time `orm:"index"`
    Views           int64     `orm:"index"`
    TopicTime       time.Time `orm:"index"`
    TopicCount      int64
    TopicLastUserID int64
}

type Servers struct {
    Id          int64
    Uid         int64
    Name        string
    Description string `orm:"size(300)"`
    Attachment  string
    //Created         time.Time `orm:"index"`
    //Updated         time.Time `orm:"index"`
    Views  int64 `orm:"index"`
    Author string
    //DelteTime       time.Time `orm:"index"`
    Cpu    int64
    Memory int64
}

func RegisterDB() {
    orm.RegisterModel(new(Category), new(Servers))
}

func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)
    // 参数1        数据库的别名，用来在ORM中切换数据库使用
    // 参数2        driverName
    // 参数3        对应的链接字符串
    // 参数4(可选)  设置最大空闲连接
    // 参数5(可选)  设置最大数据库连接 (go >= 1.2)
    // doc: https://github.com/go-sql-driver/mysql
    maxIdle := 30
    maxConn := 30
    orm.RegisterDataBase(_DBAlias, _DBDriver, _DBConn, maxIdle, maxConn)
}

func AddServer(sname string) error {
    o := orm.NewOrm()
    server := &Servers{Name: sname}

    qs := o.QueryTable("servers")
    err := qs.Filter("name", sname).One(server)
    if err == nil {
        return err
    }

    _, err = o.Insert(server)
    if err != nil {
        return err
    }

    return nil
}

func DelServer(id string) error {
    cid, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        return err
    }

    o := orm.NewOrm()
    server := &Servers{Id: cid}
    _, err = o.Delete(server)
    return err

}

func GetAllServerNmaes() ([]*Servers, error) {
    o := orm.NewOrm()
    server := make([]*Servers, 0)
    qs := o.QueryTable("servers")
    _, err := qs.All(&server)
    return server, err
}
