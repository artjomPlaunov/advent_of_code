Advent of Code in Golang

Some random text notes from this year; my favorite problem (as of writing, I still have to do 20,21,24,25) is day17 where you disassemble to create a quine:

cat disassemble.txt 
Register A: 22571680
Register B: 0
Register C: 0

Program: 2,4,1,3,7,5,0,3,4,3,1,5,5,5,3,0

START: 
  B = A%8        	(2,4)
  B = B^3    		(1,3)
  C = A/(2**B)  		(7,5)
  A = A/(2**3)  		(0,3)
  B = B^C    		(4,3)
  B = B^5    		(1,5)
  OUT += B 		(5,5)	
  If A != 0 GOTO START	(3,0)

Some scenarios I pictured for the moving boxes problem:

cat boxes.txt
##############
##......##..##
##...[][]...##
##....[]....##
##.....@....##
##..........##


..............
..............
......[].........
.......[][]..........
......[]..#.........
.....[]..[].............
......[][]...............
.......[]................
.......@..................
..............................
