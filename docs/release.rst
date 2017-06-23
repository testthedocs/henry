=======
Release
=======

Use https://github.com/goreleaser/goreleaser for releases.

Testing
=======

For testing use gox -> https://github.com/mitchellh/gox

Makking A Release
=================

- create branch
- tag branch like -> git tag 0.0.1-alpha.1
- push tag like -> git push origin 0.0.1-alpha.1
- export GitHub token -> export GITHUB_TOKEN=123456
- run goreleaser: gorelease

Check and adjust setting on GitHub release page
