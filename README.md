# About
A package to write structured data to a bytes buffer, which is a technique commonly used in games where text-based protocols such as xml or json carry too much bulk to efficiently transfer data over the wire. The package is similar to Google's protocol buffers in that it is meant to be language agnostic with its intended function, but cuts out the overhead of using protocol buffers and compiling templates for each type.

# Installation
`go get github.com/Zeroeh/binio`

# Usage
Refer to `example.go` for usage and examples.

# Notes
The package was originally developed to be used for a specific game. You may have to change some of the functions to suit your needs. Note that there is no difference between Read/WriteString and Read/WriteUTFString other than the extra 2 bytes appended for size of the string.
