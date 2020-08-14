![GitHub Workflow Status](https://img.shields.io/github/workflow/status/iverly/docker-containers-visualizer/build)
[![Go Report Card](https://goreportcard.com/badge/github.com/iverly/docker-containers-visualizer)](https://goreportcard.com/report/github.com/iverly/docker-containers-visualizer)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/iverly/docker-containers-visualizer)
![GitHub](https://img.shields.io/github/license/iverly/docker-containers-visualizer)


# docker-containers-visualizer

Visualize all container running on the host with networks, stats and others informations !

![image](https://cdn-paul.iverly.net/gh/docker-containers-visualizer/hero.819282.png)

## Installation

Using `curl`:
```bash
curl -o /usr/bin/dcv https://github.com/...
```
Using releases:

[Click here](https://github.com/iverly/docker-containers-visualizer/releases)

## Usage

#### Basic usage
```bash
$ dcv
```

##### Custom docker host
```bash
$ export DOCKER_HOST=unix:///path/to/daemon
$ dcv
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
