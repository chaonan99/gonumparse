# numparse
[![Build Status](https://drone.io/github.com/chaonan99/gonumparse/status.png)](https://drone.io/github.com/chaonan99/gonumparse/latest) [![GoDoc](https://godoc.org/github.com/chaonan99/gonumparse/numparse?status.svg)](https://godoc.org/github.com/chaonan99/gonumparse/numparse)

Package numparse is a [Go (Golang)](https://golang.org/) package converting number to English word.

## Features
* Support decimal number
* Very big number ...
* Input should be string

## Usage
* Import numparse by
    ```
    import github.com/chaonan99/gonumparse/numparse
    ```
* Parse number to word
    ```
    str := numparse.Parse("17")  // outputs "seventeen"
    str := numparse.Convert("1024")  // outputs "one thousand twenty-four"
    str := numparse.Parse("-.50")   // outputs "minus zero point five zero"
    ```

## Acknowledge
* Another [num2words](https://github.com/divan/num2words) converter in Go. This package has a little differences in the resulting word with numparse and only support integer.
* Inspired by [inflect](https://github.com/pwdyson/inflect.py) Python package.

## License
* MIT