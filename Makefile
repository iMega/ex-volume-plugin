NAME = auth
PGROUP= imega-teleport
DGROUP= imegateleport
IMG = $(DGROUP)/$(NAME)
CWD = /go/src/github.com/$(PGROUP)/$(NAME)
LINTER_FLAGS = --fast
TAG = latest


build: unit
	@docker build --build-arg CWD=$(CWD) -t $(IMG):$(TAG) .

clean:
	@-rm $(CURDIR)/mysql.log
	@TAG=$(TAG) IMG=$(IMG) docker-compose rm -sfv

.PHONY: acceptance
