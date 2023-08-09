lint: 
	docker run --volume "$(shell pwd):/workspace" --workdir /workspace yoheimuta/protolint lint proto -plugin ./example/plugin_example

lintb:
	protolint -plugin ./example/plugin_example ./proto
build-sample:
	cd example && go build -o plugin_example .
