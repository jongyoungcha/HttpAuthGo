WORK_DIR:=$(shell pwd)


server:
	echo $(WORK_DIR) \
	cd $(WORK_DIR)/server && \
	go build -o $(WORK_DIR)/build/server


