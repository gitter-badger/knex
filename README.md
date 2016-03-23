# knex
Dependency injection framework for [Go](http://golang.org/).

## Installation and Docs

Install using `go get github.com/chrisehlen/knex`.

Full documentation is available at
https://godoc.org/github.com/chrisehlen/knex

## Usage

Knex is an easy to use api that uses Go tags to define relationships between components.

**Define component(s) that provides an implementation**

```go
type StringReaderImpl struct {
	Reader `provide:"resource"`
}
func (self *StringReaderImpl) Inject() error {return nil}
func (self *StringReaderImpl) Read() (string, error) {...}
```

StringReaderImpl is an implementation of the Reader interface that uses a constant string as input. StringReaderImpl has no requires so it's Inject() method has zero arguments.

```go
type UpperCaseFilterImpl struct {
	Filter `provide:"resource"`
}
func (self *UpperCaseFilterImpl) Inject() error {return nil}
func (self *UpperCaseFilterImpl) Do(message string) (string, error) {...}
```

UpperCaseFilterImpl is an implementation of the Filter interface that converts all characters of the message to upper case. UpperCaseFilterImpl has no requires so it's Inject() method has zero arguments.

```go
type PigLatinFilterImpl struct {
	Filter `provide:"resource"`
}
func (self *PigLatinFilterImpl) Inject() error { return nil }
func (self *PigLatinFilterImpl) Do(message string) (string, error) {...}
```

PigLatinFilterImpl is an implementation of the Filter interface that converts each word, of the message, to its Pig Latin equivalent. PigLatinFilterImpl has no requires so it's Inject() method has zero arguments.

```go
type ConsoleWriterImpl struct {
	Writer `provide:"resource"`
}
func (self *ConsoleWriterImpl) Inject() error {return nil}
func (self *ConsoleWriterImpl) Write(message string) error {...}
```

ConsoleWriterImpl is an implementation of the Filter interface that sends all input to standard output. ConsoleWriterImpl has no requires so it's Inject() method has zero arguments.

**Define component(s) that require implementation(s)**

```go
type ControllerImpl struct {
	Controller `provide:"resource"`
	reader     Reader   `require:"true"`
	filters    []Filter `require:"false"`
	writer     Writer   `require:"true"`
}
func (self *ControllerImpl) Inject(reader Reader, filters []Filter, writer Writer) error {
	self.reader = reader
	self.filters = filters
	self.writer = writer
	return nil
}
func (self *ControllerImpl) ReadMessage() (string, error) {...}
func (self *ControllerImpl) ApplyFilters(string) (string, error) {...}
func (self *ControllerImpl) WriteMessage(string) error {...}
```

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

**Parent factories**

```go
knex.DefaultFactory.AddParent(spiFactory)
```

**Get an implementation by type**

**Get all implementations by type**

**Get an implementation by id**

**Graph scope**

**Factory scope**
