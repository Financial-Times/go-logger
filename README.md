# go-logger
Logging library in GO, used for structural logging inside UPP (mostly required for monitoring publish events). It is based on the [Logrus](https://github.com/sirupsen/logrus) implementation.

### Initialization
When working with this logger library, please make sure you use one of the init methods first (otherwise the logging will fail):
- `InitLogger` - requires a serviceName and a logLevel
- `InitDefaultLogger` - requires only the serviceName as a parameter

Note: You can still create your own standard logger by using the `NewLogger` function (check [Logrus](https://github.com/sirupsen/logrus/blob/master/logger.go#L69) for more details).


### Logging a Monitoring Event
The library is Logrus compatible, but it includes a few default fields, 
which help facilitate the monitoring of key application events (`@time`, `transaction_id` and `service_name`).
You can add a monitoring event to a log entry by using the following method:
- `WithMonitoringEvent` - with `transaction_id`, `eventName` and `contentType` as parameters. 
A `monitoring_event=true` field will also be added to the entry. 
This message will be picked up by the monitoring services and dashboards.

### Adding additional fields to the Entry

Beside the With... fields offered by the original Logrus Entry, the following methods can be used:
- `WithTransactionID`, to add a transaction ID to the log entry;
- `WithUUID`, to add a UUID to the log entry;
- `WithTime`, to set a custom time of the logging entry (this can be used to influence Splunk log time); 
- `WithValidFlag` to mark if a message received by an application is valid or not. 
Invalid messages will be ignored by some of the monitoring statistics (SLAs).

### Actual Logging

Use Logrus' default log methods, like Info, Warn, Error and others.

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
    "github.com/Financial-Times/go-logger/test"
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