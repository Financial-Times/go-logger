# go-logger

Logging library in Go, used for structural logging inside UPP. It is wrapper of [logrus](https://github.com/sirupsen/logrus) library for structrural logging in Go.
As wrapper of logrus logger, the UPP logger provides the same functionality and extends it with a few UPP specific methods. It also enforces the aggreed logging format in UPP -
JSON formatted logs and each of the log entries include the service name, log level and time if available.

As logrus logger implements the standard Go log interface, the UPP logger implements it as well.
So the UPP logger can be used where standard Go log is required. 

UPP logger shares the same log levels as logrus - debug, info, warning, error, fatal, panic.

### Initialization
When working with this logger library, please use one of the init methods:
- `NewUPPLogger` - requires a serviceName and a logLevel as parameters. Also there is additional optional parameter - 
configuration for the names of the field keys logged by the UPP logger methods. 
- `NewUPPInfoLogger` - requires only the serviceName as a parameter. Initializes logger with log level info. 
Also there is additional optional parameter - 
configuration for the names of the field keys logged by the UPP logger methods. 
- `NewUnstructuredLogger` - returns UPP logger but without enforced structural logging format.

Please note that using package level logger by only importing the library (supported in v1 of this library) is no longer available.

### Logging with the UPP logger
UPP logger supports structural logging as logrus supports it. Please take a look at [logging fields](https://github.com/sirupsen/logrus#fields)
as logrus method for structural logging. UPP logrus also implements `WithField` and `WithFields` methods.
As logrus UPPLogger also supports chaining of the methods that add logging fields.

For producing actual log, use the default log methods, like Info, Warn, Error and others.

### Adding additional methods to the Entry and logger

Beside the With... fields offered by the original logrus Entry and logger, the following methods can be used:
- `WithTransactionID`, to add a transaction ID to the log entry;
- `WithUUID`, to add a UUID to the log entry;
- `WithTime`, to set a custom time of the logging entry (this can be used to influence Splunk log time); 
- `WithValidFlag` to mark if a message received by an application is valid or not. 
Invalid messages will be ignored by some of the monitoring statistics (SLAs).


### Logging events
The library includes methods which help facilitate the monitoring of key application events.

- You can add a monitoring event to a log entry by using the following method:
`WithMonitoringEvent` - with transaction ID, event name and content type as parameters. 
A `monitoring_event=true` field will also be added to the entry. 
This message will be picked up by the monitoring services and dashboards.
- You can add an event with category and message by using: `WithCategorisedEvent` - with event name,
event category and event message as parameters. Using this method we are also able to produce log
with particular structure easy to be picked up and parsed by a monitoring tool.

### Examples

A monitoring log for a successful publish, with validation flag, can look like this:

```
logger.WithMonitoringEvent("Map", tid, "Annotations")
      .WithUUID(uuid)
      .WithValidFlag(true)
      .Info("Successfully mapped")
```

A monitoring log for a failed publish would log it as an error:
```
logger.WithMonitoringEvent("Map", tid, "Annotations")
      .WithUUID(uuid)
      .WithValidFlag(true)
      .WithError(err)
      .Error("Error decoding body")
```

### Test Package

The `test` package has been introduced to check through unit tests that the application is logging relevant events 
properly. The example below shows how to check that an application is logging a specific monitoring event:
```
import (
    ...
    logTest "github.com/Financial-Times/go-logger/test"
    ...
)

func TestSomething(t *testing.T) {
    hook := logTest.NewTestHook("serviceName")
    ...
    Something()
    ...
    entry := hook.LastEntry()
    test.Assert(t, entry).HasMonitoringEvent("Map", "tid_test", "annotations").HasValidFlag(true)
}

```

### Installation

```
    GO111MODULE=on go get github.com/Financial-Times/go-logger@v2
    cd $GOPATH/src/github.com/Financial-Times/go-logger/
    GO111MODULE=on go build -mod=readonly
```