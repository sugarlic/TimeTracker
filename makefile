# Определите переменные для удобства
DATABASE_URL=postgres://postgres:1@localhost:5432/postgres?sslmode=disable
MIGRATIONS_PATH=time_tracker/migrations

# Главная цель
all: build_time_tr

# Цель для сборки time_tracker
build_time_tr:
	cd time_tracker && go run ./cmd/web

# Цель для сборки people_info, которая зависит от migrate_down
build_people_info:
	cd people_info && go run ./cmd/web

# Цель для выполнения миграции вниз
migrate_down:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) down

migrate_up:
	migrate -database $(DATABASE_URL) -path $(MIGRATIONS_PATH) up


#  migrate -database postgres://postgres:1@localhost:5432/postgres?sslmode=disable -path people_info/migrations force 1