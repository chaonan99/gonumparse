// Copyright 2016 Haonan Chen. All rights reserved.
// 2016-12-28

// Package numparse implements number (can be decimal) to English word converter.
package numparse

import (
    "bytes"
    "strings"
)

var number_args = map[string]string {
    "andword": "and",
    "one": "one",
    "zero": "zero",
}

var unit = []string{"", "one", "two", "three", "four", "five",
        "six", "seven", "eight", "nine"}
var teen = []string{"ten", "eleven", "twelve", "thirteen", "fourteen",
        "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var ten = []string{"", "", "twenty", "thirty", "forty",
       "fifty", "sixty", "seventy", "eighty", "ninety"}
var mill = []string{"", " thousand", " million", " billion", " trillion", " quadrillion",
        " quintillion", " sextillion", " septillion", " octillion",
        " nonillion", " decillion"}

func unitfn(units, mindex byte) string {
    return unit[units] + millfn(mindex)
}

func millfn(ind byte) string {
    if int(ind) > len(mill) - 1 {
        panic("number out of range")
    }
    return mill[ind]
}

func tenfn(tens, units, mindex byte) string {
    if tens != 1 {
        bi := ""
        if tens > 0 && units > 0 {
            bi = "-"
        }
        // This is a better way to concatenate string as suggested in
        // http://stackoverflow.com/questions/1760757/how-to-efficiently-concatenate-strings-in-go
        // A simple benchmark provided in test/strcat_test.go
        var result bytes.Buffer
        result.WriteString(ten[tens])
        result.WriteString(bi)
        result.WriteString(unit[units])
        result.WriteString(millfn(mindex))
        return result.String()
    }
    return teen[units] + mill[mindex]
}

func hundfn(hundreds, tens, units, mindex byte) string {
    if hundreds > 0 {
        andword := ""
        if tens > 0 || units > 0 {
            var tmp bytes.Buffer
            tmp.WriteString(" ")
            tmp.WriteString(number_args["andword"])
            tmp.WriteString(" ")
            andword = tmp.String()
        }
        var result bytes.Buffer
        result.WriteString(unit[hundreds])
        result.WriteString(" hundred")
        result.WriteString(andword)
        result.WriteString(tenfn(tens, units, 0))
        result.WriteString(millfn(mindex))
        return result.String()
    }
    if tens > 0 || units > 0 {
        return tenfn(tens, units, 0) + millfn(mindex)
    }
    return ""
}

func enword(num []byte) string {
    mill_count := byte(0)
    resRev := make([]string, 0)

    for i := len(num) - 1; i >= 0; i -= 3 {
        if i == 0 {
            resRev = append(resRev, hundfn(0, 0, num[i], mill_count))
        } else if i == 1 {
            resRev = append(resRev, hundfn(0, num[i-1], num[i], mill_count))
        } else {
            beforenum := hundfn(num[i-2], num[i-1], num[i], mill_count)
            resRev = append(resRev, beforenum)
            if i > 2 && mill_count == 0 && num[i-2] == 0 && beforenum != "" {
                resRev = append(resRev, " and ")
            } else if i > 2 && beforenum != "" {
                resRev = append(resRev, ", ")
            }
            mill_count ++
        }
    }
    var result bytes.Buffer
    for i := len(resRev) - 1; i >= 0; i -- {
        result.WriteString(resRev[i])
    }
    return result.String()
}

func integerfn(num string) string {
    if num == "" || num == "0" {
        return number_args["zero"]
    }
    btArray := []byte(num)
    for i, v := range btArray {
        btArray[i] = v - 48
    }
    return enword(btArray)
}

func signfn(sign byte) string {
    if sign == '+' {
        return ""
    } else if sign == '-' {
        return "minus "
    } else {
        panic("Invalid sign")
    }
}

func yieldnum(n byte) string {
    if n > 0 {
        return unit[n]
    } else if n == 0 {
        return number_args["zero"]
    } else {
        panic("yieldnum only handle 0-9")
    }
}

func decimalfn(num string) string {
    btArray := []byte(num)
    var result bytes.Buffer
    for _, v := range btArray {
        result.WriteString(" ")
        result.WriteString(yieldnum(v - 48))
    }
    return result.String()
}

// Parse convert input number to English word
func Parse(num string) string {
    sign := ""
    if num[0] == '-' || num[0] == '+' {
        sign = signfn(num[0])
        num = num[1:len(num)]
    }
    iandd := strings.Split(num, ".")
    var result bytes.Buffer
    if len(iandd) == 1 {
        result.WriteString(sign)
        result.WriteString(integerfn(iandd[0]))
        return result.String()
    } else if len(iandd) == 2 {
        result.WriteString(sign)
        result.WriteString(integerfn(iandd[0]))
        result.WriteString(" point")
        result.WriteString(decimalfn(iandd[1]))
        return result.String()
    } else {
        panic("Multiple point in a number")
    }
}
