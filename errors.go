package viron

import "fmt"

const ParseErrorFormat = "%s value \"%s\" could not be parsed into %s value"

type ParseError string

func NewParseError(name, value, kind string) ParseError {
    return ParseError(fmt.Sprintf(ParseErrorFormat, name, value, kind))
}

func (err ParseError) Error() string {
    return string(err)
}

type InvalidArgumentError string

func NewInvalidArgumentError(value interface{}) InvalidArgumentError {
    return InvalidArgumentError(fmt.Sprintf("%v is not a non-zero struct pointer", value))
}

func (err InvalidArgumentError) Error() string {
    return string(err)
}

type RequiredFieldError string

func NewRequiredFieldError(name string) RequiredFieldError {
    return RequiredFieldError(fmt.Sprintf("%s is a required environment variable", name))
}

func (err RequiredFieldError) Error() string {
    return string(err)
}
