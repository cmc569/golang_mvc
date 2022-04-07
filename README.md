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