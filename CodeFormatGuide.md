# Your Own Formatter #

In the `main.go` file there is a call to `NewDocCreator` which is given all the dependencies for walking a directory tree, extracting test names and splitting the names into words.

The last parameter to `NewDocCreator` is the `DocFormatter` that is used to print all the tests that were found. The interface looks like this:

```
type DocFormatter interface {
	Append(testFilePath string, sentences [][]string)
	Format() string
}
```

For every found test file (`*_test.go`) the Append function is called. The first parameter is the full path to that file, e.g. "`/some/path/my_test.go`". The second parameter is a list of sentences, where each sentence is in turn a list of words. These words are the result of the `NameChopper` which takes a test name and splits it into multiple words.

The `Append` function is called once for each test file. After all files are appended, the `Format` function is called. This must return, in a single string, the documentation that is the output of the program.

All you have to do to create your own format is create a new implementation of the `DocFormatter` and pass it in the main function to the `DocCreator`.

You can use the `ListDocFormatter` as a basis for your own formatter. To get a description of how it works, why don't you use the test2doc tool on its test file and see the results. Not only will you see the format but you can also read the description of the format that was extracted from the test file!