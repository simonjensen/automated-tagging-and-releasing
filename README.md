# automated-tagging-and-releasing

[![Build](https://github.com/simonjensen/automated-tagging-and-releasing/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/simonjensen/automated-tagging-and-releasing/actions/workflows/build.yaml) [![Create Release](https://github.com/simonjensen/automated-tagging-and-releasing/actions/workflows/create-release.yaml/badge.svg)](https://github.com/simonjensen/automated-tagging-and-releasing/actions/workflows/create-release.yaml)

> This repository serves to showcase how to automate the creation of tags and releases using GitHub Actions

### Usage

To run, test and build this code from a docker container, run the following command from the root of this repository:

```sh
docker run -it --rm -v $(pwd)/src:/app -w /app golang:1.23 bash
```

Alternatively you can install Go on your local machine, following the official guides:
- https://go.dev/dl/

Once you have an environment with Go up and running, play around with the following commands from the `src` directory:

```sh
# Run the main trigger
go run main.go

# Run the test(s)
go test ./hello (specific folder)
go test ./... (all folders)

# Build the application
go build .

# Run the build/executeable
./app
```

---
### Resources

- https://github.com/marketplace/actions/setup-go-environment
- https://gitversion.net/docs/learn/how-it-works

