## 概要
今からechoでWebアプリを作りたい人が簡単にベースコードを作れるようにするためにある。
publicなので好きにcloneして使ってもらって大丈夫です。

## ローカルの開発環境構築 (local env to dev.)
https://qiita.com/riita10069/items/b654f8cbeaede51d4b0e
にまとまっているので参考にしてください。

## 本番環境


## echoの使い方 (How to use echo.)

### URLparams
#### GET
pathで渡される場合
(when pass data using path)


```
e.GET("/users/:id", getUser)

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}
```


#### POST
クエリパラメーターで渡される場合
(when pass data using query)

```
type User struct {
  ID: seq `json:id`
}
  
e.POST("/users", createUser)

func createUser(c echo.Context) error {
  u := &User{}
  if err := c.Bind(u); err != nil {
    return err
  }
```

c.Bind(u)
でgolangの構造体にcのjsonが埋め込まれる感じです。

### router

```
api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	api.GET("/todos", handler.GetTodos)
	api.POST("/todos", handler.AddTodo)
	api.DELETE("/todos/:id", handler.DeleteTodo)
	api.PUT("/todos/:id/completed", handler.UpdateTodo)
```


url := e.Group(url) とすることで、指定したURL下をグループ化することができる。
グループ化することにより、以下のように /api 下のURL(e.x. /api/todos)へのリクエスト時には必ずJWT認証を行うことを一括で指定することができる。

## Gormの使い方(How to use gorm)
http://gorm.io/ja_JP/docs/
がとてもわかりやすい。

```
package main

import (
		"fmt"
		"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {

		Id  int64 `json:"id" gorm:"column:id;primary_key"`
		Code  string `json:"code" gorm:"column:code" sql:"not null;type:varchar(200)"
		Price int8 `json:"price" gorm:"column:price" sql:"not null;type:int"
		CreatedAt  time.Tim `json:"created_at" gorm:"column:created_at" sql:"not null;type:datetime"`
		UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" sql:"not null;type:datetime"
}

func main() {

		var product Product
		var products [] Product

		// Connect
		db, err := gorm.Open("mysql", "root:secret@/go_test?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
				panic("failed to connect database")
		}
		
		// Migrate(絶対に使わない。)
		db.AutoMigrate(&Product{})


		// Create
		db.Create(&Product{Code: "test_code", Price: 1000})

		// Read
		db.First(&product, 1)
		//fmt.Println(product)

		db.First(&product, "code = ?", "test_code")
		//fmt.Println(product)

		db.Order("price desc, code").Find(&products)
		//fmt.Println(products)

		db.Select("code,price").Find(&products)
		//fmt.Println(products)

		// Update
		db.Model(&product).Update("Price", 5000)
		//fmt.Println(product)

		for _, v := range products {
				fmt.Print(v.Code)
				fmt.Print(v.Price)
		}

		//Delete
		db.Delete(&product)

}
```

gormのAutoMigrate機能は使わないでください。
まだ、開発途中で十分じゃないからです。というか、カラムが減らせないんですよ。
普通に致命傷なので、sql-migrateを使います。
以下に記述します。

## migration

### install

```
go get github.com/rubenv/sql-migrate/...
```

デフォルトと違うので注意
直下のdbconfigを移動させたら、
-config db/conf.yml

とかする。

homeのyaml多すぎでもあるので、移動させるかも。
### command

```
# ヘルプ
$ sql-migrate --help
```

```
マイグレーション作成
# 以下のコマンドでは「20181005103536-create_users.sql」みたいなマイグレーションファイルが作成される
# ファイル名のdatetimeは自動で付与

$ sql-migrate new create_users
```

```
# マイグレーションの実行
$ sql-migrate up 
```

```
# マイグレーションをdryrunで実行。実行予定のsqlが出力される
$ sql-migrate up -dryrun 
```

```
# マイグレーションのロールバック
$ sql-migrate down
```

```
# マイグレーションの実行状態確認
$ sql-migrate status
```

#### migrationfile
マイグレーション作成すると`.sql`としてファイルが生成される。sqlコメント文を使用し`up`の処理なのか`down`の処理なのかを`-- +migration Up`、`-- +migration Down`でそれぞれ指定する。

```20181005103536-create_users.sql
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (id int);

-- +migrate Down
DROP TABLE IF EXISTS users;
```

## migration table
マイグレーション初回実行時に`gorp_migrations`というテーブルが自動で作成される。
このテーブルには実行したマイグレーションファイル名をインサートして実行済みか否かを管理してる。

## docker-compose によるローカル開発環境
```
docker-compose build
docker-compose run --rm server sql-migrate up
docker-compose up
docker-compose down -v
```

### migrationファイルの作成

docker-compose run --rm sql-migrate new create_users
って感じで、内部に作ってもいいですけど、どうせマウントしてるので、
ローカルのgoにやらせる方が楽だとは思いますけどね。

## pre-commit

```
brew install pre-commit
pre-commit install
```

### データインポート

```
./cli/dev_util/run.sh db_import
```

にしたい。けど、まだできていないから、テキトーにcreate分発行しておいてください。
dumpとかすらありません。そのうち作ります。
## ディレクトリ構成(Constitution of Directory)
！！新プロジェクトではimport 文のpathを修正するのを忘れないように

- main.go
