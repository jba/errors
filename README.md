# Errors in Go

This package provides support for augmenting Go errors with more
information.

Although it includes a single interface, `Error`, with the full set of methods
described here, it is not necessary to implement `Error` in its entirety.
Package-level functions like `Code` and `Source` can work with any error, and
each looks only for the method it needs to do its job.

## Classifying Errors

Numeric codes are the universally accepted way to classify errors across
languages and systems. Integers are superior to language types, values or
functions because they are easy to transmit between processes, and have the same
meaning in every programming language.

To deal with multiple sets of error codes, we use *spaces*. A space is a unique string
that fixes the interpretation of a set of error codes. For instance, the code
404 means "not found" in the HTTP space.

To participate in this scheme, an error should implement the method
```
ErrorCode() (space string, code int)
```
The space name should be the fully-qualified name of the Go type that describes
the code, or the import path of the package that holds the list of constants if
they are untyped. For example, the HTTP space is "net/http". This package
defines constants for common spaces.

## Chaining Errors

To wrap an underlying error, implement the method
```
ErrorSource() error
```
We avoid the word "cause" because it carries too much baggage: is the
InvalidArgument error you got from that RPC really the cause of the error you're
returning? Wasn't bad user input the actual cause?

## Additional Information

To provide additional information with an error, implement the method
```
ErrorDetails() interface{}
```
This is the place to put a stack trace.

