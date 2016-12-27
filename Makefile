.PHONY: build
build:
	cd examples/scroll; go build; cd ../../
	cd examples/text-view; go build; cd ../../;
	cd examples/animating-text; go build; cd ../../;
	cd examples/clickables; go build; cd ../../;
	cd examples/drawing-primitives; go build; cd ../../;
	cd examples/image-tile; go build; cd ../../;
	cd examples/mouse-events; go build; cd ../../;
	cd examples/moving-image; go build; cd ../../;
	cd examples/moving-text; go build; cd ../../;
	cd examples/sine-wave; go build; cd ../../;
	cd examples/sound-check; go build; cd ../../;
	cd examples/sound-player; go build; cd ../../;
	cd examples/children; go build; cd ../../;
	cd examples/clipping; go build; cd ../../;
	cd examples/many-children; go build; cd ../../;

	
