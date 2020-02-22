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


## migration

### install

```
go get github.com/rubenv/sql-migrate/...
```

db/dbconfig.yml
が設定ファイル
デフォルトと違うので注意

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
$ sql-migrate up -config db/dbconfig.yml
```

```
# マイグレーションをdryrunで実行。実行予定のsqlが出力される
$ sql-migrate up -dryrun -config db/dbconfig.yml
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



## ディレクトリ構成(Constitution of Directory)
！！新プロジェクトではimport 文のpathを修正するのを忘れないように

- main.go
