# 1. Introduction

Congratulations! You've figured out how to navigate this tutorial!
You're now reading Chapter 1.

[Cookoo](https://github.com/Masterminds/cookoo) is a general-purpose
chain-of-command library written in Go. In a nutshell, it is designed to help
you map a *route name* to a *chain of sequentially-run commands*.

A *route*, in Cookoo, can be thought of as a specific task to be done.
Let's give a really lame example:

> Task: Make a peanut butter and jelly sandwhich

Cookoo is a tool for taking the task and composing it out of a list of
sequentially running commands:

1. Get an empty plate
2. Place a slice of bread on the plate
3. Spread peanut butter on the bread
4. Spread jelly atop the peanut butter
5. Place a second slice of bread on top

While the illustration above captures the idea of a route (task) and a
chain of commands (sequence of steps), Cookoo actually *can't* make you
a PBJ sandwhich. Cookoo can build different kinds of apps, though. For
example, the `cookoo/cli` library (part of Cookoo) exists to simplify
writing command line apps.

For this tutorial, the focus will be on creating **web apps with
Cookoo**.

In the next part we'll set up Cookoo.

[git checkout 2_Setup](https://github.com/Masterminds/cookoo-web-tutorial/tree/2_Setup)
