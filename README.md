merf
----

`merf` is a pretty silly, small package that implements an interface conforming to the `*testing.T` type in the `testing` package.

WHY WAT HOW ... WHO?
--------------------

I have some projects with relatively __large__ test suites written in the xunit style of tests common to Go. After [Ginkgo](https://github.com/onsi/ginkgo) was released, we wanted to switch, but did not want to slowly transform our codebase over a period of months, or take the hit to productivity and dedicate a week to rewriting tests.

Enter [ginkgo-convert](https://github.com/tjarratt/ginkgo-convert), a cli tool that converts your existing tests to Ginkgo. A lot of tests already use the `testing.T` type to make assertions and make their tests fail, so conforming interface that can call into Ginkgo is needed to help make the transition. `merf` is that package.

`merf` sounds a bit like a fake error, and like a relatively unimportant package. Additionally, it's a bit like the sound that a cat would make.

![cat wizard](http://www.blueprintrecords.ca/wp-content/uploads/2012/08/cat-wizard.jpg)

Features
--------
* implements `Errorf`

If there is a testing.T method that you'd like implemented, I will gladly accept PRs, or issues for functionality.
