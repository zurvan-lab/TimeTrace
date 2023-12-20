# TQL (Time Query Language)

Time trace is using a query language called TQL. Here is documentation and specification for TQL.  


# Commands and Operations

> [*] means implemented.

| Command   |      Action      |  Arguments |
|----------|:-------------|:------|
| CON * |  to make a connection and access to database | username - password |
| SET * |    make a new set   | set-name |
| SSET * | make a new subset | set-name - subset-name |
| PUSH * | push an element to a subset | set-name - subset-name - value-of-element - time(unix-timestamp) |
| GET * | get elements of a subset | set-name - subset-name - [last n elements (optional)] |
| CNTS * | returns count of sets |  |
| CNTSS * | returns count of subsets | set-name |
| CNTE * | returns count of elements | set-name - subset-name |
| CLN * | cleans all database sets (the sets themselves) | set-name - subset-name |
| CLNS * | cleans all sub-sets of a set | set-name |
| CLNSS * | cleans all elements of a subset | set-name - subset-name |
| DRPS * | drops a set | set-name |
| DRPSS * | drops a subset | set-name - subset-name |
| SS | takes a sanp-shot | file-name(example.ttrace) |


# A TQL Server Responses

| Message   |      reason      | 
|----------|:-------------|
| DONE | everything is ok |
| INVALID | invalid user and password to make a connection OR not enough args for a command |
| SNF | set is not found |
| SSNF | subset is not found |
| ENF | element(s) is not found |
| DATA in format of [protocol](../protocol/) |
