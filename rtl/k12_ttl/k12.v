module k12(
);

    wire [7:0] a;
    wire [7:0] b;
    wire [15:0] inst;
    wire [7:0] res;
    wire cond;
    
    k12_alu alu(
        .a(a),
        .b(b),
        .inst(inst),
        .res(res),
        .cond(cond)
    );
endmodule
