# Text Finder

Text Finder is a search tool that works similarly to `grep`. It utilizes Go's concurrency features to dramatically speed up the search process. While it is (currently) not as flexible as `grep`, for targeted uses its performance is much better, particularly when a user needs to search a large file system.

### Build:

`$ make build`

This will compile and save the executable binary as `./bin/text-finder`.

### Run unit tests:

`$ make test`

### Usage:

`$ ./bin/text-finder -r -d ~/Documents foo bar baz`

This will search for the strings `["foo", "bar", "baz"]` in all files recursively starting at `~/Documents`. It will search all lines of each non-binary file.

### Notes:

Potential future features:

- Option to terminate search after 1st match
- Option to exclude specific files
- Option to load search terms from file, and write results to file and/or copy matching files to folder
