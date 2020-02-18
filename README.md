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

## ディレクトリ構成(Constitution of Directory)
！！新プロジェクトではimport 文のpathを修正するのを忘れないように

