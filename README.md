# go-logger
Logging library in GO, used for structural logging inside UPP (mostly required for publish monitoring). It is based on the [Logrus](https://github.com/sirupsen/logrus) implementation.

When working with this logger, please make sure you use one of the init methods first (otherwise the logging will fail):
- `InitLogger` - requires a serviceName and a logLevel
- `InitDefaultLogger` - requires only the serviceName as a parameter

Note: You can still create your own standard logger by using the `NewLogger` function (check [Logrus](https://github.com/sirupsen/logrus/blob/master/logger.go#L69) for more details).

For Log entries serving the monitoring services, please use the monitoringEvent methods.
Currently we have 3 of them:
- `MonitoringEvent` - requires the `event` name (use any of your preference), `tid` (for `transaction id`), `contentType` (like `annotations`), and `message`. Logs at `INFO` level.
- `MonitoringEventWithUUID` - similar with the above, with a `uuid` field. Logs at `INFO` level.
- `MonitoringValidationEvent` - use this method for mappers, for validating the content type. The service logs at `INFO` level of validation was successful (`isValid=true`), otherwise it logs at `ERROR` level.

There are other available methods, which should make the structured logging (in json) easier.
