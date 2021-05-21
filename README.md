# Count pan kills in tf2

A quick and dirty pan kill counter.

## Usage

1. Download `tf2-pan-kill-stream-browser-source` or build it using go
2. Go to `Steam > Library > TF2 > Properties > General > Launch options` and add `-conclearlog -condebug`. This will generate a `console.log` file in your `tf` folder.
3. Start `tf2-pan-kill-stream-browser-source` with your Steam username and path to the `console.log` file (e.g. `tf2-pan-kill-stream-browser-source Avoonix "C:\Program Files (x86)\Steam\steamapps\common\Team Fortress 2\tf\console.log"` if you're on windows)
4. Add `http://localhost:6969` as browser source

## Building

```bash
GOOS=windows GOARCH=amd64 go build
# or
GOOS=linux GOARCH=amd64 go build
```
