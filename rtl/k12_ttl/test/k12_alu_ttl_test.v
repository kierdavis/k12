module k12_alu_ttl_test;
    reg [14:0] counter = 0;
    always #1 counter = counter + 1;
    
    wire [7:0] a = {counter[2], {6{counter[1]}}, counter[0]};
    wire [7:0] b = {counter[5], {6{counter[4]}}, counter[3]};
    wire [15:0] inst = {2'b00, counter[14:9], {counter[8], {6{counter[7]}}, counter[6]}};
    wire [7:0] res;
    wire cond;
    
    k12_alu_ttl u(
        .a(a),
        .b(b),
        .inst(inst),
        .res(res),
        .cond(cond)
    );
    
    initial begin
        $display("a,b,inst,res,cond");
        $monitor("%x,%x,%x,%x,%x", a, b, inst, res, cond);
        #32767
        $finish;
    end
endmodule
