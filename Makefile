build:
	tinygo build -target wasi -o deloudify.wasm main.go

build-test:
	tinygo build -o test-deloudify.wasm --target wasi test/test.go

test-plugin:
	xtp plugin test deloudify.wasm --with test-deloudify.wasm

#test:
#	extism call deloudify.wasm api_before_post_post --input "THIS is a TEST message with SOME ALL CAPS words to CHECK the FUNCTIONALITY." --wasi
