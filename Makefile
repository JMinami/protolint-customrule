lint: 
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace yoheimuta/protolint lint -plugin ./example/plugin_example proto 

lintb:
	protolint -plugin ./example/plugin_example ./proto
build-sample:
	cd example && go build -o plugin_example .
