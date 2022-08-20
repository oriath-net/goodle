goodle
======

A [gooz](https://github.com/oriath-net/gooz)-compatible shim which links
against and calls the actual Oodle library for decompression.

To use this library to replace gooz, add the following line to your `go.mod`
and rebuild:

    replace github.com/oriath-net/gooz => github.com/oriath-net/goodle v1.0.0

You will need to obtain your own copy of the Oodle library to use this module.
The library is closed-source and cannot be freely distributed.
