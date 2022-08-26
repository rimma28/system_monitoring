.PHONY: all
all:
	docker build -t system-monitor .

.PHONY: run
run: all
	docker rm system-monitor; \
	docker run \
	--net="host" \
	--pid="host" \
	-v "/:/host:ro,rslave" \
	--name system-monitor \
	system-monitor

.PHONY: test
test: all
	docker rm system-monitor-test; \
	docker run \
	--net="host" \
	--pid="host" \
	--entrypoint="go" \
	-v "/:/host:ro,rslave" \
	--name system-monitor-test \
	system-monitor test os/linux

.PHONY: clean
clean:
	docker rm system-monitor; docker rm system-monitor-test; docker rmi system-monitor