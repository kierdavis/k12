v 20130925 2
C 40000 40000 0 0 0 title-B.sym
C 43900 43700 1 0 0 74157-1.sym
{
T 44200 48240 5 10 0 0 0 0 1
device=74157
T 44200 48040 5 10 0 0 0 0 1
footprint=DIP16
T 45600 47900 5 10 1 1 0 6 1
refdes=U5
T 43900 43700 5 10 0 0 0 0 1
x-appdesc=PC input multiplexer (bits 7-4)
}
C 42300 44200 1 0 0 input-2.sym
{
T 42300 44400 5 10 0 0 0 0 1
net=-DO_LJMP:1
T 42900 44900 5 10 0 0 0 0 1
device=none
T 42800 44300 5 10 1 1 0 7 1
value=\_DO_LJMP\_
}
C 43600 43400 1 0 0 gnd-1.sym
C 46100 47400 1 0 0 output-2.sym
{
T 47000 47600 5 10 0 0 0 0 1
net=PCMUX7:1
T 46300 48100 5 10 0 0 0 0 1
device=none
T 47000 47500 5 10 1 1 0 1 1
value=PCMUX7
}
C 46100 46600 1 0 0 output-2.sym
{
T 47000 46800 5 10 0 0 0 0 1
net=PCMUX6:1
T 46300 47300 5 10 0 0 0 0 1
device=none
T 47000 46700 5 10 1 1 0 1 1
value=PCMUX6
}
C 46100 45800 1 0 0 output-2.sym
{
T 47000 46000 5 10 0 0 0 0 1
net=PCMUX5:1
T 46300 46500 5 10 0 0 0 0 1
device=none
T 47000 45900 5 10 1 1 0 1 1
value=PCMUX5
}
C 46100 45000 1 0 0 output-2.sym
{
T 47000 45200 5 10 0 0 0 0 1
net=PCMUX4:1
T 46300 45700 5 10 0 0 0 0 1
device=none
T 47000 45100 5 10 1 1 0 1 1
value=PCMUX4
}
C 42300 44600 1 0 0 input-2.sym
{
T 42300 44800 5 10 0 0 0 0 1
net=PCADD4:1
T 42900 45300 5 10 0 0 0 0 1
device=none
T 42800 44700 5 10 1 1 0 7 1
value=PCADD4
}
C 42300 45000 1 0 0 input-2.sym
{
T 42300 45200 5 10 0 0 0 0 1
net=D4:1
T 42900 45700 5 10 0 0 0 0 1
device=none
T 42800 45100 5 10 1 1 0 7 1
value=D4
}
C 42300 45400 1 0 0 input-2.sym
{
T 42300 45600 5 10 0 0 0 0 1
net=PCADD5:1
T 42900 46100 5 10 0 0 0 0 1
device=none
T 42800 45500 5 10 1 1 0 7 1
value=PCADD5
}
C 42300 45800 1 0 0 input-2.sym
{
T 42300 46000 5 10 0 0 0 0 1
net=D5:1
T 42900 46500 5 10 0 0 0 0 1
device=none
T 42800 45900 5 10 1 1 0 7 1
value=D5
}
C 42300 46200 1 0 0 input-2.sym
{
T 42300 46400 5 10 0 0 0 0 1
net=PCADD6:1
T 42900 46900 5 10 0 0 0 0 1
device=none
T 42800 46300 5 10 1 1 0 7 1
value=PCADD6
}
C 42300 46600 1 0 0 input-2.sym
{
T 42300 46800 5 10 0 0 0 0 1
net=D6:1
T 42900 47300 5 10 0 0 0 0 1
device=none
T 42800 46700 5 10 1 1 0 7 1
value=D6
}
C 42300 47000 1 0 0 input-2.sym
{
T 42300 47200 5 10 0 0 0 0 1
net=PCADD7:1
T 42900 47700 5 10 0 0 0 0 1
device=none
T 42800 47100 5 10 1 1 0 7 1
value=PCADD7
}
C 42300 47400 1 0 0 input-2.sym
{
T 42300 47600 5 10 0 0 0 0 1
net=D7:1
T 42900 48100 5 10 0 0 0 0 1
device=none
T 42800 47500 5 10 1 1 0 7 1
value=D7
}
N 43700 47500 43900 47500 4
N 43700 47100 43900 47100 4
N 43700 46700 43900 46700 4
N 43700 46300 43900 46300 4
N 43700 45900 43900 45900 4
N 43700 45500 43900 45500 4
N 43700 45100 43900 45100 4
N 43700 44700 43900 44700 4
N 43700 44300 43900 44300 4
N 43900 43900 43700 43900 4
N 43700 43900 43700 43700 4
N 45900 47500 46100 47500 4
N 45900 46700 46100 46700 4
N 45900 45900 46100 45900 4
N 45900 45100 46100 45100 4
C 51300 43700 1 0 0 74157-1.sym
{
T 51600 48240 5 10 0 0 0 0 1
device=74157
T 51600 48040 5 10 0 0 0 0 1
footprint=DIP16
T 53000 47900 5 10 1 1 0 6 1
refdes=U6
T 51300 43700 5 10 0 0 0 0 1
x-appdesc=PC input multiplexer (bits 3-0)
}
C 49700 44200 1 0 0 input-2.sym
{
T 49700 44400 5 10 0 0 0 0 1
net=-DO_LJMP:1
T 50300 44900 5 10 0 0 0 0 1
device=none
T 50200 44300 5 10 1 1 0 7 1
value=\_DO_LJMP\_
}
C 51000 43400 1 0 0 gnd-1.sym
C 53500 47400 1 0 0 output-2.sym
{
T 54400 47600 5 10 0 0 0 0 1
net=PCMUX3:1
T 53700 48100 5 10 0 0 0 0 1
device=none
T 54400 47500 5 10 1 1 0 1 1
value=PCMUX3
}
C 53500 46600 1 0 0 output-2.sym
{
T 54400 46800 5 10 0 0 0 0 1
net=PCMUX2:1
T 53700 47300 5 10 0 0 0 0 1
device=none
T 54400 46700 5 10 1 1 0 1 1
value=PCMUX2
}
C 53500 45800 1 0 0 output-2.sym
{
T 54400 46000 5 10 0 0 0 0 1
net=PCMUX1:1
T 53700 46500 5 10 0 0 0 0 1
device=none
T 54400 45900 5 10 1 1 0 1 1
value=PCMUX1
}
C 53500 45000 1 0 0 output-2.sym
{
T 54400 45200 5 10 0 0 0 0 1
net=PCMUX0:1
T 53700 45700 5 10 0 0 0 0 1
device=none
T 54400 45100 5 10 1 1 0 1 1
value=PCMUX0
}
C 49700 44600 1 0 0 input-2.sym
{
T 49700 44800 5 10 0 0 0 0 1
net=PCADD0:1
T 50300 45300 5 10 0 0 0 0 1
device=none
T 50200 44700 5 10 1 1 0 7 1
value=PCADD0
}
C 49700 45000 1 0 0 input-2.sym
{
T 49700 45200 5 10 0 0 0 0 1
net=D0:1
T 50300 45700 5 10 0 0 0 0 1
device=none
T 50200 45100 5 10 1 1 0 7 1
value=D0
}
C 49700 45400 1 0 0 input-2.sym
{
T 49700 45600 5 10 0 0 0 0 1
net=PCADD1:1
T 50300 46100 5 10 0 0 0 0 1
device=none
T 50200 45500 5 10 1 1 0 7 1
value=PCADD1
}
C 49700 45800 1 0 0 input-2.sym
{
T 49700 46000 5 10 0 0 0 0 1
net=D1:1
T 50300 46500 5 10 0 0 0 0 1
device=none
T 50200 45900 5 10 1 1 0 7 1
value=D1
}
C 49700 46200 1 0 0 input-2.sym
{
T 49700 46400 5 10 0 0 0 0 1
net=PCADD2:1
T 50300 46900 5 10 0 0 0 0 1
device=none
T 50200 46300 5 10 1 1 0 7 1
value=PCADD2
}
C 49700 46600 1 0 0 input-2.sym
{
T 49700 46800 5 10 0 0 0 0 1
net=D2:1
T 50300 47300 5 10 0 0 0 0 1
device=none
T 50200 46700 5 10 1 1 0 7 1
value=D2
}
C 49700 47000 1 0 0 input-2.sym
{
T 49700 47200 5 10 0 0 0 0 1
net=PCADD3:1
T 50300 47700 5 10 0 0 0 0 1
device=none
T 50200 47100 5 10 1 1 0 7 1
value=PCADD3
}
C 49700 47400 1 0 0 input-2.sym
{
T 49700 47600 5 10 0 0 0 0 1
net=D3:1
T 50300 48100 5 10 0 0 0 0 1
device=none
T 50200 47500 5 10 1 1 0 7 1
value=D3
}
N 51100 47500 51300 47500 4
N 51100 47100 51300 47100 4
N 51100 46700 51300 46700 4
N 51100 46300 51300 46300 4
N 51100 45900 51300 45900 4
N 51100 45500 51300 45500 4
N 51100 45100 51300 45100 4
N 51100 44700 51300 44700 4
N 51100 44300 51300 44300 4
N 51300 43900 51100 43900 4
N 51100 43900 51100 43700 4
N 53300 47500 53500 47500 4
N 53300 46700 53500 46700 4
N 53300 45900 53500 45900 4
N 53300 45100 53500 45100 4
T 50000 40700 9 10 1 0 0 0 1
PC input multiplexer (bits 7-0)
T 49900 40400 9 10 1 0 0 0 1
pcmuxl.sch
