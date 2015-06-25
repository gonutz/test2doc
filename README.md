test2doc
========

Description
-----------

When creating unit tests according to the Behavior-Driven Development (BDD) philosophy, each test name is a simple sentence characterizing one aspect of the functionality of a unit. All test names together should then summarize the functionality of the whole unit. This summary can serve as a documentation for developers.

This tool walks through a file tree and extracts for every test unit that it finds such a summary. It prints all test names as sentences (splitting at camel cases and underscores) and thus gives the developer a comprehensive overview of the behavior of a unit under test.

Per default the tool prints its results to the standard output, which can easily be redirected to a file via the > or >> operators.

The tool is very similar to TestDox for C++ and TestDox or BDoc for Java.

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

