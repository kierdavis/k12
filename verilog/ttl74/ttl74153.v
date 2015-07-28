// Dual 4:1 multiplexer
module ttl74153(
    input [1:0] sel,
    input [1:0] data0,
    input [1:0] data1,
    input [1:0] data2,
    input [1:0] data3,
    input [1:0] enable_n,
    output [1:0] out
);
    wire [1:0] selected = (sel == 2'd0) ? data0 :
                          (sel == 2'd1) ? data1 :
                          (sel == 2'd2) ? data2 :
                          (sel == 2'd3) ? data3 : 2'bxx;
    assign out = {enable_n[1] ? 1'b0 : selected[1], enable_n[0] ? 1'b0 : selected[0]};
endmodule
