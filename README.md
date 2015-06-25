test2doc
========

Description
-----------

When creating unit tests according to the Behavior-Driven Development (BDD) philosophy, each test name is a simple sentence characterizing one aspect of the functionality of a unit. All test names together should then summarize the functionality of the whole unit. This summary can serve as a documentation for developers.

This tool walks through a file tree and extracts for every test unit that it finds such a summary. It prints all test names as sentences (splitting at camel cases and underscores) and thus gives the developer a comprehensive overview of the behavior of a unit under test.

Per default the tool prints its results to the standard output, which can easily be redirected to a file via the > or >> operators.

The tool is very similar to [TestDox](http://www.eld.leidenuniv.nl/~moene/Home/projects/testdox/) for C++ and [TestDox](http://agiledox.sourceforge.net/) or [BDoc](https://code.google.com/p/bdoc/) for Java.

Example
-------

Say you a have a test file called `Adder_test.go` that contains the following tests:

	func Test3Plus15Equals18(t *testing.T) { ... }
	func Test_any_number_plus_0_is_always_the_same(t *testing.T) { ... }
	func TestAddingMoreThanTwoNumbers_IsInsanelyComplex(t *testing.T) { ... }

then the resulting description of the unit will look like this:

 	Adder:
		- 3 Plus 15 Equals 18
		- any number plus 0 is always the same
		- Adding More Than Two Numbers Is Insanely Complex

Installation and Usage
----------------------

To install test2doc you can use the go tool:

	go get code.google.com/p/test2doc

This will download and build the project. It creates a binary executable in your GOPATH's bin folder. Assuming that you added the bin folder to your PATH you can now simply call the program with the root directory of your code as the only parameter. For example you can apply the tool to its own source by typing

	test2doc code.google.com/p/test2doc

The tool will go through that directory recursively and extract the tests from all found test files (files that end in `_test.go`). If the path is not a directory but a single test file, only that file will be documented. 