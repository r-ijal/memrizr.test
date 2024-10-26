.PHONY: keypair migrate-create migrate-up migrate-down migrate-force

PWD = $(shell pwd)
ACCTPATH_UNIX  = $(PWD)/account
ACCTPATH = $(shell cygpath -w $(ACCTPATH_UNIX))
MPATH_UNIX = $(ACCTPATH_UNIX)/migrations
MPATH = $(shell cygpath -m $(MPATH_UNIX))
PORT = 5433
ENV ?= dev  # Default to 'dev' if ENV is not set
N = 1 # Default number of migrations to execute up or down

create-keypair:
	@echo "Current working directory: $(PWD)"
	@echo "Account path: $(ACCTPATH)"
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem


migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

migrate-up:
	@echo "ACCTPATH_UNIX $(ACCTPATH_UNIX)"
	@echo "MPATH_UNIX $(MPATH_UNIX)"
	@echo "MPATH $(MPATH) with cygpath -m"
	migrate -source file://$(MPATH) -database postgres://postgres:postgres@localhost:$(PORT)/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(MPATH) -database postgres://postgres:postgres@localhost:$(PORT)/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(MPATH) -database postgres://postgres:postgres@localhost:$(PORT)/postgres?sslmode=disable force $(VERSION)