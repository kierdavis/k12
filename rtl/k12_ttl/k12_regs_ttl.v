module k12_regs_ttl(
    
);

    wire [6:0] dc;

    
    wire [7:0] a;          // Register A output
    wire [7:0] amux;       // Register A multiplexer output
    wire [7:0] b;          // Register B output
    wire [7:0] c;          // Register C output
    wire       clk;        // Clock
    wire       cke_a_n;    // Register A clock enable, active low
    wire       cke_b_n;    // Register B clock enable, active low
    wire       cke_c_n;    // Register C clock enable, active low
    wire       cke_d_n;    // Register D clock enable, active low
    wire [5:0] cke_temp_n; // Interconnect nodes for clock enable logic
    wire [7:0] cmux;       // Register C multiplexer output
    wire [7:0] d;          // Register D output
    wire [7:0] dmux;       // Register D multiplexer output
    wire [7:0] mux;        // Main multiplexer output

    // Registers
    ttl74377 u1(
        .clock(clk),
        .enable_n(cke_a_n),
        .d(amux),
        .q(a)
    );
    ttl74377 u2(
        .clock(clk),
        .enable_n(cke_b_n),
        .d(mux),
        .q(b)
    );
    ttl74377 u3(
        .clock(clk),
        .enable_n(cke_c_n),
        .d(cmux),
        .q(c)
    );
    ttl74377 u4(
        .clock(clk),
        .enable_n(cke_d_n),
        .d(dmux),
        .q(d)
    );
    
    // Main multiplexer
    ttl74153 u5(
        .sel(inst[13:12]),
        .in0(alures[7:6]),
        .in1(alures[7:6]),
        .in2(c[7:6]),
        .in3(d[7:6]),
        .enable_n(2'b00),
        .out(mux[7:6])
    );
    ttl74153 u6(
        .sel(inst[13:12]),
        .in0(alures[5:4]),
        .in1(alures[5:4]),
        .in2(c[5:4]),
        .in3(d[5:4]),
        .enable_n(2'b00),
        .out(mux[5:4])
    );
    ttl74153 u7(
        .sel(inst[13:12]),
        .in0(alures[3:2]),
        .in1(alures[3:2]),
        .in2(c[3:2]),
        .in3(d[3:2]),
        .enable_n(2'b00),
        .out(mux[3:2])
    );
    ttl74153 u8(
        .sel(inst[13:12]),
        .in0(alures[1:0]),
        .in1(alures[1:0]),
        .in2(c[1:0]),
        .in3(d[1:0]),
        .enable_n(2'b00),
        .out(mux[1:0])
    );
    
    // Register A multiplexer
    ttl74153 u9(
        .sel({do_rdio_n, do_ld_n}),
        .in0(iodata[7:6]),
        .in1(iodata[7:6]),
        .in2(memdata[7:6]),
        .in3(mux[7:6]),
        .enable_n(2'b00),
        .out(amux[7:6])
    );
    ttl74153 u10(
        .sel({do_rdio_n, do_ld_n}),
        .in0(iodata[5:4]),
        .in1(iodata[5:4]),
        .in2(memdata[5:4]),
        .in3(mux[5:4]),
        .enable_n(2'b00),
        .out(amux[5:4])
    );
    ttl74153 u11(
        .sel({do_rdio_n, do_ld_n}),
        .in0(iodata[3:2]),
        .in1(iodata[3:2]),
        .in2(memdata[3:2]),
        .in3(mux[3:2]),
        .enable_n(2'b00),
        .out(amux[3:2])
    );
    ttl74153 u12(
        .sel({do_rdio_n, do_ld_n}),
        .in0(iodata[1:0]),
        .in1(iodata[1:0]),
        .in2(memdata[1:0]),
        .in3(mux[1:0]),
        .enable_n(2'b00),
        .out(amux[1:0])
    );
    
    // Register C multiplexer
    ttl74157 u13(
        .sel(do_mov_n),
        .in0(mux[7:4]),
        .in1(idout[15:12]),
        .enable_n(1'b0),
        .out(cmux[7:4])
    );
    ttl74157 u14(
        .sel(do_mov_n),
        .in0(mux[3:0]),
        .in1(idout[11:8]),
        .enable_n(1'b0),
        .out(cmux[3:0])
    );
    
    // Register D multiplexer
    ttl74157 u15(
        .sel(do_mov_n),
        .in0(mux[7:4]),
        .in1(idout[7:4]),
        .enable_n(1'b0),
        .out(dmux[7:4])
    );
    ttl74157 u16(
        .sel(do_mov_n),
        .in0(mux[3:0]),
        .in1(idout[3:0]),
        .enable_n(1'b0),
        .out(dmux[3:0])
    );
    
    // Clock enable logic
    ttl74139 u17(
        .enable_n_a(do_mon_n),
        .sel_a(inst[15:14]),
        .out_n_a({cke_temp_n[0], cke_b_n, cke_temp_n[1], cke_temp_n[2]}),
        .enable_n_b(1'b1),
        .sel_b(2'b11),
        .out_n_b(dc[3:0])
    );
    ttl7408 u18(
        .a({do_rdio_n,     cke_temp_n[3], 1'b1,  1'b1}),
        .b({do_ld_n,       cke_temp_n[0], 1'b1,  1'b1}),
        .q({cke_temp_n[3], cke_a_n,       dc[4], dc[5]})
    );
    ttl7408 u19(
        .a({do_cdinc_n,    cke_temp_n[4], cke_temp_n[1], cke_temp_n[2]}),
        .b({do_cddec_n,    do_rcall_n,    cke_temp_n[5], cke_temp_n[5]}),
        .q({cke_temp_n[4], cke_temp_n[5], cke_c_n,       cke_d_n})
    );
    
    // Incrementer-decrementer: adder
    ttl74283 u20(
        .a(idmux[3:0]),
        .b({{3{do_cddec}}, 1'b1}),
        .carry_in(1'b0),
        .sum(idout[3:0]),
        .carry_out(idcarry0)
    );
    ttl74283 u21(
        .a(idmux[7:4]),
        .b({4{do_cddec}}),
        .carry_in(idcarry0),
        .sum(idout[7:4]),
        .carry_out(idcarry1)
    );
    ttl74283 u22(
        .a(idmux[11:8]),
        .b({4{do_cddec}}),
        .carry_in(idcarry1),
        .sum(idout[11:8]),
        .carry_out(idcarry2)
    );
    ttl74283 u23(
        .a(idmux[15:12]),
        .b({4{do_cddec}}),
        .carry_in(idcarry2),
        .sum(idout[15:12]),
        .carry_out(dc[6])
    );
    
    // Incrementer-decrementer: input multiplexer
    ttl74157 u24(
        .sel(do_rcall_n),
        .in0(pc[15:12]),
        .in1(c[7:4]),
        .enable_n(1'b0),
        .out(idmux[15:12])
    );
    ttl74157 u25(
        .sel(do_rcall_n),
        .in0(pc[11:8]),
        .in1(c[3:0]),
        .enable_n(1'b0),
        .out(idmux[11:8])
    );
    ttl74157 u26(
        .sel(do_rcall_n),
        .in0(pc[7:4]),
        .in1(d[7:4]),
        .enable_n(1'b0),
        .out(idmux[7:4])
    );
    ttl74157 u27(
        .sel(do_rcall_n),
        .in0(pc[3:0]),
        .in1(d[3:0]),
        .enable_n(1'b0),
        .out(idmux[3:0])
    );
endmodule
