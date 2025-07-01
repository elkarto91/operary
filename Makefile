# Root Makefile to wrap common Operary commands

.PHONY: run build test compose-up compose-down ui clean

run:
	$(MAKE) -C backend run

build:
	$(MAKE) -C backend build

clean:
	$(MAKE) -C backend clean

test:
	$(MAKE) -C backend test

compose-up:
	docker-compose -f backend/docker-compose.yml up --build

compose-down:
	docker-compose -f backend/docker-compose.yml down

ui:
	cd ui && npm run dev

