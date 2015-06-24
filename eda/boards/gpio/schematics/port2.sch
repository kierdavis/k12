v 20130925 2
C 40000 40000 0 0 0 title-B.sym
C 45500 44100 1 0 0 74377-1.sym
{
T 45800 47550 5 10 0 0 0 0 1
device=74377
T 45800 47750 5 10 0 0 0 0 1
footprint=DIP20
T 47200 47400 5 10 1 1 0 6 1
refdes=U?
T 45500 44100 5 10 0 0 0 0 1
x-appdesc=Output port 2 (device control) register
}
C 43900 44500 1 0 0 input-2.sym
{
T 43900 44700 5 10 0 0 0 0 1
net=-DO_WRIO_2:1
T 44500 45200 5 10 0 0 0 0 1
device=none
T 44400 44600 5 10 1 1 0 7 1
value=\_DO_WRIO_2\_
}
C 43900 44200 1 0 0 input-2.sym
{
T 43900 44400 5 10 0 0 0 0 1
net=CLK:1
T 44500 44900 5 10 0 0 0 0 1
device=none
T 44400 44300 5 10 1 1 0 7 1
value=CLK
}
N 45300 47000 45500 47000 4
N 45300 46700 45500 46700 4
N 45300 46400 45500 46400 4
N 45300 45800 45500 45800 4
N 45500 46100 45300 46100 4
N 45300 45500 45500 45500 4
N 45300 45200 45500 45200 4
N 45300 44900 45500 44900 4
N 45300 44600 45500 44600 4
N 45300 44300 45500 44300 4
N 47500 47000 49300 47000 4
{
T 48000 47000 5 10 1 1 0 0 1
netname=LCD_E
}
N 47500 46700 49300 46700 4
{
T 48000 46700 5 10 1 1 0 0 1
netname=LCD_RS
}
N 47500 46400 49300 46400 4
{
T 48000 46400 5 10 1 1 0 0 1
netname=SPI0_SS
}
N 47500 46100 49300 46100 4
{
T 48000 46100 5 10 1 1 0 0 1
netname=SPI1_SS
}
N 47500 45800 49300 45800 4
N 47500 45500 49300 45500 4
N 47500 45200 49300 45200 4
N 47500 44900 49300 44900 4
C 49300 44300 1 0 0 74244-1.sym
{
T 49600 47450 5 10 0 0 0 0 1
device=74244
T 51000 47300 5 10 1 1 0 6 1
refdes=U?
T 49600 47650 5 10 0 0 0 0 1
footprint=DIP20
T 49300 44300 5 10 0 0 0 0 1
x-appdesc=Input port 2 (device control) driver
}
N 49900 44300 49900 44100 4
N 49700 44100 50700 44100 4
N 50700 44100 50700 44300 4
C 51500 46900 1 0 0 io-1.sym
{
T 52400 47100 5 10 0 0 0 0 1
net=IODATA0:1
T 51700 47500 5 10 0 0 0 0 1
device=none
T 52400 47000 5 10 1 1 0 1 1
value=IODATA0
}
C 51500 46600 1 0 0 io-1.sym
{
T 52400 46800 5 10 0 0 0 0 1
net=IODATA1:1
T 51700 47200 5 10 0 0 0 0 1
device=none
T 52400 46700 5 10 1 1 0 1 1
value=IODATA1
}
C 51500 46300 1 0 0 io-1.sym
{
T 52400 46500 5 10 0 0 0 0 1
net=IODATA2:1
T 51700 46900 5 10 0 0 0 0 1
device=none
T 52400 46400 5 10 1 1 0 1 1
value=IODATA2
}
C 51500 46000 1 0 0 io-1.sym
{
T 52400 46200 5 10 0 0 0 0 1
net=IODATA3:1
T 51700 46600 5 10 0 0 0 0 1
device=none
T 52400 46100 5 10 1 1 0 1 1
value=IODATA3
}
C 51500 45700 1 0 0 io-1.sym
{
T 52400 45900 5 10 0 0 0 0 1
net=IODATA4:1
T 51700 46300 5 10 0 0 0 0 1
device=none
T 52400 45800 5 10 1 1 0 1 1
value=IODATA4
}
C 51500 45400 1 0 0 io-1.sym
{
T 52400 45600 5 10 0 0 0 0 1
net=IODATA5:1
T 51700 46000 5 10 0 0 0 0 1
device=none
T 52400 45500 5 10 1 1 0 1 1
value=IODATA5
}
C 51500 45100 1 0 0 io-1.sym
{
T 52400 45300 5 10 0 0 0 0 1
net=IODATA6:1
T 51700 45700 5 10 0 0 0 0 1
device=none
T 52400 45200 5 10 1 1 0 1 1
value=IODATA6
}
C 51500 44800 1 0 0 io-1.sym
{
T 52400 45000 5 10 0 0 0 0 1
net=IODATA7:1
T 51700 45400 5 10 0 0 0 0 1
device=none
T 52400 44900 5 10 1 1 0 1 1
value=IODATA7
}
N 51300 47000 51500 47000 4
N 51300 46700 51500 46700 4
N 51300 46400 51500 46400 4
N 51300 46100 51500 46100 4
N 51300 45800 51500 45800 4
N 51300 45500 51500 45500 4
N 51300 45200 51500 45200 4
N 51300 44900 51500 44900 4
C 45300 47100 1 180 0 io-1.sym
{
T 44400 46900 5 10 0 0 180 0 1
net=IODATA0:1
T 45100 46500 5 10 0 0 180 0 1
device=none
T 44400 47000 5 10 1 1 180 1 1
value=IODATA0
}
C 45300 46800 1 180 0 io-1.sym
{
T 44400 46600 5 10 0 0 180 0 1
net=IODATA1:1
T 45100 46200 5 10 0 0 180 0 1
device=none
T 44400 46700 5 10 1 1 180 1 1
value=IODATA1
}
C 45300 46500 1 180 0 io-1.sym
{
T 44400 46300 5 10 0 0 180 0 1
net=IODATA2:1
T 45100 45900 5 10 0 0 180 0 1
device=none
T 44400 46400 5 10 1 1 180 1 1
value=IODATA2
}
C 45300 46200 1 180 0 io-1.sym
{
T 44400 46000 5 10 0 0 180 0 1
net=IODATA3:1
T 45100 45600 5 10 0 0 180 0 1
device=none
T 44400 46100 5 10 1 1 180 1 1
value=IODATA3
}
C 45300 45900 1 180 0 io-1.sym
{
T 44400 45700 5 10 0 0 180 0 1
net=IODATA4:1
T 45100 45300 5 10 0 0 180 0 1
device=none
T 44400 45800 5 10 1 1 180 1 1
value=IODATA4
}
C 45300 45600 1 180 0 io-1.sym
{
T 44400 45400 5 10 0 0 180 0 1
net=IODATA5:1
T 45100 45000 5 10 0 0 180 0 1
device=none
T 44400 45500 5 10 1 1 180 1 1
value=IODATA5
}
C 45300 45300 1 180 0 io-1.sym
{
T 44400 45100 5 10 0 0 180 0 1
net=IODATA6:1
T 45100 44700 5 10 0 0 180 0 1
device=none
T 44400 45200 5 10 1 1 180 1 1
value=IODATA6
}
C 45300 45000 1 180 0 io-1.sym
{
T 44400 44800 5 10 0 0 180 0 1
net=IODATA7:1
T 45100 44400 5 10 0 0 180 0 1
device=none
T 44400 44900 5 10 1 1 180 1 1
value=IODATA7
}
C 48300 44000 1 0 0 input-2.sym
{
T 48300 44200 5 10 0 0 0 0 1
net=-DO_RDIO_2:1
T 48900 44700 5 10 0 0 0 0 1
device=none
T 48800 44100 5 10 1 1 0 7 1
value=\_DO_RDIO_2\_
}
T 50000 40700 9 10 1 0 0 0 1
IO port 2 (device control port)
T 49900 40400 9 10 1 0 0 0 1
port2.sch