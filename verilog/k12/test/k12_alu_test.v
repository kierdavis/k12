module k12_alu_test;
    wire [7:0] numbers [0:7];
    assign numbers[0] = 8'h00;
    assign numbers[1] = 8'h01;
    assign numbers[2] = 8'h7E;
    assign numbers[3] = 8'h7F;
    assign numbers[4] = 8'h80;
    assign numbers[5] = 8'h81;
    assign numbers[6] = 8'hFE;
    assign numbers[7] = 8'hFF;
    
    reg [14:0] counter = 0;
    always #1 counter = counter + 1;
    
    wire [7:0] a = numbers[counter[2:0]];
    wire [7:0] b = numbers[counter[5:3]];
    wire [15:0] inst = {2'b00, counter[14:9], numbers[counter[8:6]]};
    wire [7:0] res;
    wire cond;
    
    k12_alu u(
        .a(a),
        .b(b),
        .inst(inst),
        .res(res),
        .cond(cond)
    );
    
    initial begin
        $display("a,b,inst,res,cond");
        $monitor("%x,%x,%x,%x,%x", a, b, inst, res, cond);
        #10
        $finish;
    end
endmodule
