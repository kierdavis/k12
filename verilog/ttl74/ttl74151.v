// 8:1 multiplexer
module ttl74151(
    input [2:0] sel,
    input data0,
    input data1,
    input data2,
    input data3,
    input data4,
    input data5,
    input data6,
    input data7,
    input enable_n,
    output out,
    output out_n
);
    wire selected = (sel == 3'd0) ? data0 :
                    (sel == 3'd1) ? data1 :
                    (sel == 3'd2) ? data2 :
                    (sel == 3'd3) ? data3 :
                    (sel == 3'd4) ? data4 :
                    (sel == 3'd5) ? data5 :
                    (sel == 3'd6) ? data6 :
                    (sel == 3'd7) ? data7 : 1'bx;
    assign out = enable_n ? 1'b0 : selected;
    assign out_n = ~out;
endmodule
