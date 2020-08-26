# [Design Patterns in Golang](https://github.com/Evatix/golang-go-design-patterns)

Design patterns in golang

## Explore Instruction

## Make file labels to run

## Design Patterns Categories

[Reference](https://github.com/schleary/designPatterns/blob/master/Section%201%20-%20Intro/Design-Patterns-Introduction.pdf)

- [Creational Patterns](https://bit.ly/32sgRSg)
- [Structural Pattrns](https://bit.ly/2EscOxg)

## Important Patterns

Important or commonly used patterns are marked with astrek(*).

- Creational
  - *[Singleton](https://refactoring.guru/design-patterns/singleton) | [Singleton](https://www.udemy.com/course/design-patterns-go/learn/lecture/17118426#content)
  - *[Builder](https://refactoring.guru/design-patterns/builder) | [Builder in Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17012544#content)
  - *[Factory Method](https://refactoring.guru/design-patterns/factory-method) | [Factories](https://www.udemy.com/course/design-patterns-go/learn/lecture/17039320#content)
  - *[Decorator](https://refactoring.guru/design-patterns/decorator) | [Decorator Pattern Using Java](https://www.tutorialspoint.com/design_pattern/decorator_pattern.htm)) | [Decorator | Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17249382#content)
  - [Prototype](https://refactoring.guru/design-patterns/prototype) | [Prototype Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17115840#content)
- Behavioural
  - *[Strategy](https://refactoring.guru/design-patterns/strategy) | [Strategy | Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17345762#content)
  - [State](https://refactoring.guru/design-patterns/state)
  - [Chain of Responsibility](https://refactoring.guru/design-patterns/chain-of-responsibility)
  - [Command](https://refactoring.guru/design-patterns/command)
  - [Visitor](https://refactoring.guru/design-patterns/visitor)
  - *[Observer/Broadcasting/Publisher-Subscriber](https://refactoring.guru/design-patterns/observer)
  - [Iterator](https://refactoring.guru/design-patterns/iterator)
  - [Composite](https://refactoring.guru/design-patterns/composite) | [Composite in Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17247478#content)
  - [Template Method](https://refactoring.guru/design-patterns/template-method)
- Structural
  - [Facade](https://refactoring.guru/design-patterns/facade) | [Facade | Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17252788#content)
     - Repository (eg. [Repository](https://stackoverflow.com/questions/23213543/what-type-is-repository-pattern-in))
  - [Proxy](https://refactoring.guru/design-patterns/proxy) | [Proxy | Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17262546#content)
  - *[Adapter](https://refactoring.guru/design-patterns/adapter) | [Adapter Udemy](https://www.udemy.com/course/design-patterns-go/learn/lecture/17133984#content)
  - [Flyweight](https://refactoring.guru/design-patterns/flyweight), ie. general maps

### Takeaway: Summary of GO Modules, Vendors

- Go Modules are the way to GO
  - [vgo - Accessing local packages within a go module (go 1.11) - Stack Overflow](https://stackoverflow.com/questions/52026284/accessing-local-packages-within-a-go-module-go-1-11/55347424#55347424)
- Don't use vendors (because it requires sync, everytime new-packages is added)
  - Becareful with vendors and go-modules. Specially go vendors, replaces everything in vendor folder of the project on initial run.
- Keep one gopath in the system.
- Don't use multiple go modules init or go.mod file in one project.
- [Should be avoided] To use any local package, use the referencing concept.

**go.mod file example in the root project**:

```go
module YourProjectName

go 1.14

localPackage/IndifierName => relative or absolute path to the package

require (
    localPackage/IndifierName v0.0.1
)

```

- Use the package reference as `localPackage/IndifierName` -> `localPackage.call()`

## Contributors

- [Md. Alim Ul Karim](https://github.com/aukgit) | [Md. Alim Ul Karim Linkedin](https://bd.linkedin.com/in/alimkarim)
