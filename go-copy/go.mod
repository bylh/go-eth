module go-copy

go 1.12

require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-eth/go-copy/docs v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/middleware/jwt v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/models v0.0.0-00010101000000-000000000000
	github.com/go-eth/go-copy/pkg/app v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/e v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/export v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/file v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/gredis v0.0.0-00010101000000-000000000000
	github.com/go-eth/go-copy/pkg/logging v0.0.0-00010101000000-000000000000
	github.com/go-eth/go-copy/pkg/qrcode v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/setting v0.0.0-00010101000000-000000000000
	github.com/go-eth/go-copy/pkg/upload v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/pkg/util v0.0.0-00010101000000-000000000000
	github.com/go-eth/go-copy/routers v0.0.0-00010101000000-000000000000
	github.com/go-eth/go-copy/service/article_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/auth_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/cache_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-eth/go-copy/service/tag_service v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/websocket v1.4.0 // indirect
	github.com/nntaoli-project/GoEx v1.0.3
	github.com/nubo/jwt v0.0.0-20150918093313-da5b79c3bbaf // indirect
	gopkg.in/ini.v1 v1.44.0 // indirect
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
