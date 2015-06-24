v 20130925 2
C 40000 40000 0 0 0 title-B.sym
C 45100 42800 1 0 0 7474-1.sym
{
T 46900 44600 5 10 0 0 0 0 1
device=7474
T 46600 44800 5 10 1 1 0 6 1
refdes=U14
T 46900 45600 5 10 0 0 0 0 1
footprint=DIP14
T 45100 42800 5 10 0 0 0 0 1
slot=2
}
C 43600 43900 1 0 0 7408-1.sym
{
T 44300 44800 5 10 0 0 0 0 1
device=7408
T 43900 44800 5 10 1 1 0 0 1
refdes=U15
T 44300 46200 5 10 0 0 0 0 1
footprint=DIP14
T 43600 43900 5 10 0 0 0 0 1
slot=1
T 43600 43900 5 10 0 0 0 0 1
x-appdesc=Halt/skip logic
}
N 44900 44400 45100 44400 4
C 42000 44500 1 0 0 input-2.sym
{
T 42000 44700 5 10 0 0 0 0 1
net=DO_SKIP:1
T 42600 45200 5 10 0 0 0 0 1
device=none
T 42500 44600 5 10 1 1 0 7 1
value=DO_SKIP
}
C 42000 44100 1 0 0 input-2.sym
{
T 42000 44300 5 10 0 0 0 0 1
net=COND:1
T 42600 44800 5 10 0 0 0 0 1
device=none
T 42500 44200 5 10 1 1 0 7 1
value=COND
}
N 43400 44600 43600 44600 4
N 43400 44200 43600 44200 4
C 43500 43300 1 0 0 input-2.sym
{
T 43500 43500 5 10 0 0 0 0 1
net=CLK:1
T 44100 44000 5 10 0 0 0 0 1
device=none
T 44000 43400 5 10 1 1 0 7 1
value=CLK
}
N 44900 43400 45100 43400 4
C 47300 47100 1 0 0 7408-1.sym
{
T 48000 48000 5 10 0 0 0 0 1
device=7408
T 47600 48000 5 10 1 1 0 0 1
refdes=U15
T 48000 49400 5 10 0 0 0 0 1
footprint=DIP14
T 47300 47100 5 10 0 0 0 0 1
slot=2
}
C 45100 45800 1 0 0 7474-1.sym
{
T 46900 47600 5 10 0 0 0 0 1
device=7474
T 46600 47800 5 10 1 1 0 6 1
refdes=U14
T 46900 48600 5 10 0 0 0 0 1
footprint=DIP14
T 45100 45800 5 10 0 0 0 0 1
slot=1
T 45100 45800 5 10 0 0 0 0 1
x-appdesc=Halt/skip latches
}
C 49000 43100 1 0 0 7408-1.sym
{
T 49700 44000 5 10 0 0 0 0 1
device=7408
T 49300 44000 5 10 1 1 0 0 1
refdes=U15
T 49700 45400 5 10 0 0 0 0 1
footprint=DIP14
T 49000 43100 5 10 0 0 0 0 1
slot=3
}
C 43500 46300 1 0 0 input-2.sym
{
T 43500 46500 5 10 0 0 0 0 1
net=CLK:1
T 44100 47000 5 10 0 0 0 0 1
device=none
T 44000 46400 5 10 1 1 0 7 1
value=CLK
}
N 44900 46400 45100 46400 4
C 43500 47300 1 0 0 input-2.sym
{
T 43500 47500 5 10 0 0 0 0 1
net=DO_HALT:1
T 44100 48000 5 10 0 0 0 0 1
device=none
T 44000 47400 5 10 1 1 0 7 1
value=DO_HALT
}
N 44900 47400 45100 47400 4
N 46700 47400 47300 47400 4
{
T 46800 47400 5 10 1 1 0 0 1
netname=HALT
}
C 45500 48900 1 0 0 input-2.sym
{
T 45500 49100 5 10 0 0 0 0 1
net=ENABLE:1
T 46100 49600 5 10 0 0 0 0 1
device=none
T 46000 49000 5 10 1 1 0 7 1
value=ENABLE
}
C 49000 47500 1 0 0 output-2.sym
{
T 49900 47700 5 10 0 0 0 0 1
net=RUN:1
T 49200 48200 5 10 0 0 0 0 1
device=none
T 49900 47600 5 10 1 1 0 1 1
value=RUN
}
N 47300 47800 47100 47800 4
N 47100 47800 47100 49000 4
N 46900 49000 47100 49000 4
N 45900 45800 45900 45000 4
C 45300 45300 1 0 0 vcc-1.sym
N 45500 45300 45500 45200 4
N 45500 45200 45900 45200 4
C 45300 42300 1 0 0 vcc-1.sym
N 45900 42800 45900 42200 4
N 45900 42200 45500 42200 4
N 45500 42200 45500 42300 4
N 48600 47600 49000 47600 4
N 49000 43800 48800 43800 4
N 48800 43800 48800 47600 4
N 46700 43400 49000 43400 4
{
T 46800 43400 5 10 0 1 0 0 1
netname=-SKIP
}
C 50500 43500 1 0 0 output-2.sym
{
T 51400 43700 5 10 0 0 0 0 1
net=EXEC:1
T 50700 44200 5 10 0 0 0 0 1
device=none
T 51400 43600 5 10 1 1 0 1 1
value=EXEC
}
N 50300 43600 50500 43600 4
C 43500 48100 1 0 0 input-2.sym
{
T 43500 48300 5 10 0 0 0 0 1
net=-WAKE:1
T 44100 48800 5 10 0 0 0 0 1
device=none
T 44000 48200 5 10 1 1 0 7 1
value=\_WAKE\_
}
N 45900 48000 45900 48200 4
N 45900 48200 44900 48200 4
T 49900 40400 9 10 1 0 0 0 1
enable.sch
T 53800 40400 9 10 1 0 0 0 1
REVISION
T 50000 40700 9 10 1 0 0 0 1
Enable, halt and skip logic
C 49000 46200 1 0 0 7400-1.sym
{
T 49500 47100 5 10 0 0 0 0 1
device=7400
T 49300 47100 5 10 1 1 0 0 1
refdes=U16
T 49500 48450 5 10 0 0 0 0 1
footprint=DIP14
T 49000 46200 5 10 0 0 0 0 1
slot=1
T 49000 46200 5 10 0 0 0 0 1
x-appdesc=Halt/skip & instruction decoding logic
}
N 49000 46900 48800 46900 4
N 49000 46500 48800 46500 4
C 50500 46600 1 0 0 output-2.sym
{
T 51400 46800 5 10 0 0 0 0 1
net=-RUN:1
T 50700 47300 5 10 0 0 0 0 1
device=none
T 51400 46700 5 10 1 1 0 1 1
value=\_RUN\_
}
N 50300 46700 50500 46700 4
C 53800 44900 1 0 0 7408-1.sym
{
T 54500 45800 5 10 0 0 0 0 1
device=7408
T 54100 45800 5 10 1 1 0 0 1
refdes=U15
T 54500 47200 5 10 0 0 0 0 1
footprint=DIP14
T 53800 44900 5 10 0 0 0 0 1
slot=4
}
T 54400 44700 9 10 1 0 0 3 1
(unused slot)
C 53400 45800 1 0 0 vcc-1.sym
N 53800 45600 53600 45600 4
N 53600 45200 53600 45800 4
N 53800 45200 53600 45200 4
T 46800 43400 5 10 1 0 0 0 1
\_SKIP\_
