# 2. Setup

In this chapter we will get the environment all set up to begin building
web apps with Cookoo.

We'll assume that you have a working Go environment already.

## Installation

Installing Cookoo is no different from installing any other Go app. The
standard way to get it is to do this:

```
$ go get github.com/Masterminds/cookoo
```

I am a big fan of [gpm](https://github.com/pote/gpm), which uses a
`Godeps` file to declare a project's dependencies. You may notice the
`Godeps` file in this chapter. If you want to use `gpm` (and I suggest
also using [gvp](https://github.com/pote/gvp)), then you can skip the
`go get` step and do this instead:

```
$ gpm install
```

That's really all we need to do before we can head into our first Cookoo
program.

**Note:** Any time we include external libraries in our examples, they
will not only show up in the text and code, but also in the `Godeps`
file. So for any chapter, running `gpm install` should get your
dependencies all configured for you.

[git checkout 3_Hello_World](https://github.com/Masterminds/cookoo-web-tutorial/tree/3_Hello_World)
