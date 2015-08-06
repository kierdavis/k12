// Dual 4:1 multiplexer
module ttl74153(
    input [1:0] sel,
    input [1:0] in0,
    input [1:0] in1,
    input [1:0] in2,
    input [1:0] in3,
    input [1:0] enable_n,
    output [1:0] out
);
    wire [1:0] selected = (sel == 2'd0) ? in0 :
                          (sel == 2'd1) ? in1 :
                          (sel == 2'd2) ? in2 :
                          (sel == 2'd3) ? in3 : 2'bxx;
    assign out = {enable_n[1] ? 1'b0 : selected[1], enable_n[0] ? 1'b0 : selected[0]};
endmodule
