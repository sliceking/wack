# wack

I guess I'll take a swing

## WordPress action generator

Theres a lot of little pieces that take place when making an ajax request into your wordpress project and I always forget. This is a little tool that lets me actually forget and put faith in my tool.


## Installation

- Make sure you have the latest version of go. You can grab a copy [here](https://golang.org/dl/).
- `go get github.com/svwielga4/wack` to get a copy locally
- run `go build` in the newly downloaded directory. It should be `$GOPATH/src/github.com/svwielga4/wack`
- open up your .bash_profile and add the executable to your path


## Usage

### php action
- Run `wack php {actionName}` in a wordpress theme folder to append a stubbed action with linked function, in the functions.php file.

- Run `wack js {filename} {actionName}` to append a jQuery ajax request to a js file in your wordpress theme
