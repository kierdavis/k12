v 20130925 2
C 40000 40000 0 0 0 title-B.sym
C 49000 43600 1 0 0 LCD-1.sym
{
T 50000 49175 5 10 0 0 0 0 1
device=LCD
}
C 48700 43500 1 0 0 gnd-1.sym
N 49000 43900 48800 43900 4
N 48800 43800 48800 48400 4
N 49000 48400 48800 48400 4
N 49000 47200 48800 47200 4
C 48200 48200 1 0 0 vcc-1.sym
N 48400 48200 48400 48100 4
N 48400 48100 49000 48100 4
C 47200 47400 1 0 0 input-2.sym
{
T 47200 47600 5 10 0 0 0 0 1
net=LCD_RS:1
T 47800 48100 5 10 0 0 0 0 1
device=none
T 47700 47500 5 10 1 1 0 7 1
value=LCD_RS
}
C 47200 46800 1 0 0 input-2.sym
{
T 47200 47000 5 10 0 0 0 0 1
net=LCD_E:1
T 47800 47500 5 10 0 0 0 0 1
device=none
T 47700 46900 5 10 1 1 0 7 1
value=LCD_E
}
C 47200 46500 1 0 0 input-2.sym
{
T 47200 46700 5 10 0 0 0 0 1
net=LCD_DB0:1
T 47800 47200 5 10 0 0 0 0 1
device=none
T 47700 46600 5 10 1 1 0 7 1
value=LCD_DB0
}
C 47200 46200 1 0 0 input-2.sym
{
T 47200 46400 5 10 0 0 0 0 1
net=LCD_DB1:1
T 47800 46900 5 10 0 0 0 0 1
device=none
T 47700 46300 5 10 1 1 0 7 1
value=LCD_DB1
}
C 47200 45900 1 0 0 input-2.sym
{
T 47200 46100 5 10 0 0 0 0 1
net=LCD_DB2:1
T 47800 46600 5 10 0 0 0 0 1
device=none
T 47700 46000 5 10 1 1 0 7 1
value=LCD_DB2
}
C 47200 45600 1 0 0 input-2.sym
{
T 47200 45800 5 10 0 0 0 0 1
net=LCD_DB3:1
T 47800 46300 5 10 0 0 0 0 1
device=none
T 47700 45700 5 10 1 1 0 7 1
value=LCD_DB3
}
C 47200 45300 1 0 0 input-2.sym
{
T 47200 45500 5 10 0 0 0 0 1
net=LCD_DB4:1
T 47800 46000 5 10 0 0 0 0 1
device=none
T 47700 45400 5 10 1 1 0 7 1
value=LCD_DB4
}
C 47200 45000 1 0 0 input-2.sym
{
T 47200 45200 5 10 0 0 0 0 1
net=LCD_DB5:1
T 47800 45700 5 10 0 0 0 0 1
device=none
T 47700 45100 5 10 1 1 0 7 1
value=LCD_DB5
}
C 47200 44700 1 0 0 input-2.sym
{
T 47200 44900 5 10 0 0 0 0 1
net=LCD_DB6:1
T 47800 45400 5 10 0 0 0 0 1
device=none
T 47700 44800 5 10 1 1 0 7 1
value=LCD_DB6
}
C 47200 44400 1 0 0 input-2.sym
{
T 47200 44600 5 10 0 0 0 0 1
net=LCD_DB7:1
T 47800 45100 5 10 0 0 0 0 1
device=none
T 47700 44500 5 10 1 1 0 7 1
value=LCD_DB7
}
N 48600 44500 49000 44500 4
N 48600 44800 49000 44800 4
N 48600 45100 49000 45100 4
N 48600 45700 49000 45700 4
N 48600 45400 49000 45400 4
N 48600 46000 49000 46000 4
N 48600 46300 49000 46300 4
N 48600 46600 49000 46600 4
N 48600 46900 49000 46900 4
N 48600 47500 49000 47500 4
C 46100 48300 1 270 0 pot-1.sym
{
T 47000 47500 5 10 0 0 270 0 1
device=VARIABLE_RESISTOR
T 46500 47600 5 10 1 1 0 0 1
refdes=R?
T 47600 47500 5 10 0 0 270 0 1
footprint=none
}
N 46700 47800 49000 47800 4
C 46100 47000 1 0 0 gnd-1.sym
C 46000 48400 1 0 0 vcc-1.sym
N 46200 48400 46200 48300 4
N 46200 47400 46200 47300 4
C 46300 44300 1 90 0 resistor-1.sym
{
T 45900 44600 5 10 0 0 90 0 1
device=RESISTOR
T 46000 45000 5 10 1 1 180 0 1
refdes=R?
}
C 46200 45300 1 90 0 switch-spst-1.sym
{
T 45500 45700 5 10 0 0 90 0 1
device=SPST
T 45900 45800 5 10 1 1 180 0 1
refdes=S?
}
C 46000 46200 1 0 0 vcc-1.sym
N 46200 46100 46200 46200 4
N 46200 45300 46200 45200 4
N 46200 44300 46200 44200 4
N 46200 44200 49000 44200 4
