# go-logger
Logging library in GO, used for structural logging inside UPP (mostly required for monitoring publish events). It is based on the [Logrus](https://github.com/sirupsen/logrus) implementation.

### Initialization
When working with this logger library, please make sure you use one of the init methods first (otherwise the logging will fail):
- `InitLogger` - requires a serviceName and a logLevel
- `InitDefaultLogger` - requires only the serviceName as a parameter

Note: You can still create your own standard logger by using the `NewLogger` function (check [Logrus](https://github.com/sirupsen/logrus/blob/master/logger.go#L69) for more details).


### Logging a Monitoring Event
The library maintains most of the Logrus's way of logging, but it adds some default fields 
(`@time`, `transaction_id` and `service_name`) to facilitate monitoring of relevant application events.
You can add a monitoring event to a log entry by using the following method:
- `WithMonitoringEvent` - with `transaction_id`, `eventName` and `contentType` as parameters. 
Beside these, a `monitoring_event=true` field will also be added to the entry. 
This message will be picked up by the monitoring services and dashboards.

### Adding additional fields to the Entry

Beside the With... fields offered by the original Logrus Entry, the following methods can be used:
- `WithTransactionID`, to add a transaction ID to the log entry;
- `WithUUID`, To add a UUID to the log entry;
- `WithTime`, to set a custom time of the logging entry; 
- `WithValidFlag` to mark the if a message processed by an application is valid. 
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
