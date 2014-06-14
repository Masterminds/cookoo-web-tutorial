# 5. Routes

In the last chapter we built a basic `Hello Web` app. This chapter
focuses on naming routes for web apps.

We've seen two route names in the last few chapters:

* `hello`: We used this when we were creating a from-scratch non-web
  application.
* `GET /`: We used this in the last chapter to declare that our route is
  listening for `GET` requests to the document root.

The code in this chapter begins with `GET /` and gives several other
examples:

* `GET /hello`: This listens for `GET` requests at the path `/hello`.
* `POST /hello`: This listens for only `POST` requests at the path
  `/hello`.
* `* /goodeby`: This listens for any request to the path `/goodbye`. It
  will answer whether the verb is `GET`, `POST`, `DELETE`, or anything
  else.
* `GET /goodbye/*`: This will listen for `GET` requests on any path that
  begins with `/goodbye/` and doesn't contain slashes. Examples:
  - `GET /goodbye/foo`
  - `GET /goodbye/some-long-string`
* `GET /goodbye/**`: This will listen for `GET` requests for any path
  that starts with `/goodbye/`. Examples:
  - `GET /goodbye/foo`
  - `GET /goodbye/foo/bar/baz`

Take a moment to start up the server and try a few requests. Here's a
basic `curl` command to get you started. (`-X` lets you tell curl what
method you want to use):

```
$ curl -v -X POST localhost:8080/goodbye
```

As you test things out, you may discover that:

1. Order is important. Since `/goodbye/*` comes before `/goodbye/**`, it
   will be matched first.
2. Cookoo will allow non-standard HTTP methods. This is great for
   protocols like WebDAV that specify additional methods.
3. Cookoo views slashes as important characters. Omitting them may have
   consequences.

If you're feeling wild, you can try building some patterns of your own.
I suggest experimenting with something like this:

```
GET /t?st/*/science
```

**Tip:** Cookoo's path patching library supports POSIX-style regular
expressions. Check out Go's `path.Match` function to get an idea for the
syntax.

[git checkout 6_Request_Data](https://github.com/Masterminds/cookoo-web-tutorial/tree/6_Request_Data)
