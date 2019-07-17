# Instructions

In this challenge we will assess your coding skills. Create a little project implementing a solution for the problem shown below. After you're done, please `zip` or `tar` all needed files and send the compressed archive to the shared folder sent to you by email.

# Number Translation Problem

You have arrived in a long forgotten island, called kwego. After 3 days on this island, you realized that their numeric system is the same as the [roman system](https://en.wikipedia.org/wiki/Roman_numerals) but with different names and came up with the following kwegonian to roman translation table:

```
kil 	I
jin 	V
pol 	X
kilow 	L
jij 	C
jinjin	D
polsx	M
```

To make your life easier, you will build an Rest API to translate their kwegonian numbers into DECIMAL numbers.

Here's an example of the request payload

```
{
 Â kwego: "polsx polsx pol jin kil"
}
```

And the response payload is

```
{
 Â success: true,
 Â decimal: 2016
}
```

In case of invalid numbers, `success` key must be false, and a `error_msg` key may be added to receive the error string. 

Make the API more restful as it makes sense for you.

You can use external libraries to build and test, but not to do the translation itself.

# Packages
github.com/gorilla/mux

# Usage
```
go run main.go
```
