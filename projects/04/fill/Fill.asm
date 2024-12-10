// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/4/Fill.asm

// Runs an infinite loop that listens to the keyboard input. 
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel. When no key is pressed, 
// the screen should be cleared.

(LOOP)
  @SCREEN
  D=A
  @R0
  M=D

  (WHITE)
    @KBD
    D=M
    @FILL
    D;JGT

    @R0
    D=M

    @KBD
    D=D-A
    @LOOP
    D;JEQ

    @R0
    A=M
    M=0

    @R0
    M=M+1

    @WHITE
    0;JEQ
  
  (FILL)
    @KBD
    D=M
    @WHITE
    D;JEQ

    @R0
    D=M

    @KBD
    D=D-A
    @LOOP
    D;JEQ

    @1
    D=-A
    @R0
    A=M
    M=D

    @R0
    M=M+1

    @FILL
    0;JEQ