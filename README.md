Ijod is web application that allows you to synchronise the watching of
a video over multiple devices.

# How to install?

Building Ijod requires [Go][go] 1.16 or newer to be installed.

When everything has been installed, run

	$ go build ./cmd/ijod

and the file `ijod` should appear in the working directory.

# How to start?

Given the `ijod` binary, run

	$ ./ijod

and a server should start listening on port 8080, on all hosts.  Use

	$ ./ijod -help

to generate a listing of options.  This is a fat, static binary so it
can be copied wherever you might want to use it.

If you don't want to build, then run Ijod, you can also do it all in
one step:

	$ go run ./cmd/ijod

followed by any options you'd otherwise use.

# Further Notes

- Ijod has been primarily tested with Firefox and  Chromium.
- Since the video-files are streamed directly from your file system, it
  might be necessary to prepare them for streaming.

  For .mp4 files, running

      $ ffmpeg -i file.mp4 -movflags +faststart new-file.mp4

  with [FFmpeg][ffmpeg] should suffice to quickly create a streamable
  file. If not, consider converting it to the [WebM][webm] format using:

      $ ffmpeg -i file.ext -c:v libvpx-vp9 -c:a libopus -movflags +faststart file.webm

- Tools like [youtube-dl] should automatically produce files fit for
  streaming.
- If any issues or questions come up, send an email to my [public
  inbox][mail].

# Legal

Ijod is distributed under [CC0 Universal Public Domain Dedication][cc0]
(CC0 1.0).

[go]: https://golang.org/
[ffmpeg]: https://ffmpeg.org/
[webm]: https://www.webmproject.org/
[youtube-dl]: https://ytdl-org.github.io/youtube-dl/
[cc0]: https://creativecommons.org/publicdomain/zero/1.0/deed
[mail]: https://lists.sr.ht/~pkal/public-inbox
