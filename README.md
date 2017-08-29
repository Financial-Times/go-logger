# go-logger
Logging library in GO, used for structural logging inside UPP (mostly required for monitoring publish events). It is based on the [Logrus](https://github.com/sirupsen/logrus) implementation.

### Initialization
When working with this logger library, please make sure you use one of the init methods first (otherwise the logging will fail):
- `InitLogger` - requires a serviceName and a logLevel
- `InitDefaultLogger` - requires only the serviceName as a parameter

Note: You can still create your own standard logger by using the `NewLogger` function (check [Logrus](https://github.com/sirupsen/logrus/blob/master/logger.go#L69) for more details).


### Creating an Entry
The library maintains most of the Logrus's way of logging, but it adds some default fields (`@time`, `transaction_id` and `service_name`) for the created logging entries.
For entry creation you can use the following methods:
- `NewEntry` - with `transaction_id` as a parameter
- `NewMonitoringEntry` - with `transaction_id`, `eventName` and `contentType` as parameters. Beside these, a `monitoring_event=true` field will also be added to the entry. This message will be picked up by the monitoring services and dashboards.

### Adding additional fields to the Entry

Beside the With... fields offered by the original Logrus Entry, the `WithUUID`, `WithTime` and `WithValidFlag` methods can also be used.
The validation flag should mark whether the message is accepted as valid by the service it uses. Invalid messages will be ignored by some of the monitoring statistics (SLAs).

### Actual Logging

Use Logrus' default log methods, like Info, Warn, Error and others.

### Examples

A monitoring log for a successful publish, with validation flag, can look like this:

```
logger.NewMonitoringEntry("Map", tid, "Annotations")
      .WithUUID(uuid)
      .WithValidFlag(true)
      .Info("Successfully mapped")
```

A monitoring log for a failed publish would log it as an error:
```
logger.NewMonitoringEntry("Map", tid, "Annotations")
      .WithUUID(uuid)
      .WithValidFlag(true)
      .WithError(err)
      .Error("Error decoding body")
```
