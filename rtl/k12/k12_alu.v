module k12_alu(
    input [7:0] a,
    input [7:0] b,
    input [15:0] inst,
    output [7:0] res,
    output cond
);

    wire [7:0] bi = inst[12] ? inst[7:0] : b;
    
    wire subtract = ~inst[11] | inst[8];
    wire [7:0] adder_a = a;
    wire [7:0] adder_b = subtract ? ~bi : bi;
    wire [7:0] adder_result;
    wire carry;
    assign {carry, adder_result} = {1'd0, adder_a} + {1'd0, adder_b} + {8'd0, subtract};
    
    wire [2:0] func = inst[10:8];
    
    assign res = (func == 3'h0) ? a :
                 (func == 3'h1) ? (a & bi) :
                 (func == 3'h2) ? (a | bi) :
                 (func == 3'h3) ? (a ^ bi) :
                 (func == 3'h4) ? adder_result :
                 (func == 3'h5) ? adder_result :
                 (func == 3'h6) ? {a[7], a[7:1]} :
                 (func == 3'h7) ? bi : 8'hxx;
    
    wire zero = adder_result == 8'd0;
    wire negative = adder_result[7];
    wire borrow = ~carry;
    wire overflow = (adder_a[7] ^ adder_result[7]) & (adder_b[7] ^ adder_result[7]);
    wire ule = borrow | zero;
    wire slt = negative ^ overflow;
    wire sle = slt | zero;
    
    wire raw_cond = (func == 3'h0) ? zero :
                    (func == 3'h1) ? negative :
                    (func == 3'h2) ? borrow :
                    (func == 3'h3) ? overflow :
                    (func == 3'h4) ? borrow :
                    (func == 3'h5) ? ule :
                    (func == 3'h6) ? slt :
                    (func == 3'h7) ? sle : 1'hx;
    
    assign cond = inst[13] ? ~raw_cond : raw_cond;
    
endmodule
