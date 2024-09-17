db_username = dev
db_password = dev
db_host = localhost
db_port = 5432
db_name = orion_pay

db-start:
	docker run -dt --name ${db_name} -p ${db_port}:${db_port} -e POSTGRES_USER=${db_username} -e POSTGRES_PASSWORD=${db_password} -e POSTGRES_DB=${db_name} -e TZ=Asia/Jakarta -e PGTZ=Asia/Jakarta postgres:alpine -c 'max_connections=500'

migrate-diff:
	./atlas migrate --env gorm diff

migrate-apply:
	./atlas migrate --env gorm apply

test:
	go test -v cover ./...