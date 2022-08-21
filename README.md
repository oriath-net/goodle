goodle
======

A [gooz](https://github.com/oriath-net/gooz)-compatible shim which links
against and calls the actual Oodle library for decompression.

To use this library to replace gooz, add the following line to your `go.mod`
and rebuild:

    replace github.com/oriath-net/gooz => github.com/oriath-net/goodle v1.1.0

You will need to obtain your own copy of the Oodle library to use this module.
The library is closed-source and cannot be freely distributed.


Windows
-------

The Oodle library must be in the current working directory or the same
directory as the Go executable. `oo2core_9_win64.dll` is preferred, but older
versions will also generally work.

If you have a newer version of the Oodle library or if it's in an unusual
location, set the environment variable `OODLE_DLL` to the path to that
library.


Linux / macOS
-------------

The library must be called `liboo2core.so` (or `liboo2core.dylib` on macOS)
and must be in one of your system's standard library directories, like
`/usr/local/lib`.

Finding these libraries may be difficult, and is left as an exercise to the
reader.
