VERILOG_DIR = ../..
VERILOG_SOURCES = $(shell find ${VERILOG_DIR} -type f -name '*.v')
all: alu
k12_alu_test.vvp: $(VERILOG_SOURCES)
	iverilog -o k12_alu_test.vvp -s k12_alu_test ${VERILOG_SOURCES}
k12_alu_test.csv.gz: k12_alu_test.vvp
	./k12_alu_test.vvp | gzip > k12_alu_test.csv.gz
alu: k12_alu_test.csv.gz
clean:
  rm -f k12_alu_test.vvp k12_alu_test.csv.gz
