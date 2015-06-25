test2doc
========

Description
-----------

When creating unit tests according to the Behavior-Driven Development (BDD) philosophy, each test name is a simple sentence characterizing one aspect of the functionality of a unit. All test names together should then summarize the functionality of the whole unit. This summary can serve as a documentation for developers.

This tool walks through a file tree and extracts for every test unit that it finds such a summary. It prints all test names as sentences (splitting at camel cases and underscores) and thus gives the developer a comprehensive overview of the behavior of a unit under test.

Per default the tool prints its results to the standard output, which can easily be redirected to a file via the > or >> operators.

The tool is very similar to TestDox for C++ and TestDox or BDoc for Java.
