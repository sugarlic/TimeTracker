all: build_people_info
	

build_time_tr:
	cd time_tracker && go run ./cmd/web

build_people_info:
	cd people_info && go run ./cmd/web