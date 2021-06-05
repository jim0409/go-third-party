package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/jimweng/crawler/pipeline/grpcproto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
)

var (
	port = ":50051"
	OpDB = NewDBConfiguration("root", "secret", "mysql", "mysql", "3306", os.Getenv("DBADD")).NewDBConnection()
)

type DBConfig struct {
	User      string
	Password  string
	DBType    string
	DBName    string
	DBAddress string
	DBPort    string
	DBUri     string
}

type OperationDatabase struct {
	DB *gorm.DB
}

type OPDB interface {
	queryWithName(name string) string
	Closed() error
	debug()
}

func (db *OperationDatabase) Closed() error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("Error happended while closing db : %v\n", err)
	}
	log.Fatalln("Going to close DB")
	return nil
}

// 透過使用Debug()可以轉譯語言為SQL語法
func (db *OperationDatabase) debug() {
	db.DB = db.DB.Debug()
}

type PKGContent struct {
	// gorm.Model
	Name     string `gorm:"primary_key"`
	Parent   string `gorm:"primary_key"`
	Synopsis string
	Href     string
}

// Read
func (db *OperationDatabase) queryWithName(name string) string {
	queryNameInfo := db.DB.Select([]string{"parent", "href", "synopsis"}).Where("name = ?", name).Find(&PKGContent{})
	if err := queryNameInfo.Error; err != nil {
		return fmt.Sprintf("Can't find the parent with "+name, err)
	}

	respContext := queryNameInfo.Value.(*PKGContent)

	return fmt.Sprintf("{name:'%v',parent:'%v',href:'%v',synopsis:'%v'}", name, respContext.Parent, respContext.Href, respContext.Synopsis)
}

func (dbc *DBConfig) NewDBConnection() OPDB {
	db, err := gorm.Open(dbc.DBType, dbc.DBUri)
	if err != nil {
		panic(err)
	}
	return &OperationDatabase{DB: db}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer OpDB.Closed()
}

func NewDBConfiguration(user string, password string, dbtype string, dbname string, dbport string, dbaddress string) *DBConfig {
	return &DBConfig{
		User:      user,
		Password:  password,
		DBType:    dbtype,
		DBName:    dbname,
		DBPort:    dbport,
		DBAddress: dbaddress,
		DBUri:     user + ":" + password + "@tcp(" + dbaddress + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local",
	}
}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Server Received: %v", in.Name)

	return &pb.HelloReply{
		Message: OpDB.queryWithName(in.Name),
	}, nil
}
