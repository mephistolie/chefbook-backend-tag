module github.com/mephistolie/chefbook-backend-tag

go 1.20

require (
	github.com/google/uuid v1.3.0
	github.com/jackc/pgx/v5 v5.3.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/mephistolie/chefbook-backend-common/log v0.6.0
	github.com/mephistolie/chefbook-backend-common/migrate/sql v0.6.0
	github.com/mephistolie/chefbook-backend-common/responses v0.9.0
	github.com/mephistolie/chefbook-backend-common/shutdown v0.6.0
	github.com/peterbourgon/ff/v3 v3.3.0
	google.golang.org/grpc v1.54.0
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b
)

require (
	github.com/golang-migrate/migrate/v4 v4.15.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.8.0 // indirect
	github.com/jackc/pgerrcode v0.0.0-20201024163028-a0d42d470451 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.6.2 // indirect
	github.com/jackc/pgx/v4 v4.10.1 // indirect
	github.com/lib/pq v1.10.0 // indirect
	github.com/sirupsen/logrus v1.9.2 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/mephistolie/chefbook-backend-tag/api => ./api
