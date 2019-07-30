module go-eth

go 1.12

require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-eth/docs v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/middleware/jwt v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/models v0.0.0-00010101000000-000000000000
	github.com/go-eth/pkg/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/pkg/e v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/pkg/export v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/pkg/file v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/pkg/gredis v0.0.0-00010101000000-000000000000
	github.com/go-eth/pkg/logging v0.0.0-00010101000000-000000000000
	github.com/go-eth/pkg/qrcode v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/pkg/setting v0.0.0-00010101000000-000000000000
	github.com/go-eth/pkg/upload v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/pkg/util v0.0.0-00010101000000-000000000000
	github.com/go-eth/routers v0.0.0-00010101000000-000000000000
	github.com/go-eth/service/article_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/service/auth_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/service/cache_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/service/tag_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/nntaoli-project/GoEx v1.0.3
	github.com/nubo/jwt v0.0.0-20150918093313-da5b79c3bbaf // indirect
	gopkg.in/ini.v1 v1.44.0 // indirect
)

replace github.com/go-eth/middleware/jwt => ./middleware/jwt

replace github.com/go-eth/docs => ./docs

replace github.com/go-eth/models => ./models

replace github.com/go-eth/routers => ./routers

replace github.com/go-eth/service/article_service => ./service/article_service

replace github.com/go-eth/service/auth_service => ./service/auth_service

replace github.com/go-eth/service/cache_service => ./service/cache_service

replace github.com/go-eth/service/tag_service => ./service/tag_service

replace github.com/go-eth/pkg/app => ./pkg/app

replace github.com/go-eth/pkg/e => ./pkg/e

replace github.com/go-eth/pkg/export => ./pkg/export

replace github.com/go-eth/pkg/file => ./pkg/file

replace github.com/go-eth/pkg/gredis => ./pkg/gredis

replace github.com/go-eth/pkg/logging => ./pkg/logging

replace github.com/go-eth/pkg/qrcode => ./pkg/qrcode

replace github.com/go-eth/pkg/setting => ./pkg/setting

replace github.com/go-eth/pkg/upload => ./pkg/upload

replace github.com/go-eth/pkg/util => ./pkg/util
