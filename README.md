# knex
Dependency injection framework for [Go](http://golang.org/).

## Installation and Docs

Install using `go get github.com/chrisehlen/knex`.

Full documentation is available at
https://godoc.org/github.com/chrisehlen/knex

## Usage

Knex is an easy to use api that uses Go tags to define relationships between components.

Examples below can be found at [knex-example](https://github.com/chrisehlen/knex-example/)

**Define component(s) that provides an implementation**

```go
type StringReaderImpl struct {
	spi.Reader `provide:"resource" "scope:"factory"`
}
func (self *StringReaderImpl) Inject() error {return nil}
func (self *StringReaderImpl) Read() (string, error) {...}
```

[StringReaderImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/StringReaderImpl.go) is an implementation of the [Reader](https://github.com/chrisehlen/knex-example/blob/master/spi/Reader.go) interface that uses a constant string as input. [StringReaderImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/StringReaderImpl.go) has no requires so it's Inject() method has zero arguments.

```go
type UpperCaseFilterImpl struct {
	spi.Filter `provide:"resource"`
}
func (self *UpperCaseFilterImpl) Inject() error {return nil}
func (self *UpperCaseFilterImpl) Do(message string) (string, error) {...}
```

[UpperCaseFilterImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/UpperCaseFilterImpl.go) is an implementation of the [Filter](https://github.com/chrisehlen/knex-example/blob/master/spi/Filter.go) interface that converts all characters of the message to upper case. [UpperCaseFilterImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/UpperCaseFilterImpl.go) has no requires so it's Inject() method has zero arguments.

```go
type PigLatinFilterImpl struct {
spi.Filter `provide:"resource"`
}
func (self *PigLatinFilterImpl) Inject() error { return nil }
func (self *PigLatinFilterImpl) Do(message string) (string, error) {...}
```

[PigLatinFilterImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/PigLatinFilterImpl.go) is an implementation of the [Filter](https://github.com/chrisehlen/knex-example/blob/master/spi/Filter.go) interface that converts each word, of the message, to its [Pig Latin](https://en.wikipedia.org/wiki/Pig_Latin) equivalent. [PigLatinFilterImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/PigLatinFilterImpl.go) has no requires so it's Inject() method has zero arguments.

```go
type ConsoleWriterImpl struct {
	spi.Writer `provide:"resource" "scope:"factory"`
}
func (self *ConsoleWriterImpl) Inject() error {return nil}
func (self *ConsoleWriterImpl) Write(message string) error {...}
```

[ConsoleWriterImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/ConsoleWriterImpl.go) is an implementation of the [Writer](https://github.com/chrisehlen/knex-example/blob/master/spi/Writer.go) interface that sends all input to standard output. [ConsoleWriterImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/ConsoleWriterImpl.go) has no requires so it's Inject() method has zero arguments.

**Define component(s) that require implementation(s)**

```go
type SimpleControllerImpl struct {
	api.Controller `provide:"resource" id:"controller" "scope:"graph"`
	reader     Reader   `require:"true"`
	filters    []Filter `require:"false"`
	writer     Writer   `require:"true"`
}
func (self *SimpleControllerImpl) Inject(reader Reader, filters []Filter, writer Writer) error {
	self.reader = reader
	self.filters = filters
	self.writer = writer
	return nil
}
func (self *SimpleControllerImpl) ReadMessage() (string, error) {...}
func (self *SimpleControllerImpl) ApplyFilters(string) (string, error) {...}
func (self *SimpleControllerImpl) WriteMessage(string) error {...}
```

[SimpleControllerImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/SimpleControllerImpl.go) is an implementation of the [Controller](https://github.com/chrisehlen/knex-example/blob/master/api/Controller.go) interface that delegates calls to its dependencies. [SimpleControllerImpl](https://github.com/chrisehlen/knex-example/blob/master/lib/SimpleControllerImpl.go) has three requires so it's Inject() method has three arguments which are in the order they are defined in the struct.

**Register components**

```go
	// Register controller implementation
	knex.DefaultFactory.Register(new(lib.ControllerImpl))

	// Register SPI implementations
	spiFactory := knex.NewFactory()
	spiFactory.Register(new(lib.ConsoleWriterImpl))
	spiFactory.Register(new(lib.PigLatinFilterImpl))
	spiFactory.Register(new(lib.StringReaderImpl))
	spiFactory.Register(new(lib.UpperCaseFilterImpl))
```

[Factory.Register(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.Register) method associates a given implementation to an interface type.

**Parent factories**

```go
knex.DefaultFactory.AddParent(spiFactory)
```

[Factory.AddParent(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.AddParent) assignes a parent factory to a child factory.  If an implementation can not be retrieved from a child factory then the child factory will check if the implementation can be retrieved from one of its parents.

**Get an implementation by type**

```go
iController, err := knex.DefaultFactory.GetByType(new(api.Controller))
controller := iController.(api.Controller)
```

[Factory.GetByType(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.GetByType) gets an implementation based on a given interface type.  If an implementaion has not been registered or multiple implementations have been registered for this interface type [Factory.GetByType(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.GetByType) will return an error.

**Get all implementations by type**

```go
iFilters, err := knex.DefaultFactory.GetAllOfType(new(spi.Filter))
filters := iFilters.([]spi.Filter)
```

[Factory.GetAllOfType(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.GetAllOfType) gets all implementations based on a given interface type.  If an implementation has not been registered [Factory.GetAllOfType(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.GetAllOfType) returns an empty slice.

**Get an implementation by id**

```go
iController, err := knex.DefaultFactory.GetById("controller")
controller := iController.(api.Controller)
```

[Factory.GetById(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.GetById) gets an implementation based on a given id.  If an implementaion has not been registered for the given id [Factory.GetById(...)](https://godoc.org/github.com/chrisehlen/knex#Factory.GetById) will return an error.
