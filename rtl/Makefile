# This file was automatically generated by build-makefile.sh
SOURCES = $(shell find . -type f -name '*.v')
all: alu alu_ttl
alu: build/k12_alu_test.csv
build/k12_alu_test.csv: build/k12_alu_test.vvp
	build/k12_alu_test.vvp > build/k12_alu_test.csv
build/k12_alu_test.vvp: ${SOURCES}
	mkdir -p build
	iverilog -o build/k12_alu_test.vvp -s k12_alu_test ${SOURCES}
alu_ttl: build/k12_alu_ttl_test.csv
build/k12_alu_ttl_test.csv: build/k12_alu_ttl_test.vvp
	build/k12_alu_ttl_test.vvp > build/k12_alu_ttl_test.csv
build/k12_alu_ttl_test.vvp: ${SOURCES}
	mkdir -p build
	iverilog -o build/k12_alu_ttl_test.vvp -s k12_alu_ttl_test ${SOURCES}
clean:
	rm -rf build
.PHONY: all alu alu_ttl clean
