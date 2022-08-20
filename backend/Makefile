postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Password123 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root gev_example

dropdb:
	docker exec -it postgres dropdb gev_example

migrateup:
	migrate -path database/migration -database "postgresql://root:Password123@localhost:5432/gev_example?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migration -database "postgresql://root:Password123@localhost:5432/gev_example?sslmode=disable" -verbose down

generatesqlc:
	docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

userproto:
	protoc models/protobuff/m_user.proto --js_out=import_style=commonjs:./ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./ --go_out=plugins=grpc:.

.PHONY: postgres createdb dropdb migrateup migratedown generatesqlc userproto