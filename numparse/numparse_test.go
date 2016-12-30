package numparse

import (
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func TestConvert(t *testing.T) {
    Convey("Should convert correctly", t, func() {
        Convey("Small numbers should convert correctly", func() {
            So(Parse("0"), ShouldEqual, "zero")
            So(Parse("1"), ShouldEqual, "one")
            So(Parse("5"), ShouldEqual, "five")
            So(Parse("10"), ShouldEqual, "ten")
            So(Parse("11"), ShouldEqual, "eleven")
            So(Parse("12"), ShouldEqual, "twelve")
            So(Parse("17"), ShouldEqual, "seventeen")
        })
        Convey("Tens should convert correctly", func() {
            So(Parse("20"), ShouldEqual, "twenty")
            So(Parse("30"), ShouldEqual, "thirty")
            So(Parse("40"), ShouldEqual, "forty")
            So(Parse("50"), ShouldEqual, "fifty")
            So(Parse("60"), ShouldEqual, "sixty")
            So(Parse("90"), ShouldEqual, "ninety")
        })
        Convey("Combined numbers should convert correctly", func() {
            So(Parse("21"), ShouldEqual, "twenty-one")
            So(Parse("34"), ShouldEqual, "thirty-four")
            So(Parse("49"), ShouldEqual, "forty-nine")
            So(Parse("53"), ShouldEqual, "fifty-three")
            So(Parse("68"), ShouldEqual, "sixty-eight")
            So(Parse("99"), ShouldEqual, "ninety-nine")
        })
        Convey("Big numbers should convert correctly", func() {
            So(Parse("100"), ShouldEqual, "one hundred")
            So(Parse("200"), ShouldEqual, "two hundred")
            So(Parse("500"), ShouldEqual, "five hundred")
            So(Parse("123"), ShouldEqual, "one hundred and twenty-three")
            So(Parse("666"), ShouldEqual, "six hundred and sixty-six")
            So(Parse("1024"), ShouldEqual, "one thousand and twenty-four")
        })
        Convey("Negative numbers should convert correctly", func() {
            So(Parse("-123"), ShouldEqual, "minus one hundred and twenty-three")
        })
        Convey("Decimal numbers should convert correctly", func() {
            So(Parse("-.50"), ShouldEqual, "minus zero point five zero")
        })
    })
}

func ExampleParse() {
    var str string
    str = Parse("17")   // outputs "seventeen"
    str = Parse("10024") // outputs "ten thousand and twenty-four"
    str = Parse("-123.98") // outputs "minus one hundred and twenty-three point nine eight"
    str = Parse("-.50") // outputs "minus zero point five zero"
    _ = str
}