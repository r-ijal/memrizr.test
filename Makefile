.PHONY: create-keypair

PWD = $(shell pwd)
ACCTPATH_UNIX  = $(PWD)/account
ACCTPATH = $(shell cygpath -w $(ACCTPATH_UNIX))
ENV ?= dev  # Default to 'dev' if ENV is not set

create-keypair:
	@echo "Current working directory: $(PWD)"
	@echo "Account path: $(ACCTPATH)"
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

