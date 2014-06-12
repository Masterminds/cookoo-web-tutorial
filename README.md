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
