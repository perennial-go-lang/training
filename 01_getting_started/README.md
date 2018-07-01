# GETTING STARTED

In this section, we will focus on the core constructs and the basic program writing techniques of the language.

Go source code is stored in .go files. The filename consists of lower-case letters with multiple words separated by underscores. Spaces or any special characters are not recommended in a file name.

A typical Go program consists of keywords, constants, variables, operators, types and functions. Go is a case-sensitive language. A valid identifier name begins with a letter (any letter in the UTF-8) followed by zero or more letters or Unicode digits. An identifier may not start with a digit, be a keyword or have operators.

For ex:

__Valid Identifiers:__ A30, table2, _r24, i, β32\
__Invalid Identifiers:__ 3d (starts with a digit), defer (keyword in go), some-thing (no operators allowed)


The following are keywords in Go

__break__, __case__, __chan__, __const__, __continue__, __default__, __defer__, __else__, __fallthrough__, __for__, __func__, __go__, __goto__, __if__, __import__, __interface__, __map__, __pacakge__, __range__, __return__, __select__, __struct__, __switch__, __type__, __var__

Go also has a set of 36 *predeclared identifiers*. These include names of elementary types and some basic build-in functions. We shall explore these over the length of our training session.

__append__, __bool__, __byte__, __cap__, __close__, __complex__, __complex64__, __complex128__, __copy__, __false__, __float32__, __float64__, __imag__, __int__, __int8__, __int16__, __int32__, __int64__, __iota__, __len__, __make__, __new__, __nil__, __panic__, __print__, __println__, __real__, __recover__, __string__, __true__, __uint__, __uint8__, __uint16__, __uint32__, __uint64__, __uintptr__


