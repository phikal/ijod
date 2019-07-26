# Ijod

Ijod is a quick-to-setup web app, to easily watch videos at the same
time, over multiple hosts.

## How to use?

All you need is the `ijod` binary, and a directory with videos. Then run

	./ijod

and a server should start listening on port 8080, on all IP addresses. No
output should be generated, unless something has gone wrong.

## Configuration Options

The above example should cover most use-cases, but if not, Ijod can be
configured to act differently with a few extra flags.

These are:

- `-addr`: on what port to listen on.

	The default value is `:8080`, meaning on every address, on port
	8080, but `localhost:80` would mean locally on port 80.
- `-auth`: require a basic auth password.

	The argument MUST contain a colon (`:`) that separates the
	username-part from the password.
- `-words`: path to a file with words that will be used for names.

	To disable names, point to `/dev/null`, otherwise it will attempt to
	use `/usr/share/dict/words`.
- `-debug`: print debugging information on stdout.

## How to install?

Building Ijod requires [Go][go] to be installed, as well as
[go-bindata][bindata].

When everything has been installed, run

	go generate
	go build

and `ijod` should appear in the working directory.

## Further Notes

- Ijod has been tested with Firefox (primarily) and Chromium
  (secondarily).
- If the generated binary size is too large, using tools like [`strip
  -a`][strip] or [UPX][upx] might help to reduce the file size.
- Since the video-files are streamed directly from your file system, it
  might be necessary to prepare them.

  For most formats, running

		ffmpeg -i file.mp4 -c:v copy -c:a copy -movflags +faststart new-file.mp4

  with [FFmpeg][ffmpeg] should suffice to quickly create a streamable
  file. If not, consider converting it to the [WebM][webm] format using:

		ffmpeg -i file.mp4 -c:v libvpx-vp9 -c:a libopus -crf 23 -movflags +faststart file.webm
- Tools like [youtube-dl][] should automatically produce files fit for
  streaming.
- Ijod is inspired and conceptually based on an older project of mine,
  called [Doji][doji].
- If any issues or questions come up, send an email to my [public inbox][mail].

## Legal

Ijod is distributed under [CC0 Universal Public Domain Dedication][cc0]
(CC0 1.0).

[go]: https://golang.org/
[bindata]: https://github.com/jteeuwen/go-bindata
[strip]: https://man.openbsd.org/strip
[upx]: https://upx.github.io/
[ffmpeg]: https://ffmpeg.org/
[webm]: https://www.webmproject.org/
[youtube-dl]: https://ytdl-org.github.io/youtube-dl/
[cc0]: https://creativecommons.org/publicdomain/zero/1.0/deed
[doji]: http://zge.us.to/proj/doji/
[mail]: https://lists.sr.ht/~zge/public-inbox
