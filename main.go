package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynq"
	"github.com/jmoiron/sqlx"
	"github.com/moon-wind/sparks/model"
	"github.com/moon-wind/sparks/utils"
	"github.com/moon-wind/sparks/utils/golimit"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

var client *http.Client

func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	//f6()
	//f7()
	//f8()
	//f9()
	f10()
}

// multi-post
func f1() {
	data := make(map[string]interface{})
	//data["account"] = "abcc"
	//data["passWord"] = "a123456"
	//data["captchaId"] = "vTaGqftKjdEObFBjeNzh"
	//data["captcha"] = "443872"

	//data["order_id"] = 48
	//data["answer"] = "1"

	client = &http.Client{}
	begin := time.Now()
	//url := "http://172.16.4.114:8888/pen/check"
	url := "http://172.16.4.114:8888/awd_service/rfsc"
	//url := "http://172.16.4.114:9000/api/nine_user/login"
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result := utils.CurlPost(client, url, data)
			fmt.Println(result)
		}()
	}
	wg.Wait()
	fmt.Printf("time consumed: %fs", time.Now().Sub(begin).Seconds())
}

// multi-post
func f2() {
	data := make(map[string]interface{})
	data["order_id"] = 5
	data["answer"] = "8ff6ce00747fae9c17846c8219c75b8f"
	client = &http.Client{}
	begin := time.Now()
	url := "http://172.16.15.89:8888/ctf/check"
	wg := &sync.WaitGroup{}
	g := golimit.NewG(5000)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		g.Run(func() {
			defer wg.Done()
			result := utils.CurlPost(client, url, data)
			fmt.Println(result)
		})
	}
	wg.Wait()
	fmt.Printf("time consumed: %fs", time.Now().Sub(begin).Seconds())
}

// mongo
func f3() {
	var (
		client     = utils.GetMgoCli()
		db         *mongo.Database
		collection *mongo.Collection
		//iResult    *mongo.InsertOneResult
		//id         primitive.ObjectID
		cursor *mongo.Cursor
		err    error
	)
	//record := model.Record{}
	//record.Name = "apple"
	//record.Title = "A"

	record := struct {
		Name string `bson:"name"`
	}{Name: "apple"}

	//2.选择数据库 my_db
	db = client.Database("test")

	//选择表 my_collection
	collection = db.Collection("col")

	//插入某一条数据
	if iResult, err := collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err, iResult)
		return
	}

	cursor, err = collection.Find(context.TODO(), record, options.Find().SetSkip(0), options.Find().SetLimit(2))
	if err != nil {
		fmt.Println(cursor, err)
		return
	}

	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.Record
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}

	//fmt.Println("结果:", cursor)
}

func f4() {
	sshConfig := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("cimer@@2021@1qaz")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         15 * time.Second,
	}

	sshClient, err := ssh.Dial("tcp", "172.16.2.10:22", sshConfig)
	if err != nil {
		panic(err)
	}

	// 建立新会话
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatalf("new session error: %s", err.Error())
	}

	defer session.Close()

	//var b bytes.Buffer
	//session.Stdout = &b
	//if err := session.Run("ls"); err != nil {
	//	panic("Failed to run: " + err.Error())
	//}
	//fmt.Println(b.String())

	//session.Stdout = os.Stdout // 会话输出关联到系统标准输出设备
	//session.Stderr = os.Stderr // 会话错误输出关联到系统标准错误输出设备
	//session.Stdin = os.Stdin   // 会话输入关联到系统标准输入设备
	//modes := ssh.TerminalModes{
	//	ssh.ECHO:          0,     // 禁用回显（0禁用，1启动）
	//	ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
	//	ssh.TTY_OP_OSPEED: 14400, //output speed = 14.4kbaud
	//}
	//if err = session.RequestPty("linux", 32, 160, modes); err != nil {
	//	log.Fatalf("request pty error: %s", err.Error())
	//}
	//if err = session.Shell(); err != nil {
	//	log.Fatalf("start shell error: %s", err.Error())
	//}
	//if err = session.Wait(); err != nil {
	//	log.Fatalf("return error: %s", err.Error())
	//}

	rdb := redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort("127.0.0.1", "6379"),
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return sshClient.Dial(network, addr)
		},
		// Disable timeouts, because SSH does not support deadlines.
		ReadTimeout:  -1,
		WriteTimeout: -1,
	})

	ctx := context.Background()
	//rdb.HSet(ctx, "myhash", map[string]interface{}{"key1-1": "value1", "key1-2": "value2"})
	rdb.HMSet(ctx, "user_test", map[string]string{"name": "test", "age": "20"})
	rdb.HMSet(ctx, "user_test", map[string]string{"name": "test1", "age": "18"})
	////val, err := rdb.HMGet(ctx, "myhash", "key1-2", "key1-2").Result()
	//val := rdb.HGetAll(ctx, "myhash")
	//fields, _ := rdb.HMGet(ctx, "user_test", "name", "age").Result()
	//fmt.Println(val)
	rdb.HDel(ctx, "user_test", "age")
	age, _ := rdb.HGet(ctx, "user_test", "name").Result()
	fmt.Println(age)
}

func f5() {
	var db *sqlx.DB
	dsn := "cimer:cimer@@123@Z02O@tcp(172.16.2.10:3308)/wcwx"
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("connect server failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	type User struct {
		Uid      int    `db:"uid"`
		UserName string `db:"username"`
	}
	sqlStr := "SELECT uid, username FROM member WHERE uid = ?"
	var u User
	if err := db.Get(&u, sqlStr, 1); err != nil {
		fmt.Printf("get data failed, err:%v\n", err)
		return
	}
	//fmt.Printf("uid:%d, name:%s", u.Uid, u.UserName)
	fmt.Println(u.UserName)
}

func f6() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"172.16.2.10:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

func f7() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	//router.GET("/someProtoBuf", returnProto)
	router.Run(":8083")

}

//使用ProtoBuf
//func returnProto(c *gin.Context) {
//	course := []string{"python", "golang", "java", "c++"}
//	user := &proto.Teacher{
//		Name:   "ice_moss",
//		Course: course,
//	}
//	//返回protobuf
//	c.ProtoBuf(http.StatusOK, user)
//}

func moreJSON(c *gin.Context) {
	var msg struct {
		Nmae    string `json:"UserName"`
		Message string
		Number  int
	}
	msg.Nmae = "ice_moss"
	msg.Message = "This is a test of JSOM"
	msg.Number = 101

	c.JSON(http.StatusOK, msg)
}

func sendWelcomeEmail(ctx context.Context, t *asynq.Task) error {
	var p struct {
		// ID for the email recipient.
		UserID int
	}
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Welcome Email to User %d", p.UserID)
	return nil
}

func sendReminderEmail(ctx context.Context, t *asynq.Task) error {
	var p struct {
		// ID for the email recipient.
		UserID int
	}
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf(" [*] Send Reminder Email to User %d", p.UserID)
	return nil
}

func f8() {
	const redisAddr = "172.16.2.10:36379"
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr, Password: "G62m50oigInC30sf"})
	defer client.Close()

	// Create a task with typename and payload.
	payload1, err := json.Marshal(map[string]interface{}{"user_id": 42})
	t1 := asynq.NewTask(
		"send_welcome_email",
		payload1)

	payload2, err := json.Marshal(map[string]interface{}{"user_id": 42})
	t2 := asynq.NewTask(
		"send_reminder_email",
		payload2)

	// Process the task immediately.
	info, err := client.Enqueue(t1)
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// Process the task 2 minutes later.
	info, err = client.Enqueue(t2, asynq.ProcessIn(2*time.Minute))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

}

func f9() {
	const redisAddr = "172.16.2.10:36379"
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr, Password: "G62m50oigInC30sf"},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 10,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("send_welcome_email", sendWelcomeEmail)
	mux.HandleFunc("send_reminder_email", sendReminderEmail)
	// mux.Handle(tasks.TypeImageResize, tasks.NewImageProcessor())
	// ...register other handlers...

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func f10() {
	//// to produce messages
	//topic := "test"
	//partition := 1
	//
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "172.16.2.10:9092", topic, partition)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}
	//
	//conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	//_, err = conn.WriteMessages(
	//	kafka.Message{Value: []byte("one!")},
	//	kafka.Message{Value: []byte("two!")},
	//	kafka.Message{Value: []byte("three!")},
	//)
	//if err != nil {
	//	log.Fatal("failed to write messages:", err)
	//}
	//
	//if err := conn.Close(); err != nil {
	//	log.Fatal("failed to close writer:", err)
	//}

	// -------------------------------------------------
	// to consume messages
	topic := "test"
	partition := 1

	conn, err := kafka.DialLeader(context.Background(), "tcp", "172.16.2.10:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
