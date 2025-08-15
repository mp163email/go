package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
	Email    string             `bson:"email"`
	CreateAt time.Time          `bson:"createAt"`
}

func main() {

	//连接数据库
	uri := "mongodb://root:password123@10.1.9.121:27018/miaopeng-dream?authSource=admin&maxPoolSize=40"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//指定数据库和集合
	collection := client.Database("miaopeng-dream").Collection("user")

	//插入文档
	//insert(err, collection)

	//批量插入文档
	//batchInsert(err, collection)

	//查询单个文档
	findOne(err, collection)

	//查询多个文档
	findMany(err, collection)

	//更新文档
	update(err, collection)

	//删除文档
	delete(err, collection)

	//分页查询
	findByPage(err, collection)

}

func findByPage(err error, collection *mongo.Collection) {
	pageSize := int64(2)
	pageNumber := int64(1)
	opts := options.Find().
		SetLimit(pageSize).
		SetSkip((pageNumber - 1) * pageSize).
		SetSort(bson.M{"createAt": -1})
	cur, err := collection.Find(context.TODO(), bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var pageResults []User
	if err = cur.All(context.Background(), &pageResults); err != nil {
		log.Fatal(err)
	}
	fmt.Println("分页查询结果")
	for _, user := range pageResults {
		fmt.Printf("user: %+v\n", user)
	}
}

func delete(err error, collection *mongo.Collection) {
	one, err := collection.DeleteOne(context.TODO(), bson.M{"name": "lisi"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Delete %v %v \n", one, one.DeletedCount)
}

func update(err error, collection *mongo.Collection) {
	filter := bson.M{"name": "zhangsan"}
	update := bson.M{"$set": bson.M{"age": 30, "email": "zhangsan@qq.com"}}
	one, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v document(s)，, docId= %v \n", one.ModifiedCount, one.UpsertedID)
}

func findMany(err error, collection *mongo.Collection) {
	cursor, err := collection.Find(context.TODO(), bson.M{"age": bson.M{"$gt": 18}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO()) //这里为什么要关闭cursor
	var results []User
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, user := range results {
		fmt.Printf("user: %+v\n", user)
	}
}

func findOne(err error, collection *mongo.Collection) {
	var result User
	err = collection.FindOne(context.TODO(), bson.M{"name": "zhangsan"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
}

func insert(err error, collection *mongo.Collection) bool {
	user1 := User{
		Name:     "zhangsan",
		Age:      18,
		Email:    "zhangsan@163.com",
		CreateAt: time.Now(),
	}
	one, err := collection.InsertOne(context.TODO(), user1)
	if err != nil {
		return true
	}
	fmt.Printf("Inserted a single document: %+v, id: %v\n", one, one.InsertedID)
	return false
}

func batchInsert(err error, collection *mongo.Collection) {
	//批量插入
	users := []interface{}{
		User{
			Name:     "lisi",
			Age:      28,
			Email:    "lisi@163.com",
			CreateAt: time.Now(),
		},
		User{
			Name:     "wangwu",
			Age:      38,
			Email:    "wangwu@163.com",
			CreateAt: time.Now(),
		},
	}
	many, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted many documents: %+v, ids: %v\n", many, many.InsertedIDs)
}
