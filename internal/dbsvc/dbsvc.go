package dbsvc

import (
	"context"
	"fmt"
	"log"
	"student-server/internal/model"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Servicer interface {
	// admin
	FindUser(params model.DupuserRequest) (result model.UserDb, err error)
	CreateUser(params model.LoginRequest) (err error)
	UpdatePassword(params model.LoginRequest, password string) (result model.UserDb, err error)
	RefreshToken(params model.LoginRequest, token string) (result model.UserDb, err error)

	// student
	FindOneStudentByID(id uint32) (result model.StudentDb)
	FindStudents(params model.QueryStudentRequest) (result model.QueryStudentResponse, err error)
	CreateStudent(params model.CreateStudentRequest) (err error)
	BatchInsertStudents(students []model.StudentDb) (totalInserted int64, err error)
	ChangeStudent(params model.ChangeStudentRequest) (err error)
	DeleteStudent(id uint32) (err error)

	// tag
	FindTags(params model.QueryTagRequest) (result model.QueryTagResponse, err error)
	CreateTag(params model.CreateTagRequest) (err error)
	ChangeTag(params model.ChangeTagRequest) (err error)
	DeleteTag(params model.DeleteTagRequest) (err error)
	FindOneTagByLabel(label string) (result model.TagDb, err error)

	RedisDataBase
}

type basicService struct {
	gdb *gorm.DB
	rdb *redis.Client
	ctx context.Context
}

var dbServiceInstance *basicService

func newPsql() *gorm.DB {
	dbName := viper.GetString("database.name")
	host := viper.GetString(dbName + ".host")
	port := viper.GetInt(dbName + ".port")
	if port == 0 {
		port = 5432
	}
	user := viper.GetString(dbName + ".user")
	password := viper.GetString(dbName + ".password")
	// password = crypto.AESDecryptWithSaltBase64(password)
	fmt.Printf("dbName: %v", dbName)
	gdb := NewGormContext(host, port, user, password, dbName)
	// 自动迁移
	if viper.GetBool("database.migrate") {
		err := gdb.AutoMigrate(&model.UserDb{}, &model.StudentDb{}, &model.TagDb{})
		if err != nil {
			log.Fatalf("migrate table error[%s] exited \n", err)
		}
		log.Print("Init database success \n")
	}
	return gdb
}

func newRedis() (*redis.Client, context.Context) {

	stu_rdb := viper.GetString("redis.name")
	fmt.Printf("rdb_name: %v \n", stu_rdb)
	addr := fmt.Sprintf("%s:%s", viper.GetString(stu_rdb+".host"), viper.GetString(stu_rdb+".port"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	s, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Init redis filed: %v. \n", err)
	}
	log.Printf("Init redis success: %v. \n", s)

	return rdb, ctx
}
func NewDbService() Servicer {
	gdb := newPsql()
	rdb, ctx := newRedis()

	dbServiceInstance = &basicService{gdb: gdb, rdb: rdb, ctx: ctx}

	return dbServiceInstance
}
