// bitstype - Bit wise collection methods in the Go language
package main
/*
Program:  bitstype.go
Language: go
Author:   David Fahey - nzkiwi1g@gmail.com
Date:     07Apr2016
Purpose:  To demonstrate bit wise collection methods in the Go language.

History:  07Apr2016  Initial coding.                                    DAF
          02Feb2019  Added commentary                                   DAF

Notes:    Ref: https://golang.org/pkg/os/ 
          Bit numbering within integers is high order bit on the left, thus
          in a uint8 single unsigned integer, bit 7 refers to the left
          most bit or 0x80 and bit 0 refers to the right most bit or 0x01
          
Output:
          Go Demo Program: Bits
                Const BIT0= 0x01        00000001
                Const BIT1= 0x02        00000010
                Const BIT2= 0x04        00000100
                Const BIT3= 0x08        00001000
                Const BIT4= 0x10        00010000
                Const BIT5= 0x20        00100000
                Const BIT6= 0x40        01000000
                Const BIT7= 0x80        10000000
                Var b= 0x13     00010011
                Var c= 0x00     00000000
                c.Set(BIT7),    c= 0x80 10000000
                c.Set(BIT0),    c= 0x81 10000001
                b.Clear(BIT1),  b= 0x11 00010001
                b.Flip(BIT7),   b= 0x91 10010001
                b.Flip(BIT3),   b= 0x99 10011001
                b.Flip(BIT3),   b= 0x91 10010001
                b.Flip(BIT3),   b= 0x99 10011001
                b.Test(BIT4), t= true
                b.Test(BIT2), t= false

End.
*/

import (
    "fmt"
)

type Bits uint8

func (bits *Bits) Set(b Bits)   { *bits |= b }
func (bits *Bits) Clear(b Bits) { *bits &^= b }
func (bits *Bits) Flip(b Bits)  { *bits ^= b }
func (bits *Bits) Test(b Bits) bool {
    if *bits&b != 0 {
        return true
    } else {
        return false
    }
}
func (bits Bits) String() string { return fmt.Sprintf("%#.2x\t%.8b", uint8(bits), uint8(bits)) }

const (
    BIT0 Bits = 1 << iota
    BIT1
    BIT2
    BIT3
    BIT4
    BIT5
    BIT6
    BIT7
)

var (
    b Bits = 0x13
    c Bits = 0x00
)

func main() {

    fmt.Printf("Go Demo Program: Bits \n")

    // Print initial values
    fmt.Printf("\tConst BIT0= %s\n", BIT0)
    fmt.Printf("\tConst BIT1= %s\n", BIT1)
    fmt.Printf("\tConst BIT2= %s\n", BIT2)
    fmt.Printf("\tConst BIT3= %s\n", BIT3)
    fmt.Printf("\tConst BIT4= %s\n", BIT4)
    fmt.Printf("\tConst BIT5= %s\n", BIT5)
    fmt.Printf("\tConst BIT6= %s\n", BIT6)
    fmt.Printf("\tConst BIT7= %s\n", BIT7)
    fmt.Printf("\tVar b= %s\n", b)
    fmt.Printf("\tVar c= %s\n", c)

    // Do bit operations

    // Set a bit: BYTE |= (1<<bit#)
    c.Set(BIT7) // set left bit
    fmt.Printf("\tc.Set(BIT7),    c= %s\n", c)
    c.Set(BIT0) // set right bit
    fmt.Printf("\tc.Set(BIT0),    c= %s\n", c)

    // Clear a bit: BYTE &^= (1<<bit#)
    b.Clear(BIT1) // clear bit 1
    fmt.Printf("\tb.Clear(BIT1),  b= %s\n", b)

    // Toggle a bit: BYTE ^= (1<<bit#)
    b.Flip(BIT7) // flip bit 7
    fmt.Printf("\tb.Flip(BIT7),   b= %s\n", b)
    b.Flip(BIT3) // flip bit 3
    fmt.Printf("\tb.Flip(BIT3),   b= %s\n", b)
    b.Flip(BIT3) // flip bit 3
    fmt.Printf("\tb.Flip(BIT3),   b= %s\n", b)
    b.Flip(BIT3) // flip bit 3
    fmt.Printf("\tb.Flip(BIT3),   b= %s\n", b)

    // Test bit state BYTE & (1<<bit#)
    ok := b.Test(BIT4) // test bit 4
    fmt.Printf("\tb.Test(BIT4), t= %v \n", ok)
    ok = b.Test(BIT2) // test bit 2
    fmt.Printf("\tb.Test(BIT2), t= %v \n", ok)
}
