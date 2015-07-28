// Quad 2:1 multiplexer
module ttl74157(
    input sel,
    input [3:0] data0,
    input [3:0] data1,
    input enable_n,
    output [3:0] out
);
    wire [3:0] selected = (sel == 1'd0) ? data0 :
                          (sel == 1'd1) ? data1 : 4'bxxxx;
    assign out = enable_n ? 4'b0000 : selected;
endmodule
