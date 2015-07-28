// Quad 2-input NOR gate
module ttl7402(
    input [3:0] a,
    input [3:0] b,
    output [3:0] q
);
    assign q = ~(a | b);
endmodule
