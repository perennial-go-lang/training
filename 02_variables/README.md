# Variables

In this section, we will take a look at the various types in which we can declare a variable in Go. It is worth noting that in go, declaration and initialization happen together. Thus, any variable takes up a default value (zero value) upon declaration. Following are the zero values depending on the type of variable declared:

| Type | Zero Value |
|:---|:---|
|bool|false|
|string|0|
|int/int8 etc|0|
|float32/float64|0.0|

When declaring variables with an initial value, there are three ways to do so:

1. Declaration with specified type
2. Declaration with inferred type
3. Short-hand declaration