# Fork of MIRA

For the original project by @thecsw, please see https://github.com/thecsw/mira/.

### Why

While I liked the base of the original library, especially [mira/models](https://github.com/thecsw/mira/models) was full of duplicate code. In this fork, I tried to tailor it down and have a more "go" approach at handling the different types of JSON data returned from Reddit, while minimizing duplicated code.

Because of the project I need the library for, I also added some mod relevant stuff that I haven't found in any other Go Reddit API Wrapper so far - mainly mod functions.

### How

Please read the documentation over at https://pkg.go.dev/github.com/ttgmpsn/mira. It has some examples to get you started.