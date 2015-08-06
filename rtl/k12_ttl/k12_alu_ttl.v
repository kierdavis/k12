module k12_alu_ttl(
    input [7:0] a,     // register A
    input [7:0] b,     // register B
    input [15:0] inst, // instruction
    output [7:0] res,  // result
    output cond        // condition
);

    wire [10:0] dc;  // submodule output ports that are not connected to anything

    wire [7:0] add;  // output of adder
    wire [7:0] addb; // output of adder input inverters
    wire [7:0] and_; // output of AND logic
    wire [7:0] bi;   // output of B/immediate multiplexer
    wire borrow;     // borrow flag (= ~carry)
    wire carry;      // carry out from adder
    wire condmux;    // output from condition multiplexer
    wire halfcarry;  // carry out of lower 4 bits / carry in to higher 4 bits
    wire inst8_n;    // = ~inst[8]
    wire n0;         // = add[7] ^ a[7]
    wire n1;         // = add[7] ^ addb[7]
    wire n2;         // = |add[1:0]
    wire n3;         // = |add[3:2]
    wire n4;         // = |add[5:4]
    wire n5;         // = |add[7:6]
    wire n6;         // = |add[3:0]
    wire n7;         // = |add[7:4]
    wire [7:0] or_;  // output of OR logic
    wire overflow_n; // = ~overflow
    wire sle;        // signed less-than-or-equal-to flag
    wire slt;        // signed less-than flag
    wire sub;        // high for subtraction, low for addition
    wire ule;        // unsigned less-than-or-equal-to flag
    wire [7:0] xor_; // output of XOR logic
    wire zero;       // zero flag

    // B/immediate multiplexer
    ttl74157 u1(
        .sel(inst[12]),
        .in0(b[7:4]),
        .in1(inst[7:4]),
        .enable_n(1'b0),
        .out(bi[7:4])
    );
    ttl74157 u2(
        .sel(inst[12]),
        .in0(b[3:0]),
        .in1(inst[3:0]),
        .enable_n(1'b0),
        .out(bi[3:0])
    );
    
    // AND, OR & XOR logic
    ttl7408 u3(
        .a(a[7:4]),
        .b(bi[7:4]),
        .q(and_[7:4])
    );
    ttl7408 u4(
        .a(a[3:0]),
        .b(bi[3:0]),
        .q(and_[3:0])
    );
    ttl7432 u5(
        .a(a[7:4]),
        .b(bi[7:4]),
        .q(or_[7:4])
    );
    ttl7432 u6(
        .a(a[3:0]),
        .b(bi[3:0]),
        .q(or_[3:0])
    );
    ttl7486 u7(
        .a(a[7:4]),
        .b(bi[7:4]),
        .q(xor_[7:4])
    );
    ttl7486 u8(
        .a(a[3:0]),
        .b(bi[3:0]),
        .q(xor_[3:0])
    );
    
    // Adder
    ttl74283 u9(
        .a(a[3:0]),
        .b(addb[3:0]),
        .carry_in(sub),
        .sum(add[3:0]),
        .carry_out(halfcarry)
    );
    ttl74283 u10(
        .a(a[7:4]),
        .b(addb[7:4]),
        .carry_in(halfcarry),
        .sum(add[7:4]),
        .carry_out(carry)
    );
    
    // Adder input inverters (invert B input when in subtraction mode).
    ttl7486 u11(
        .a({4{sub}}),
        .b(bi[7:4]),
        .q(addb[7:4])
    );
    ttl7486 u12(
        .a({4{sub}}),
        .b(bi[3:0]),
        .q(addb[3:0])
    );
    
    // Condition logic
    ttl7400 u13(
        .a({inst[8], inst[11], n0,         1'b1}),
        .b({inst[8], inst8_n,  n1,         1'b1}),
        .q({inst8_n, sub,      overflow_n, dc[0]})
    );
    ttl7432 u14(
        .a({add[0],  add[2],   add[4],     add[6]}),
        .b({add[1],  add[3],   add[5],     add[7]}),
        .q({n2,      n3,       n4,         n5})
    );
    ttl7432 u15(
        .a({n2,      zero,     zero,       n4}),
        .b({n3,      borrow,   slt,        n5}),
        .q({n6,      ule,      sle,        n7})
    );
    ttl7402 u16(
        .a({n6,      carry,    overflow_n, 1'b1}),
        .b({n7,      carry,    overflow_n, 1'b1}),
        .q({zero,    borrow,   overflow,   dc[1]})
    );
    ttl7486 u17(
        .a({add[7],  add[7],   add[7],     inst[13]}),
        .b({a[7],    addb[7],  overflow,   condmux}),
        .q({n0,      n1,       slt,        cond})
    );
    ttl74151 u18(
        .sel(inst[10:8]),
        .in0(zero),
        .in1(add[7]),
        .in2(borrow),
        .in3(overflow),
        .in4(borrow),
        .in5(ule),
        .in6(slt),
        .in7(sle),
        .enable_n(1'b0),
        .out(condmux),
        .out_n(dc[2])
    );
    
    // Output multiplexers
    ttl74151 u20(
        .sel(inst[10:8]),
        .in0(a[0]),
        .in1(and_[0]),
        .in2(or_[0]),
        .in3(xor_[0]),
        .in4(add[0]),
        .in5(add[0]),
        .in6(a[1]),
        .in7(bi[0]),
        .enable_n(1'b0),
        .out(res[0]),
        .out_n(dc[3])
    );
    ttl74151 u21(
        .sel(inst[10:8]),
        .in0(a[1]),
        .in1(and_[1]),
        .in2(or_[1]),
        .in3(xor_[1]),
        .in4(add[1]),
        .in5(add[1]),
        .in6(a[2]),
        .in7(bi[1]),
        .enable_n(1'b0),
        .out(res[1]),
        .out_n(dc[4])
    );
    ttl74151 u22(
        .sel(inst[10:8]),
        .in0(a[2]),
        .in1(and_[2]),
        .in2(or_[2]),
        .in3(xor_[2]),
        .in4(add[2]),
        .in5(add[2]),
        .in6(a[3]),
        .in7(bi[2]),
        .enable_n(1'b0),
        .out(res[2]),
        .out_n(dc[5])
    );
    ttl74151 u23(
        .sel(inst[10:8]),
        .in0(a[3]),
        .in1(and_[3]),
        .in2(or_[3]),
        .in3(xor_[3]),
        .in4(add[3]),
        .in5(add[3]),
        .in6(a[4]),
        .in7(bi[3]),
        .enable_n(1'b0),
        .out(res[3]),
        .out_n(dc[6])
    );
    ttl74151 u24(
        .sel(inst[10:8]),
        .in0(a[4]),
        .in1(and_[4]),
        .in2(or_[4]),
        .in3(xor_[4]),
        .in4(add[4]),
        .in5(add[4]),
        .in6(a[5]),
        .in7(bi[4]),
        .enable_n(1'b0),
        .out(res[4]),
        .out_n(dc[7])
    );
    ttl74151 u25(
        .sel(inst[10:8]),
        .in0(a[5]),
        .in1(and_[5]),
        .in2(or_[5]),
        .in3(xor_[5]),
        .in4(add[5]),
        .in5(add[5]),
        .in6(a[6]),
        .in7(bi[5]),
        .enable_n(1'b0),
        .out(res[5]),
        .out_n(dc[8])
    );
    ttl74151 u26(
        .sel(inst[10:8]),
        .in0(a[6]),
        .in1(and_[6]),
        .in2(or_[6]),
        .in3(xor_[6]),
        .in4(add[6]),
        .in5(add[6]),
        .in6(a[7]),
        .in7(bi[6]),
        .enable_n(1'b0),
        .out(res[6]),
        .out_n(dc[9])
    );
    ttl74151 u27(
        .sel(inst[10:8]),
        .in0(a[7]),
        .in1(and_[7]),
        .in2(or_[7]),
        .in3(xor_[7]),
        .in4(add[7]),
        .in5(add[7]),
        .in6(a[7]),
        .in7(bi[7]),
        .enable_n(1'b0),
        .out(res[7]),
        .out_n(dc[10])
    );
endmodule
