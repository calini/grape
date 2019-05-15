# Grape

This is a fork of `philipithomas/iterscraper`. Thanks, Philip I. Thomas.

The link can contain either an incrementing id or a token that can be passed from a file (more on that later).
Information is retrieved from HTML elements, and outputted as a CSV.

Thanks [Francesc](https://github.com/campoy) for featuring the original repo in episode #1 of [Just For Func](https://twitter.com/justforfunc). [Watch The Video](https://www.youtube.com/watch?list=PL64wiCrrxh4Jisi7OcCJIUpguV_f5jGnZ&v=eIWFnNz8mF4) or [Review Francesc's pull request](https://github.com/philipithomas/iterscraper/pull/1).

## Installation
```sh
go get go.ilie.io/grape
```

## Modes
There are three modes you can query for data.
1. Iterative
```sh
grape                               \
-url      https://github.com/%d     \
-from     100                       \
-to       105                       \
-query    ".p-name .p-org .p-label"
```
This mode will iterate over the indexes 100-105. (Interesting to see that username `100` exists)

2. Dictionary
```sh
grape                            \
-url      https://github.com/%s  \
-dict     dicts/users.txt        \
-query    ".p-name"
```
This mode will use a dictionary and query each term in it.

An example result looks like this
```
url                          id        .p-name
https://github.com/calini    calini    Calin Ilie
```

3. Dictionary range
```sh
grape                                \
-url      https://github.com/%s      \
-dict     dicts/users.txt            \
-from     2                          \
-to       4                          \
-query    ".p-name .p-org .p-label"
```
This mode will use a dictionary and query each term within the specified range.



## Selector Syntax
You can select HTML elements with classic JQUery syntax (thanks to GoQuery).
The only difference is, I have added the ability to use `§` as a separator to be able to for attributes of the element, not only it's text.
Example:
```sh
grape \
-url      https://github.com/%s       \
-dict     dicts/users.txt             \
-query    ".p-name .u-photo>img§src"
```
Will produce:
```
url                          id        .p-name       .u-photo>img§src 
https://github.com/calini    calini    Calin Ilie    https://avatars2.githubusercontent.com/u/9298529?s=460&v=4
```


## Flags

The manatory flag is `-url`.


For an explanation of the options, type `iterscraper -help`

General usage of iterscraper:

```
TODO REPLACE THIS WITH `grape -help`
```


## Errata

* On a `429 - too many requests` error, the app logs and continues, ignoring the request.
* On a `404 - not found` error, the system will log the miss, then continue. It is not exported to the CSV.
* The package will [follow up to 10 redirects](https://golang.org/pkg/net/http/#Get)
