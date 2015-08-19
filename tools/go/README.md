This directory contains implementations in [Go][go] of tools for working with
the K12 computer.

The tools are designed to be built using [gb][gb], a project-based build tool
that functions similarly to the standard `go` tool for compilation.

[go]: http://golang.org/
[gb]: http://getgb.io/

# Build instructions

As per the rules for a `gb` workspace, the source code lives inside the `src`
subdirectory and is organised into packages, much like `GOPATH` directories.
However, since the packages defined here aren't globally unique, they don't
have to have globally unique identifiers (for example, the `emulator` package
is referred to simply as `"emulator"`, not as
`"github.com/kierdavis/k12/(etc)/emulator"`). The `vendor` subdirectory contains
"vendored" code; essentially frozen snapshots of continously integrated packages
available on Github.

First, install `gb` if you haven't already:

    go get github.com/constabulary/gb/...

Then, make sure you are in this directory (`REPO/tools/go`) and run:

    gb build all

to build all the tools. The compiled packages are placed in a directory named
`pkg`, and the executables in a directory named `bin` (again much like a normal
Go workspace).
