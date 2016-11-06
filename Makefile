.PHONY: build
build:
	echo "Building examples..."
	go build -o examples/animating-text/animating-text examples/animating-text/main.go
	go build -o examples/clickables/clickables examples/clickables/main.go
	go build -o examples/drawing-primitives/drawing-primitives examples/drawing-primitives/main.go
	go build -o examples/image-tile/image-tile examples/image-tile/main.go
	go build -o examples/mouse-events/mouse-events examples/mouse-events/main.go
	go build -o examples/moving-image/moving-image examples/moving-image/main.go
	go build -o examples/moving-text/moving-text examples/moving-text/main.go
	go build -o examples/sine-wave/sine-wave examples/sine-wave/main.go
	go build -o examples/sound-check/sound-check examples/sound-check/main.go
	go build -o examples/sound-player/sound-player examples/sound-player/main.go