import os, re

tb_file_regex = re.compile("^k12_([a-z_]+)_test\\.v$")

targets = []
for file in os.listdir():
    m = tb_file_regex.match(file)
    if m is not None:
        targets.append(m.group(1))

print("VERILOG_DIR = ../..")
print("VERILOG_SOURCES = $(shell find ${VERILOG_DIR} -type f -name '*.v')")
print("all: %s" % " ".join(targets))

build_files = []

for target in targets:
    module = "k12_%s_test" % target
    vvp = "%s.vvp" % module
    csvgz = "%s.csv.gz" % module
    
    print("%s: $(VERILOG_SOURCES)" % (vvp))
    print("\tiverilog -o %s -s %s ${VERILOG_SOURCES}" % (vvp, module))
    
    print("%s: %s" % (csvgz, vvp))
    print("\t./%s | gzip > %s" % (vvp, csvgz))
    
    print("%s: %s" % (target, csvgz))
    
    build_files.extend([vvp, csvgz])

print("clean:")
print("\trm -f %s" % " ".join(build_files))

#print("all: alu")

#print("alu: $(VERILOG_SOURCES)")
#        iverilog -o k12_alu_test.vvp -s k12_alu_test ${VERILOG_SOURCES}
#        ./k12_alu_test.vvp | gzip > k12_alu_test.csv.gz

#clean:
#        rm *.vvp *.csv.gz

