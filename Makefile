export DB_HOST=127.0.0.1
export DB_USER=postgres
export DB_PASSWORD=1234
export DB_NAME=todolistulit
export DB_SSL_MODE=disable
export DB_PORT=6969

dev:
	air

debug:
	dlv debug

install_fe_deps:
	cd ./web/; npm i; cd -

tailwind_dev:
	npm run tailwind:dev --prefix ./web/

goose_status:
	goose -dir ./sql/schema postgres "host=$$DB_HOST user=$$DB_USER password=$$DB_PASSWORD dbname=$$DB_NAME port=$$DB_PORT sslmode=$$DB_SSL_MODE" status

goose_up:
	goose -dir ./sql/schema postgres "host=$$DB_HOST user=$$DB_USER password=$$DB_PASSWORD dbname=$$DB_NAME port=$$DB_PORT sslmode=$$DB_SSL_MODE" up

sqlc_generate:
	sqlc generate
