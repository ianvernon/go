We love pull requests from everyone. By participating in this project, you
agree to abide by the thoughtbot [code of conduct].

We expect everyone to follow the code of conduct
anywhere in thoughtbot's project codebases,
issue trackers, chatrooms, and mailing lists.

[code of conduct]: https://thoughtbot.com/open-source-code-of-conduct

## Develop

Run `bin/setup` to get dependencies.

Run tests with `go test`.

Check your code with `bin/vet`.

Build with `go build`. Install locally with `go install`.

## Contributing

1. Fork the repo.

2. Run the tests. We only take pull requests with passing tests, and it's great
to know that you have a clean slate: `go test ./...`

3. Add a test for your change. Only refactoring and documentation changes
require no new tests. If you are adding functionality or fixing a bug, we need
a test!

4. Make the test pass.

5. Push to your fork and submit a pull request.

At this point you're waiting on us. We like to at least comment on, if not
accept, pull requests within three business days (and, typically, one business
day). We may suggest some changes or improvements or alternatives.

Some things that will increase the chance that your pull request is accepted:

  - the tests pass on CI
  - `bin/vet` runs without errors
  - `golint` prints no warnings

And in case we didn't emphasize it enough: we love tests!
