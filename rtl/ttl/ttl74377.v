// Octal positive edge-triggered D flip-flop with data enable
module ttl74377(
    input clock,
    input enable_n,
    input [7:0] d,
    output reg [7:0] q
);
    always @(posedge clock)
        if (~enable_n)
            q <= d;
endmodule
