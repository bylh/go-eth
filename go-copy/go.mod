module go-copy

go 1.12

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/PuerkitoBio/purell v1.1.0
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
	github.com/Unknwon/com v0.0.0-20190321035513-0fed4efef755
	github.com/astaxie/beego v1.12.0
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-contrib/sse v0.1.0
	github.com/gin-gonic/gin v1.4.0
	github.com/go-eth/go-copy/docs v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/middleware/jwt v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/e v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/export v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/file v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/gredis v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/logging v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/qrcode v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/setting v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/upload v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/util v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/routers v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/article_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/auth_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/cache_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/tag_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-ini/ini v1.44.0
	github.com/go-openapi/jsonpointer v0.17.0
	github.com/go-openapi/jsonreference v0.19.0
	github.com/go-openapi/spec v0.19.0
	github.com/go-openapi/swag v0.17.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.3.1
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.10
	github.com/jinzhu/inflection v1.0.0
	github.com/json-iterator/go v1.1.6
	github.com/mailru/easyjson v0.0.0-20180823135443-60711f1a8329
	github.com/mattn/go-isatty v0.0.8
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
	github.com/modern-go/reflect2 v1.0.1
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.2
	github.com/tealeg/xlsx v1.0.3
	github.com/ugorji/go v1.1.5-pre
	golang.org/x/image v0.0.0-20190703141733-d6a02ce849c9
	golang.org/x/net v0.0.0-20190611141213-3f473d35a33a
	golang.org/x/text v0.3.2
	gopkg.in/go-playground/validator.v8 v8.18.2
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/go-eth/go-copy/middleware/jwt => ./middleware/jwt

replace github.com/go-eth/go-copy/docs => ./docs

replace github.com/go-eth/go-copy/models => ./models

replace github.com/go-eth/go-copy/routers => ./routers

replace github.com/go-eth/go-copy/service/article_service => ./service/article_service

replace github.com/go-eth/go-copy/service/auth_service => ./service/auth_service

replace github.com/go-eth/go-copy/service/cache_service => ./service/cache_service

replace github.com/go-eth/go-copy/service/tag_service => ./service/tag_service

replace github.com/go-eth/go-copy/pkg/app => ./pkg/app

replace github.com/go-eth/go-copy/pkg/e => ./pkg/e

replace github.com/go-eth/go-copy/pkg/export => ./pkg/export

replace github.com/go-eth/go-copy/pkg/file => ./pkg/file

replace github.com/go-eth/go-copy/pkg/gredis => ./pkg/gredis

replace github.com/go-eth/go-copy/pkg/logging => ./pkg/logging

replace github.com/go-eth/go-copy/pkg/qrcode => ./pkg/qrcode

replace github.com/go-eth/go-copy/pkg/setting => ./pkg/setting

replace github.com/go-eth/go-copy/pkg/upload => ./pkg/upload

replace github.com/go-eth/go-copy/pkg/util => ./pkg/util
