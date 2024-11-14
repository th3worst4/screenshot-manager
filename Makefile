build:
	cd cmd; \
	go build -o ../build/screenshot-manager main.go fileroutines.go GUI.go

run:
	cd cmd; \
	go build -o ../build/screenshot-manager main.go fileroutines.go GUI.go; \
	cd ../build; \
	./screenshot-manager

test:
	make build; \
	cd build; \
	./screenshot-manager ../assets/hi
