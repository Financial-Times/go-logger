# go-logger
Logging library in GO, used for structural logging inside UPP.

For Log entries serving the monitoring services, please use the monitoringEvent methods.
Currently we have 3 of them defined:
- `MonitoringEvent` - requires the `event` name (use any of your preference), `tid` (for `transaction id`), `contentType` (like `annotations`), and `message`. Logs at `INFO` level.
- `MonitoringEventWithUUID` - similar with the above, with a `uuid` field. Logs at `INFO` level.
- `MonitoringValidationEvent` - use this method for mappers, for validating the content type. The service logs at `INFO` level of validation was successful (`isValid=true`), otherwise it logs at `ERROR` level.

The service has other methods too, which should make the structured logging (in json) easier.
