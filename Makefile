.PHONY: binary supervisor run clean tool lint help
all: binary supervisor

BUILD=./workdir
DEB=./deb
TARGET=student_server
TARGET_BIN=./bin

SERVICE=${BUILD}/service
SERVICE_BIN=${SERVICE}/pitrix
SRC_CTR=${BUILD}/control

SUPERVISOR=${BUILD}/supervisor
SUPERVISOR_BIN=${SUPERVISOR}/pitrix
SUPERVISOR_DEBIAN=${BUILD}/supervisor/DEBIAN
SUPERVISOR_CONTROL=${SUPERVISOR_DEBIAN}/control
PITRIX_QKE_APISERVER_DRI=${BUILD}/supervisor/pitrix/lib/student-server

PITRIX=${BUILD}/pitrix
PITRIX_BIN=${PITRIX_QKE_APISERVER_DRI}

GET_SIZE_SH=${BUILD}/getsize.sh
DEB_SIZE=`${GET_SIZE_SH} ${SUPERVISOR}`

binary:
	@echo "--> Building..."
	@cp -r ./build ${BUILD}
	@cp go.mod ${BUILD}
	@mkdir -p ${TARGET_BIN}
	@mkdir -p ${PITRIX}
	# GOARCH=arm64  GOOS=darwin go build -o ${TARGET_BIN}/${TARGET} main.go
	# GOARCH=amd64  GOOS=darwin go build -o ${TARGET_BIN}/${TARGET} main.go
	GOARCH=amd64  GOOS=linux go build -o ${TARGET_BIN}/${TARGET} main.go 
	# GOARCH=amd64  GOOS=windows go build -o ${TARGET_BIN}/${TARGET} main.go 
	@chmod 755 ${TARGET_BIN}/*
	@chmod 755 ${SUPERVISOR_DEBIAN}/*
	@cp -r ${BUILD}/go.mod go.mod
	
supervisor:binary
	@echo "--> packing supervisor..."
	chmod 755 ${GET_SIZE_SH}
	chmod 755 ${SUPERVISOR_DEBIAN}/*
	mkdir -p ${DEB}
	mkdir -p ${PITRIX_BIN}
	cp -rf ${TARGET_BIN}/${TARGET} ${PITRIX_BIN}
	cp -rf ${PITRIX} ${SUPERVISOR}
	cp -rf ${SRC_CTR} ${SUPERVISOR_CONTROL}
	echo "Installed-Size: ${DEB_SIZE}" >> ${SUPERVISOR_CONTROL}
	dpkg-deb -Z gzip -b ${SUPERVISOR} ${DEB}
	@cp -r ${BUILD}/go.mod go.mod
	rm -rf ${BUILD}

run:binary
	${TARGET_BIN}/${TARGET} -swagger=false

help:
	@echo "make 格式化go代码 并编译生成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make test 执行测试case"
	@echo "make check 格式化go代码"
	@echo "make cover 检查测试覆盖率"
	@echo "make run 直接运行程序"
	@echo "make lint 执行代码检查"
	@echo "make docker 构建docker镜像"
