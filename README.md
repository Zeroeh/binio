# About
A package to write structured data to a bytes buffer, which is a technique commonly used in games where text-based protocols such as xml or json carry too much bulk to efficiently transfer data over the wire.

# Installation
`go get github.com/Zeroeh/binio`

# Usage


# Notes
The package was developed to be used for specific game. You may have to change some of the functions to suit your needs. Note that there is no difference between Read/WriteString and Read/WriteUTFString other than the extra 2 bytes appended for size of the string.
Eventually I plan to add something similar to "encoding/binary"'s Read() and Write() functions which can auto parse full structs to a bytes buffer.
