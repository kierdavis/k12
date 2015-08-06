module k12_alu(
    input [7:0] a,
    input [7:0] b,
    input [15:0] inst,
    output [7:0] res,
    output cond
);

    wire [7:0] bi = inst[12] ? inst[7:0] : b;
    
    wire [2:0] func = inst[10:8];
    assign res = (func == 3'h0) ? a :
                 (func == 3'h1) ? (a & bi) :
                 (func == 3'h2) ? (a | bi) :
                 (func == 3'h3) ? (a ^ bi) :
                 (func == 3'h4) ? (a + bi) :
                 (func == 3'h5) ? (a - bi) :
                 (func == 3'h6) ? {a[7], a[7:1]} :
                 (func == 3'h7) ? bi : 8'hxx;
    
    wire [7:0] cmp_b = ~bi;
    wire [8:0] cmp_res = {1'd0, a} + {1'd0, cmp_b} + 9'd1;
    
    wire zero = cmp_res[7:0] == 8'd0;
    wire negative = cmp_res[7];
    wire borrow = ~cmp_res[8];
    wire overflow = (a[7] ^ cmp_res[7]) & (cmp_b[7] ^ cmp_res[7]);
    wire ule = borrow | zero;
    wire slt = negative ^ overflow;
    wire sle = slt | zero;
    
    assign cond = (func == 3'h0) ? zero :
                  (func == 3'h1) ? negative :
                  (func == 3'h2) ? borrow :
                  (func == 3'h3) ? overflow :
                  (func == 3'h4) ? borrow :
                  (func == 3'h5) ? ule :
                  (func == 3'h6) ? slt :
                  (func == 3'h7) ? sle : 1'hx;
    
endmodule
