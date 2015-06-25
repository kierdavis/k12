v 20130925 2
C 40000 40000 0 0 0 title-B.sym
C 43700 41000 1 0 0 74377-1.sym
{
T 44000 44450 5 10 0 0 0 0 1
device=74377
T 44000 44650 5 10 0 0 0 0 1
footprint=DIP20
T 43700 41000 5 10 0 0 0 0 1
x-appdesc=Output port 1 (LCD data) register
T 45400 44300 5 10 1 1 0 6 1
refdes=U?
}
C 42100 41400 1 0 0 input-2.sym
{
T 42100 41600 5 10 0 0 0 0 1
net=-DO_WRIO_1:1
T 42700 42100 5 10 0 0 0 0 1
device=none
T 42600 41500 5 10 1 1 0 7 1
value=\_DO_WRIO_1\_
}
C 42100 41100 1 0 0 input-2.sym
{
T 42100 41300 5 10 0 0 0 0 1
net=CLK:1
T 42700 41800 5 10 0 0 0 0 1
device=none
T 42600 41200 5 10 1 1 0 7 1
value=CLK
}
C 45900 43800 1 0 0 output-2.sym
{
T 46800 44000 5 10 0 0 0 0 1
net=LCD_DB0:1
T 46100 44500 5 10 0 0 0 0 1
device=none
T 46800 43900 5 10 1 1 0 1 1
value=LCD_DB0
}
C 45900 43500 1 0 0 output-2.sym
{
T 46800 43700 5 10 0 0 0 0 1
net=LCD_DB1:1
T 46100 44200 5 10 0 0 0 0 1
device=none
T 46800 43600 5 10 1 1 0 1 1
value=LCD_DB1
}
C 45900 43200 1 0 0 output-2.sym
{
T 46800 43400 5 10 0 0 0 0 1
net=LCD_DB2:1
T 46100 43900 5 10 0 0 0 0 1
device=none
T 46800 43300 5 10 1 1 0 1 1
value=LCD_DB2
}
C 45900 42900 1 0 0 output-2.sym
{
T 46800 43100 5 10 0 0 0 0 1
net=LCD_DB3:1
T 46100 43600 5 10 0 0 0 0 1
device=none
T 46800 43000 5 10 1 1 0 1 1
value=LCD_DB3
}
C 45900 42600 1 0 0 output-2.sym
{
T 46800 42800 5 10 0 0 0 0 1
net=LCD_DB4:1
T 46100 43300 5 10 0 0 0 0 1
device=none
T 46800 42700 5 10 1 1 0 1 1
value=LCD_DB4
}
C 45900 42300 1 0 0 output-2.sym
{
T 46800 42500 5 10 0 0 0 0 1
net=LCD_DB5:1
T 46100 43000 5 10 0 0 0 0 1
device=none
T 46800 42400 5 10 1 1 0 1 1
value=LCD_DB5
}
C 45900 42000 1 0 0 output-2.sym
{
T 46800 42200 5 10 0 0 0 0 1
net=LCD_DB6:1
T 46100 42700 5 10 0 0 0 0 1
device=none
T 46800 42100 5 10 1 1 0 1 1
value=LCD_DB6
}
C 45900 41700 1 0 0 output-2.sym
{
T 46800 41900 5 10 0 0 0 0 1
net=LCD_DB7:1
T 46100 42400 5 10 0 0 0 0 1
device=none
T 46800 41800 5 10 1 1 0 1 1
value=LCD_DB7
}
N 43500 43900 43700 43900 4
N 43500 43600 43700 43600 4
N 43500 43300 43700 43300 4
N 43500 42700 43700 42700 4
N 43700 43000 43500 43000 4
N 43500 42400 43700 42400 4
N 43500 42100 43700 42100 4
N 43500 41800 43700 41800 4
N 43500 41500 43700 41500 4
N 43500 41200 43700 41200 4
N 45700 43900 45900 43900 4
N 45700 43600 45900 43600 4
N 45700 43300 45900 43300 4
N 45700 43000 45900 43000 4
N 45700 42700 45900 42700 4
N 45700 42400 45900 42400 4
N 45700 42100 45900 42100 4
N 45700 41800 45900 41800 4
C 43500 44000 1 180 0 io-1.sym
{
T 42600 43800 5 10 0 0 180 0 1
net=IODATA0:1
T 43300 43400 5 10 0 0 180 0 1
device=none
T 42600 43900 5 10 1 1 180 1 1
value=IODATA0
}
C 43500 43700 1 180 0 io-1.sym
{
T 42600 43500 5 10 0 0 180 0 1
net=IODATA1:1
T 43300 43100 5 10 0 0 180 0 1
device=none
T 42600 43600 5 10 1 1 180 1 1
value=IODATA1
}
C 43500 43400 1 180 0 io-1.sym
{
T 42600 43200 5 10 0 0 180 0 1
net=IODATA2:1
T 43300 42800 5 10 0 0 180 0 1
device=none
T 42600 43300 5 10 1 1 180 1 1
value=IODATA2
}
C 43500 43100 1 180 0 io-1.sym
{
T 42600 42900 5 10 0 0 180 0 1
net=IODATA3:1
T 43300 42500 5 10 0 0 180 0 1
device=none
T 42600 43000 5 10 1 1 180 1 1
value=IODATA3
}
C 43500 42800 1 180 0 io-1.sym
{
T 42600 42600 5 10 0 0 180 0 1
net=IODATA4:1
T 43300 42200 5 10 0 0 180 0 1
device=none
T 42600 42700 5 10 1 1 180 1 1
value=IODATA4
}
C 43500 42500 1 180 0 io-1.sym
{
T 42600 42300 5 10 0 0 180 0 1
net=IODATA5:1
T 43300 41900 5 10 0 0 180 0 1
device=none
T 42600 42400 5 10 1 1 180 1 1
value=IODATA5
}
C 43500 42200 1 180 0 io-1.sym
{
T 42600 42000 5 10 0 0 180 0 1
net=IODATA6:1
T 43300 41600 5 10 0 0 180 0 1
device=none
T 42600 42100 5 10 1 1 180 1 1
value=IODATA6
}
C 43500 41900 1 180 0 io-1.sym
{
T 42600 41700 5 10 0 0 180 0 1
net=IODATA7:1
T 43300 41300 5 10 0 0 180 0 1
device=none
T 42600 41800 5 10 1 1 180 1 1
value=IODATA7
}
C 43300 46800 1 0 0 switch-dip8-1.sym
{
T 44700 49375 5 8 0 0 0 0 1
device=SWITCH_DIP8
T 43600 49550 5 10 1 1 0 0 1
refdes=SW?
T 43300 46800 5 10 0 0 0 0 1
x-appdesc=DIP switches
}
T 50000 40700 9 10 1 0 0 0 1
IO port 1 (LCD data / DIP switches)
T 49900 40400 9 10 1 0 0 0 1
port1.sch
C 50000 46600 1 0 0 74244-1.sym
{
T 50300 49750 5 10 0 0 0 0 1
device=74244
T 51700 49600 5 10 1 1 0 6 1
refdes=U?
T 50300 49950 5 10 0 0 0 0 1
footprint=DIP20
T 50000 46600 5 10 0 0 0 0 1
x-appdesc=Input port 1 (DIP switches) driver
}
C 44900 47000 1 270 0 resistor-1.sym
{
T 45300 46700 5 10 0 0 270 0 1
device=RESISTOR
T 45200 46800 5 10 1 1 270 0 1
refdes=R?
T 44900 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 7)
}
C 45400 47000 1 270 0 resistor-1.sym
{
T 45800 46700 5 10 0 0 270 0 1
device=RESISTOR
T 45700 46800 5 10 1 1 270 0 1
refdes=R?
T 45400 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 6)
}
C 45900 47000 1 270 0 resistor-1.sym
{
T 46300 46700 5 10 0 0 270 0 1
device=RESISTOR
T 46200 46800 5 10 1 1 270 0 1
refdes=R?
T 45900 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 5)
}
C 46400 47000 1 270 0 resistor-1.sym
{
T 46800 46700 5 10 0 0 270 0 1
device=RESISTOR
T 46700 46800 5 10 1 1 270 0 1
refdes=R?
T 46400 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 4)
}
C 46900 47000 1 270 0 resistor-1.sym
{
T 47300 46700 5 10 0 0 270 0 1
device=RESISTOR
T 47200 46800 5 10 1 1 270 0 1
refdes=R?
T 46900 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 3)
}
C 47400 47000 1 270 0 resistor-1.sym
{
T 47800 46700 5 10 0 0 270 0 1
device=RESISTOR
T 47700 46800 5 10 1 1 270 0 1
refdes=R?
T 47400 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 2)
}
C 47900 47000 1 270 0 resistor-1.sym
{
T 48300 46700 5 10 0 0 270 0 1
device=RESISTOR
T 48200 46800 5 10 1 1 270 0 1
refdes=R?
T 47900 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 1)
}
C 48400 47000 1 270 0 resistor-1.sym
{
T 48800 46700 5 10 0 0 270 0 1
device=RESISTOR
T 48700 46800 5 10 1 1 270 0 1
refdes=R?
T 48400 47000 5 10 0 0 0 0 1
x-appdesc=DIP switch pull-down resistor (bit 0)
}
N 44600 49300 50000 49300 4
{
T 49000 49300 5 10 1 1 0 0 1
netname=DIPSW7
}
N 45000 49300 45000 47000 4
N 44600 49000 50000 49000 4
{
T 49000 49000 5 10 1 1 0 0 1
netname=DIPSW6
}
N 45500 49000 45500 47000 4
N 44600 48700 50000 48700 4
{
T 49000 48700 5 10 1 1 0 0 1
netname=DIPSW5
}
N 46000 48700 46000 47000 4
N 44600 48400 50000 48400 4
{
T 49000 48400 5 10 1 1 0 0 1
netname=DIPSW4
}
N 46500 48400 46500 47000 4
N 44600 48100 50000 48100 4
{
T 49000 48100 5 10 1 1 0 0 1
netname=DIPSW3
}
N 47000 48100 47000 47000 4
N 44600 47800 50000 47800 4
{
T 49000 47800 5 10 1 1 0 0 1
netname=DIPSW2
}
N 47500 47800 47500 47000 4
N 44600 47500 50000 47500 4
{
T 49000 47500 5 10 1 1 0 0 1
netname=DIPSW1
}
N 48000 47500 48000 47000 4
N 44600 47200 50000 47200 4
{
T 49000 47200 5 10 1 1 0 0 1
netname=DIPSW0
}
N 48500 47200 48500 47000 4
C 48400 45400 1 0 0 gnd-1.sym
C 42900 49500 1 0 0 vcc-1.sym
N 43300 49300 43100 49300 4
N 43100 47200 43100 49500 4
N 43300 49000 43100 49000 4
N 43300 48700 43100 48700 4
N 43300 48400 43100 48400 4
N 43300 48100 43100 48100 4
N 43300 47800 43100 47800 4
N 43300 47500 43100 47500 4
N 43300 47200 43100 47200 4
N 48500 46100 48500 45700 4
N 48000 46100 48000 45900 4
N 45000 45900 48500 45900 4
N 47500 46100 47500 45900 4
N 47000 46100 47000 45900 4
N 46500 46100 46500 45900 4
N 46000 46100 46000 45900 4
N 45500 46100 45500 45900 4
N 45000 46100 45000 45900 4
N 52000 49300 52200 49300 4
N 52000 49000 52200 49000 4
N 52000 48700 52200 48700 4
N 52000 48400 52200 48400 4
N 52000 48100 52200 48100 4
N 52000 47800 52200 47800 4
N 52000 47500 52200 47500 4
N 52000 47200 52200 47200 4
N 50600 46600 50600 46400 4
N 50600 46400 51400 46400 4
N 51400 45900 51400 46600 4
C 52200 47100 1 0 0 io-1.sym
{
T 53100 47300 5 10 0 0 0 0 1
net=IODATA0:1
T 52400 47700 5 10 0 0 0 0 1
device=none
T 53100 47200 5 10 1 1 0 1 1
value=IODATA0
}
C 52200 47400 1 0 0 io-1.sym
{
T 53100 47600 5 10 0 0 0 0 1
net=IODATA1:1
T 52400 48000 5 10 0 0 0 0 1
device=none
T 53100 47500 5 10 1 1 0 1 1
value=IODATA1
}
C 52200 47700 1 0 0 io-1.sym
{
T 53100 47900 5 10 0 0 0 0 1
net=IODATA2:1
T 52400 48300 5 10 0 0 0 0 1
device=none
T 53100 47800 5 10 1 1 0 1 1
value=IODATA2
}
C 52200 48000 1 0 0 io-1.sym
{
T 53100 48200 5 10 0 0 0 0 1
net=IODATA3:1
T 52400 48600 5 10 0 0 0 0 1
device=none
T 53100 48100 5 10 1 1 0 1 1
value=IODATA3
}
C 52200 48300 1 0 0 io-1.sym
{
T 53100 48500 5 10 0 0 0 0 1
net=IODATA4:1
T 52400 48900 5 10 0 0 0 0 1
device=none
T 53100 48400 5 10 1 1 0 1 1
value=IODATA4
}
C 52200 48600 1 0 0 io-1.sym
{
T 53100 48800 5 10 0 0 0 0 1
net=IODATA5:1
T 52400 49200 5 10 0 0 0 0 1
device=none
T 53100 48700 5 10 1 1 0 1 1
value=IODATA5
}
C 52200 48900 1 0 0 io-1.sym
{
T 53100 49100 5 10 0 0 0 0 1
net=IODATA6:1
T 52400 49500 5 10 0 0 0 0 1
device=none
T 53100 49000 5 10 1 1 0 1 1
value=IODATA6
}
C 52200 49200 1 0 0 io-1.sym
{
T 53100 49400 5 10 0 0 0 0 1
net=IODATA7:1
T 52400 49800 5 10 0 0 0 0 1
device=none
T 53100 49300 5 10 1 1 0 1 1
value=IODATA7
}
C 49800 45800 1 0 0 input-2.sym
{
T 49800 46000 5 10 0 0 0 0 1
net=-DO_RDIO_1:1
T 50400 46500 5 10 0 0 0 0 1
device=none
T 50300 45900 5 10 1 1 0 7 1
value=\_DO_RDIO_1\_
}
N 51200 45900 51400 45900 4
