# golang_template

## 程式結構

api => 處理request事務

config => 初始化設置

database => 處理db connection, migrate, seeder 事務

docs => swagger

model => 定義db_table結構, crud事務

service => 處理request與api中間層的事物

util => 共用func

.env => git push會被省略，供開發者本地開發使用

.env.example => 開發者第一次架設env使用

.gitignore

main.go => 系統入口

production_config.go => 打包切換gin mode機制

test_config.go  => 打包切換gin mode機制

router.go => 設置路由

<hr>

## 本地開發啟動程式
```
go run api
```
## 打包相關設定
### 根據不同的作業系統需選擇不同的環境

(Linux, windows , darwin)

GOOS=linux

(amd64, 386)

GOARCH=amd64
```
#正式
env GOOS=linux GOARCH=amd64 go build -o linux_linux
```
```
#測試
env GOOS=linux GOARCH=amd64 go build -tags test -o linux_linux
```

## 提醒

```
要記得啟用go module
go env | grep GO111MODULE
```
