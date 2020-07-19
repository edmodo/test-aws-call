DREG = registry.edmodo.io
GITHUB_ACCESS_TOKEN_FILE = .github.token
GITHUB_ACCESS_TOKEN = `cat $(GITHUB_ACCESS_TOKEN_FILE)`

all: rubyhack py27hack py37hack

base:
	# curl -sfX POST https://vaultmodo.clubmodo.com/edmodo-base-image.tar.gz --data-binary @${SECRET_FILE} -o vault_configs.tar.gz
	docker build -t $(DREG)/base -f Dockerfile.base . && \
	touch .build.base

ruby23: .build.base
	docker build -t $(DREG)/ruby23 -f Dockerfile.ruby23 . && \
	touch .build.ruby23

rubyhack: .build.ruby23
	docker build -t $(DREG)/rubyhack --build-arg GITHUB_ACCESS_TOKEN=$(GITHUB_ACCESS_TOKEN) -f Dockerfile.rubyhack . && \
	touch .build.rubyhack

py27: .build.base
	docker build -t $(DREG)/python27 -f Dockerfile.python27 . && \
	touch .build.py27

py27hack: .build.py27
	curl -H "Authorization: token ${GITHUB_ACCESS_TOKEN}" \
		-H 'Accept: application/vnd.github.v3.raw' \
		-o endpoints.json.patch \
		-L https://raw.githubusercontent.com/edmodo/aws-endpoints/master/python-boto/endpoints.json && \
	docker build -t $(DREG)/py27hack -f Dockerfile.py27hack . && \
	rm endpoints.json.patch && \
	touch .build.py27hack

py37: .build.base
	docker build -t $(DREG)/python37 -f Dockerfile.python37 . && \
	touch .build.py37

py37hack: .build.py37
	curl -H "Authorization: token ${GITHUB_ACCESS_TOKEN}" \
		-H 'Accept: application/vnd.github.v3.raw' \
		-o endpoints.json.patch \
		-L https://raw.githubusercontent.com/edmodo/aws-endpoints/master/python-boto/endpoints.json && \
	docker build -t $(DREG)/py37hack -f Dockerfile.py37hack . && \
	rm endpoints.json.patch && \
	touch .build.py37hack

clean:
	rm .build.*
