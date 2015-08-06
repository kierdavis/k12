// 4-bit binary adder
module ttl74283(
    input [3:0] a,
    input [3:0] b,
    input carry_in,
    output [3:0] sum,
    output carry_out
);
    assign {carry_out, sum} = {1'd0, a} + {1'd0, b} + {4'd0, carry_in};
    /*
    // From Figure 4 of http://www.farnell.com/datasheets/1846165.pdf
    wire w[26:0];
    assign w[0] = ~carry_in;
    assign w[1] = ~(a[0] | b[0]);
    assign w[2] = ~(a[0] & b[0]);
    assign w[3] = ~(a[1] | b[1]);
    assign w[4] = ~(a[1] & b[1]);
    assign w[5] = ~(a[2] | b[2]);
    assign w[6] = ~(a[2] & b[2]);
    assign w[7] = ~(a[3] | b[3]);
    assign w[8] = ~(a[3] & b[3]);
    assign w[9] = ~w[0];
    assign w[10] = ~w[1] & w[2];
    assign w[11] = w[0] & w[2];
    assign w[12] = ~w[3] & w[4];
    assign w[13] = w[0] & w[2] & w[4];
    assign w[14] = w[4] & w[1];
    assign w[15] = ~w[5] & w[6];
    assign w[16] = w[0] & w[2] & w[4] & w[6];
    assign w[17] = w[4] & w[6] & w[1];
    assign w[18] = w[6] & w[3];
    assign w[19] = ~w[7] & w[8];
    assign w[20] = w[0] & w[2] & w[4] & w[6] & w[8];
    assign w[21] = w[4] & w[6] & w[8] & w[1];
    assign w[22] = w[6] & w[8] & w[3];
    assign w[23] = w[8] & w[5];
    assign w[24] = ~(w[11] | w[1]);
    assign w[25] = ~(w[13] | w[14] | w[3]);
    assign w[26] = ~(w[16] | w[17] | w[18] | w[5]);
    assign sum[0] = w[9] ^ w[10];
    assign sum[1] = w[24] ^ w[12];
    assign sum[2] = w[25] ^ w[15];
    assign sum[3] = w[26] ^ w[19];
    assign carry_out = ~(w[20] | w[21] | w[22] | w[23] | w[7]);
    */
endmodule
