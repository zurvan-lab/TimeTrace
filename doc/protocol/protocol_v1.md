# Protocol V1

This document will define the time trace protocol, which you need to know about (beside the [TQL](../TQL/)) to be able to make an client or driver for time trace.

# Requests

All of the data requesting from a time-trace database, will follow the [TQL](../TQL/) specification.

# Receiving Data

The TQL `GET` request data format:

```
<EOS-3bytes><value><S-1bytes><time><SOS-3bytes><protocolVersion-1bytes><payloadSize-8bytes>
```

EOS: end of sequence
SOS: start of sequence
S: separator (between each time and value)

> NOTE: value and time part can be repeated more than once depend on your request.
