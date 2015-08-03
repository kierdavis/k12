// Dual 2:4 decoder
module ttl74139(
    input enable_n_a,
    input [1:0] sel_a,
    output [0:3] out_n_a,
    input enable_n_b,
    input [1:0] sel_b,
    output [0:3] out_n_b
);
    assign decoded_a = (sel_a == 2'd0) ? 4'b0111 :
                       (sel_a == 2'd1) ? 4'b1011 :
                       (sel_a == 2'd2) ? 4'b1101 :
                       (sel_a == 2'd3) ? 4'b1110 : 4'bxxxx;
    assign out_n_a = enable_n_a ? 4'b1111 : decoded_a;
    
    assign decoded_b = (sel_b == 2'd0) ? 4'b0111 :
                       (sel_b == 2'd1) ? 4'b1011 :
                       (sel_b == 2'd2) ? 4'b1101 :
                       (sel_b == 2'd3) ? 4'b1110 : 4'bxxxx;
    assign out_n_b = enable_n_b ? 4'b1111 : decoded_b;
endmodule
