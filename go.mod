module github.com/mephistolie/chefbook-backend-tag

go 1.25.0

require (
	github.com/jackc/pgx/v5 v5.9.2
	github.com/jmoiron/sqlx v1.4.0
	github.com/mephistolie/chefbook-backend-common/log v0.6.0
	github.com/mephistolie/chefbook-backend-common/migrate/sql v0.6.0
	github.com/mephistolie/chefbook-backend-common/responses v0.9.0
	github.com/mephistolie/chefbook-backend-common/shutdown v0.6.0
	github.com/mephistolie/chefbook-backend-tag/api v0.0.0-00010101000000-000000000000
	github.com/peterbourgon/ff/v3 v3.4.0
	google.golang.org/grpc v1.80.0
	k8s.io/utils v0.0.0-20260319190234-28399d86e0b5
)

require (
	github.com/golang-migrate/migrate/v4 v4.19.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgerrcode v0.0.0-20250907135507-afb5586c32a6 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.4 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/lib/pq v1.12.3 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	golang.org/x/crypto v0.50.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260427160629-7cedc36a6bc4 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/mephistolie/chefbook-backend-tag/api => ./api
