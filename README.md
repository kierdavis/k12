# K12

K12 is a project to design & build a working 8-bit CPU out of 7400-series TTL
logic gates.

Quick links:

* [Architecture][arch]

[arch]: https://github.com/kierdavis/k12/wiki/Architecture

## Current status

As of 5th August 2015, I am currently:

* working on a test setup (probably consisting of Arduinos & my FPGA dev board) that can be used to test the ALU board prototype
* implementing two functionally identical models of the CPU in Verilog (one using normal Verilog and one using submodules representing TTL ICs) that can be used for simulation & testing
* ordering IC sockets, perfboards and LEDs (lots of them!), ready to begin construction of the computer

Things that have been done so far:

* design the architecture of the CPU
* draw out schematics for most of the CPU (ALU board, register board, control board, parts of the other boards)
* obtain some ICs & breadboards
* build a prototype of the adder section of the ALU board on breadboards
