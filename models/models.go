package models

import (
	"time"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"github.com/astaxie/beego"
)

const (
	_DB_NAME = "data/beeblog.db"
	_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64                              //id默认主键
	Title           string                             //orm默认认为是255字节
	Created         time.Time        `orm:"index"`     //时间,添加索引
	Views           int64                `orm:"index"` //浏览次数,添加索引
	TopicTime       time.Time        `orm:"index"`     //发布时间
	TopicCount      int64                              //分类中有多少文章
	TopicLastUserId int64                              //最后一个操作这个分类的用户id
}

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:size(5000)` //设置内容长度5000字节
	Attachment       string                  //附件
	Created          time.Time        `orm:"index"`
	Updated          time.Time        `orm:"index"`
	Views            int64                `orm:"index"`
	Author           string
	ReplyTime        time.Time        `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

func RegisterDB() {
	//创建目录
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	//创建模型
	//注册模型
	orm.RegisterModel(new(Category), new(Topic))
	//初始driver
	orm.RegisterDriver(_DRIVER, orm.DRSqlite)
	//注册数据库
	//允许注册多个数据库,但是必须有一个默认数据库  叫default
	orm.RegisterDataBase("default", _DRIVER, _DB_NAME, 10)        //最后一位为最大连接数
}

func AddCate(name string) error {
	orm := orm.NewOrm()

	cate := &Category{
		Title:name,
		Created:time.Now(),
		TopicTime:time.Now(),
	}

	qs := orm.QueryTable("Category")
	err := qs.Filter("title", name).One(cate)
	beego.Info(cate,111)
	if err == nil {
		return err
	}

	_, err = orm.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCate() (cate []Category, err error) {
	orm := orm.NewOrm()

	cate = make([]Category, 0)
	qs := orm.QueryTable("Category")
	_, err = qs.All(&cate)

	if err != nil {
		return
	}
	return
}

func DelCate(id string) error {
	cid , err := strconv.ParseInt(id,10,64)
	if err != nil{
		return err
	}

	cate := &Category{Id:cid}
	orm := orm.NewOrm()
	//qs := orm.QueryTable("Category")
	_,err = orm.Delete(cate)
	return err
}