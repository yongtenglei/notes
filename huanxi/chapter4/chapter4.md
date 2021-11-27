# GORM

## Quick Start

连接数据库

```go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	dsn := "root:123456@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

```

使用 logger 打印日志信息, 设置 Info 的日志级别

```go
package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

```

## 创建模型 自动迁移

声明模型时, 按照约定数据库的名称会变为结构体的复数, 例如, Tag 在数据库中为 tags. 使用 TableName 覆盖约定.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	dsn := "root:123456@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	fmt.Println("Init finished")
}

```

<div align=center><img src="https://tva1.sinaimg.cn/large/006cK6rNgy1gwem4m4fz2j30jn0cpjwg.jpg"></div>

## GORM 的字段标签, 大小写不敏感

<div align=center>
<table>
<thead>
<tr>
<th>标签名</th>
<th>说明</th>
</tr>
</thead>
<tbody><tr>
<td>column</td>
<td>指定 db 列名</td>
</tr>
<tr>
<td>type</td>
<td>列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：<code>not null</code>、<code>size</code>, <code>autoIncrement</code>… 像 <code>varbinary(8)</code> 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：<code>MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT</code></td>
</tr>
<tr>
<td>size</td>
<td>指定列大小，例如：<code>size:256</code></td>
</tr>
<tr>
<td>primaryKey</td>
<td>指定列为主键</td>
</tr>
<tr>
<td>unique</td>
<td>指定列为唯一</td>
</tr>
<tr>
<td>default</td>
<td>指定列的默认值</td>
</tr>
<tr>
<td>precision</td>
<td>指定列的精度</td>
</tr>
<tr>
<td>scale</td>
<td>指定列大小</td>
</tr>
<tr>
<td>not null</td>
<td>指定列为 NOT NULL</td>
</tr>
<tr>
<td>autoIncrement</td>
<td>指定列为自动增长</td>
</tr>
<tr>
<td>autoIncrementIncrement</td>
<td>自动步长，控制连续记录之间的间隔</td>
</tr>
<tr>
<td>embedded</td>
<td>嵌套字段</td>
</tr>
<tr>
<td>embeddedPrefix</td>
<td>嵌入字段的列名前缀</td>
</tr>
<tr>
<td>autoCreateTime</td>
<td>创建时追踪当前时间，对于 <code>int</code> 字段，它会追踪秒级时间戳，您可以使用 <code>nano</code>/<code>milli</code> 来追踪纳秒、毫秒时间戳，例如：<code>autoCreateTime:nano</code></td>
</tr>
<tr>
<td>autoUpdateTime</td>
<td>创建/更新时追踪当前时间，对于 <code>int</code> 字段，它会追踪秒级时间戳，您可以使用 <code>nano</code>/<code>milli</code> 来追踪纳秒、毫秒时间戳，例如：<code>autoUpdateTime:milli</code></td>
</tr>
<tr>
<td>index</td>
<td>根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 索引 获取详情</td>
</tr>
<tr>
<td>uniqueIndex</td>
<td>与 <code>index</code> 相同，但创建的是唯一索引</td>
</tr>
<tr>
<td>check</td>
<td>创建检查约束，例如 <code>check:age &gt; 13</code>，查看 约束 获取详情</td>
</tr>
<tr>
<td>&lt;-</td>
<td>设置字段写入的权限， <code>&lt;-:create</code> 只创建、<code>&lt;-:update</code> 只更新、<code>&lt;-:false</code> 无写入权限、<code>&lt;-</code> 创建和更新权限</td>
</tr>
<tr>
<td>-&gt;</td>
<td>设置字段读的权限，<code>-&gt;:false</code> 无读权限</td>
</tr>
<tr>
<td>-</td>
<td>忽略该字段，<code>-</code> 无读写权限</td>
</tr>
<tr>
<td>comment</td>
<td>迁移时为字段添加注释</td>
</tr>
</tbody></table>
</div>

## 插入

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	fmt.Println("=========单行插入===========")
	fmt.Println("==========结构体============")
	t1 := Tag{Name: "t1"}
	result := db.Create(&t1)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t2 := map[string]interface{}{"name": "t2"}
	result = db.Model(&Tag{}).Create(t2)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	fmt.Println("=========多行插入===========")
	fmt.Println("==========结构体============")
	t3s := []Tag{
		{Name: "t3"},
		{Name: "t4"},
		{Name: "t5"},
	}
	result = db.Create(&t3s)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t4s := []map[string]interface{}{
		{"name": "t6"},
		{"name": "t7"},
		{"name": "t8"},
	}
	result = db.Model(&Tag{}).Create(&t4s)
	fmt.Println(result.Error)        // 返回 error
	fmt.Println(result.RowsAffected) // 返回插入记录的条数

	fmt.Println("=========多行插入InBatch===========")
	fmt.Println("==========结构体============")
	t5s := []Tag{
		{Name: "t9"},
		{Name: "t10"},
		{Name: "t11"},
	}
	result = db.CreateInBatches(&t5s, 2) // 一次2条数据进行插入
	fmt.Println(result.Error)            // 返回 error
	fmt.Println(result.RowsAffected)     // 返回插入记录的条数
	fmt.Println("==========MAP============")
	t6s := []map[string]interface{}{
		{"name": "t12"},
		{"name": "t13"},
		{"name": "t14"},
	}
	result = db.Model(&Tag{}).CreateInBatches(t6s, 3) // 一次三条数据插入
	fmt.Println(result.Error)                         // 返回 error
	fmt.Println(result.RowsAffected)                  // 返回插入记录的条数

}

```

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gwen6tlt88j30mc0baahm.jpg"></div>

⚠️ 注意三点:

1. 使用 map 插入数据时, 需要指定 Model.

2. 使用 map 插入数据时, 不会自动更新 created_at, updated_at 等子段. 原因是, gorm 只会记录 map 中的字段.

3. 使用 map 插入数据时, association 不会被调用，且主键也不会自动填充.

## 修改 Update

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

// TableName 会将 Tag 的表名重写为 `my_user`
func (User) TableName() string {
	return "my_user"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{}, User{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	// 创建User
	//u1 := User{Name: "u1", Password: "passwd1"}
	//u2 := User{Name: "u2", Password: "passwd2"}
	//us := []User{u1, u2}
	//db.Create(&us)

	fmt.Println("=========Save===========")
	// Save 会保存所有的字段，即使字段是零值
	var t1 Tag
	db.First(&t1)
	fmt.Println(t1)

	t1.Name = "t1 modified"
	db.Save(&t1)
	fmt.Println(t1)

	fmt.Println("=========Update 更新单列===========")
	var t2 Tag
	db.First(&t2, 2) // id = 2
	fmt.Println(t2)

	db.Model(&Tag{}).Where("id = ?", 2).Update("name", "t2 modified")
	db.First(&t2, 2) // id = 2
	fmt.Println(t2)

	fmt.Println("=========Updates 更新多列===========")
	// Updates 方法支持 struct 和 map[string]interface{} 参数。
	// 当使用 struct 更新时，默认情况下，GORM 只会更新非零值的字段

	fmt.Println("=========结构体===========")
	var u1 User
	db.First(&u1)
	fmt.Println(u1)

	db.Model(&User{}).Where("id = ?", 1).Updates(User{Name: "u1 modified", Password: ""})
	db.First(&u1, 1)
	fmt.Println(u1)

	fmt.Println("=========MAP===========")
	var u2 User
	db.First(&u2, 2)
	fmt.Println(u2)

	db.Model(&User{}).Where("id = ?", 2).Updates(map[string]interface{}{"name": "u2 modified", "password": ""})
	db.First(&u2, 2)
	fmt.Println(u2)
}

```

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gweodqray5j30x907yn2n.jpg"></div>

<div align=center><img src="https://tvax3.sinaimg.cn/large/006cK6rNgy1gweo94yvk0j30py0guajn.jpg"></div>

⚠️ 注意四点 :

1. 新建 User 表, 可以看到 Save 会保存所有值, 即使是 0 值.

2. Update 同样会保存 0 值, 没有在结构体或 map 中的字段, 会被忽略.

3. Updates, 在使用结构体更新时, 不会更新 0 值的字段, map 则会更新所有存在的字段.

如果想使用结构体更新 0 值字段, 使用指针类型或 sql.NullXXX, 如 sql.NullString, sql.NullInt64, 包装一个 string or int64 类型的值, 以及一个 bool 值.

使用 map 更新字段时, 需要指定 model.

4. 指定更新的字段使用 Select, 忽略更新的字段使用 Omit.

## 查询 Select

```go
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Tag struct {
	gorm.Model
	Name string `json:"name"`
}

/* 等价与
type Tag struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt DeletedAt `gorm:"index"`
	Name string `json:"name"`
}*/

// TableName 会将 Tag 的表名重写为 `my_tag`
func (Tag) TableName() string {
	return "my_tag"
}

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

// TableName 会将 Tag 的表名重写为 `my_user`
func (User) TableName() string {
	return "my_user"
}

var db *gorm.DB
var err error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,         // 启用彩色打印
		},
	)

	dsn := "root:1@tcp(127.0.0.1:3306)/huanxi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Tag{}, User{})
	if err != nil {
		panic("AutoMigrate Failed")
	}

}

func main() {
	fmt.Println("=========查询单条数据===========")
	var u1 User
	var u2 User
	var u3 User
	// 获取第一条记录（主键升序）
	db.First(&u1)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	fmt.Println(u1)

	// 获取一条记录，没有指定排序字段
	db.Take(&u2)
	// SELECT * FROM users LIMIT 1;
	fmt.Println(u2)

	// 获取最后一条记录（主键降序）
	db.Last(&u3)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	fmt.Println(u3)

	result := db.First(&u3)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("record not found")
	} else {
		fmt.Println("record found successfully")
	}

	fmt.Println("=========查询多条数据===========")
	// 查询所有数据
	var users []User
	result = db.Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(users)

	result = db.Where("id in ?", []int64{1, 2}).Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(users)

}

```

<div align=center><img src="https://tvax4.sinaimg.cn/large/006cK6rNgy1gwep2lgroej30zk094wk7.jpg"></div>

⚠️:

查询时, 使用结构体需要传入指针或指定模型 Model(). 使用 map 时, 需要指定 Table()

## 删 Delete

```go
// Email 的 ID 是 `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";

db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);
```

⚠️ :

    1. GORM 推荐使用软删除, 如果模型中有 gorm.DeletedAt, 则自动进行软删除.

给模型添加软删除:

```go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```

    2. 查询软删除的数据 Unscoped()

```go
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

    3. 不使用软删除 Unscoped()

```go
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```

<++>
