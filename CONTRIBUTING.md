# Contributing to Serve

Thank you so much for contributing to Serve. I appreciate your time and help.

## Code of Conduct

Be kind and respectful. Harassment of any kind will not be tolerated.

## Questions

## Filing a bug or feature

1. Before filing an issue, please check the existing issues to see if a
   similar one was already opened. If there is one already opened, feel free
   to comment on it.
2. If you believe you've found a bug, please provide detailed steps of
   reproduction. Also state the current behavior vs. the expected behavior.
3. If you'd like to see a feature or an enhancement please open an issue with
   a clear title and description of what the feature is and why it would be
   beneficial to the project and its users.

## Submitting changes

1. Tests: If you are submitting code, please ensure you have adequate tests
   for the feature. Tests can be run via `go test ./...` or `make test`.
1. Since this is golang project, ensure the new code is properly formatted to
   ensure code consistency. Run `make format`.

### Quick steps to contribute

1. Fork the project.
1. Download your fork to your PC (`git clone https://github.com/your_username/serve && cd serve`)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Make changes and run tests (`make build && make test`)
1. Add them to staging (`git add .`)
1. Commit your changes (`git commit -m 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new pull request
