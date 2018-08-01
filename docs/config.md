The major config sections are the `blacklist` array, and the `[[aggregation]]`, `[[rewriter]]` and `[[route]]` entries.


You can also create routes, populate the blacklist, etc via the `init` config array using the same commands as the telnet interface, detailed below.

# Blacklist

example:
```
blacklist = [
  'prefix collectd.localhost',
  'regex ^foo\..*\.cpu+'
]
```

# Aggregators

Examples:
```
[[aggregation]]
# aggregate timer metrics with sums
function = 'sum'
prefix = ''
substr = ''
regex = '^stats\.timers\.(app|proxy|static)[0-9]+\.requests\.(.*)'
format = 'stats.timers._sum_$1.requests.$2'
interval = 10
wait = 20

[[aggregation]]
# aggregate timer metrics with averages
function = 'avg'
prefix = ''
substr = 'requests'
regex = '^stats\.timers\.(app|proxy|static)[0-9]+\.requests\.(.*)'
format = 'stats.timers._avg_$1.requests.$2'
interval = 5
wait = 10
cache = true

[[aggregation]]
# aggregate timer metrics with percentiles
# computes p25, p50, p75, p90, p95, p99 and writes each as a single metric.
# pxx gets appended to the corresponding metric path.
function = 'percentiles'
prefix = ''
substr = 'requests'
regex = '^stats\.timers\.(app|proxy|static)[0-9]+\.requests\.(.*)'
format = 'stats.timers.$1.requests.$2'
interval = 5
wait = 10
dropRaw = false
```

# Rewriters

example:
```
[[rewriter]]
# rewrite all instances of testold to testnew
old = 'testold'
new = 'testnew'
not = ''
max = -1
```

# Routes

## carbon route

setting        | mandatory | values                                        | default | description 
---------------|-----------|-----------------------------------------------|---------|------------
key            |     Y     | string                                        | N/A     |
type           |     Y     | sendAllMatch/sendFirstMatch/consistentHashing | N/A     | send to all destinations vs first matching destination vs distribute via consistent hashing
prefix         |     N     | string                                        | ""      |
sub            |     N     | string                                        | ""      |
regex          |     N     | string                                        | ""      |
replicationFactor  |     N     | int                                           | 1       | only for consistentHashing, defines the total number of destinations a matching point will be written to. It must be `0 < replicationFactor < number of destinations`

### Examples

```
[[route]]
# a plain carbon route that sends all data to the specified carbon (graphite) server
key = 'carbon-default'
type = 'sendAllMatch'
# prefix = ''
# substr = ''
# regex = ''
destinations = [
  '127.0.0.1:2003 spool=true pickle=false'
]

[[route]]
# all metrics with '=' in them are metrics2.0, send to carbon-tagger service
key = 'carbon-tagger'
type = 'sendAllMatch'
substr = '='
destinations = [
  '127.0.0.1:2006'
]

[[route]]
# send to the first carbon destination that matches the metric
key = 'analytics'
type = 'sendFirstMatch'
regex = '(Err/s|wait_time|logger)'
destinations = [
  'graphite.prod:2003 prefix=prod. spool=true pickle=true',
  'graphite.staging:2003 prefix=staging. spool=true pickle=true'
]

[[route]]
# a carbon route that sends to multiple carbon-cache servers
key = 'carbon-default'
type = 'consistentHashing'
prefix = 'consistentHashing'
# substr = ''
# regex = ''
 replicationFactor = 2
 destinations = [
  'carbon-cache1:2003 spool=true pickle=false',
  'carbon-cache2:2003 spool=true pickle=false',
  'carbon-cache3:2003 spool=true pickle=false'
]
```

## carbon destination

setting              | mandatory | values        | default | description 
---------------------|-----------|---------------|---------|------------
addr                 |     Y     |  string       | N/A     |
prefix               |     N     |  string       | ""      |
sub                  |     N     |  string       | ""      |
regex                |     N     |  string       | ""      |
flush                |     N     |  int (ms)     | 1000    | flush interval
reconn               |     N     |  int (ms)     | 10k     | reconnection interval
pickle               |     N     |  true/false   | false   | pickle output format instead of the default text protocol
spool                |     N     |  true/false   | false   | disk spooling
connbuf              |     N     |  int          | 30k     | connection buffer (how many metrics can be queued, not written into network conn)
iobuf                |     N     |  int (bytes)  | 2M      | buffered io connection buffer
spoolbuf             |     N     |  int          | 10k     | num of metrics to buffer across disk-write stalls. practically, tune this to number of metrics in a second
spoolmaxbytesperfile |     N     |  int          | 200MiB  | max filesize for spool files
spoolsyncevery       |     N     |  int          | 10k     | sync spool to disk every this many metrics
spoolsyncperiod      |     N     |  int  (ms)    | 1000    | sync spool to disk every this many milliseconds
spoolsleep           |     N     |  int (micros) | 500     | sleep this many microseconds(!) in between ingests from bulkdata/redo buffers into spool
unspoolsleep         |     N     |  int (micros) | 10      | sleep this many microseconds(!) in between reads from the spool, when replaying spooled data

## grafanaNet route

setting        | mandatory | values      | default | description 
---------------|-----------|-------------|---------|------------
key            |     Y     |  string     | N/A     |
addr           |     Y     |  string     | N/A     |
apiKey         |     Y     |  string     | N/A     |
schemasFile    |     Y     |  string     | N/A     |
prefix         |     N     |  string     | ""      |
sub            |     N     |  string     | ""      |
regex          |     N     |  string     | ""      |
sslverify      |     N     |  true/false | true    |
spool          |     N     |  true/false | false   | ** disk spooling. not implemented yet **
blocking       |     N     |  true/false | false   | if false, full buffer drops data. if true, full buffer puts backpressure on the table, possibly affecting ingestion and other routes
concurrency    |     N     |  int        | 100     | number of concurrent connections to ingestion endpoint
bufSize        |     N     |  int        | 10M     | buffer size. assume +- 100B per message, so 10M is about 1GB of RAM
flushMaxNum    |     N     |  int        | 5000    | max number of metrics to buffer before triggering flush
flushMaxWait   |     N     |  int (ms)   | 500     | max time to buffer before triggering flush
timeout        |     N     |  int (ms)   | 10000   | abort and retry requests to api gateway if takes longer than this.
orgId          |     N     |  int        | 1       |

### Example
example route for https://grafana.com/cloud/metrics

```
[[route]]
key = 'grafanaNet'
type = 'grafanaNet'
addr = 'your-base-url/metrics'
apikey = 'your-grafana.net-api-key'
schemasFile = 'examples/storage-schemas.conf'
# require the destination address to have a valid SSL certificate
#sslverify=true
# Number of metrics to spool in in-memory buffer. since a message is typically around 100B this is 1GB
#bufSize=10000000
# When the in-memory buffer reaches capacity we can either "block" ingestion or "drop" metrics.
#blocking=false
# maximum number of metrics to send in a single batch to grafanaCloud
#flushMaxNum=5000
# maximum time in ms that metrics can wait in buffers before being sent.
#flushMaxWait=500
# time in ms, before giving up trying to send a batch of data to grafanaCloud.
#timeout=10000
# number of concurrent connections to use for sending data.
concurrency=100
```

## kafkaMdm route

setting        | mandatory | values      | default | description 
---------------|-----------|-------------|---------|------------
key            |     Y     |  string     | N/A     |
brokers        |     Y     |  []string   | N/A     |
topic          |     Y     |  string     | N/A     |
codec          |     Y     |  string     | N/A     |
partitionBy    |     Y     |  string     | N/A     |
schemasFile    |     Y     |  string     | N/A     |
prefix         |     N     |  string     | ""      |
sub            |     N     |  string     | ""      |
regex          |     N     |  string     | ""      |
blocking       |     N     |  true/false | false   | if false, full buffer drops data. if true, full buffer puts backpressure on the table, possibly affecting ingestion and other routes
bufSize        |     N     |  int        | 10M     | buffer size. assume +- 100B per message, so 10M is about 1GB of RAM
flushMaxNum    |     N     |  int        | 10k     | max number of metrics to buffer before triggering flush
flushMaxWait   |     N     |  int (ms)   | 500     | max time to buffer before triggering flush
timeout        |     N     |  int (ms)   | 2000    |
orgId          |     N     |  int        | 1       |

## Google PubSub route

setting        | mandatory | values      | default       | description 
---------------|-----------|-------------|---------------|------------
key            |     Y     |  string     | N/A           |
project        |     Y     |  string     | N/A           | Google Cloud Project containing the topic
topic          |     Y     |  string     | N/A           | The Google PubSub topic to publish metrics onto
codec          |     N     |  string     | gzip          | Optional compression codec to compress published messages. Valid: "none" (no compression), "gzip" (default)
prefix         |     N     |  string     | ""            |
sub            |     N     |  string     | ""            |
regex          |     N     |  string     | ""            |
blocking       |     N     |  true/false | false         | if false, full buffer drops data. if true, full buffer puts backpressure on the table, possibly affecting ingestion and other routes
bufSize        |     N     |  int        | 10M           | buffer size. assume +- 100B per message, so 10M is about 1GB of RAM
flushMaxSize   |     N     |  int        | (10MB - 4KB)  | max message size before triggering flush. The size is before any compression is calculated. PubSub message limit is 10MB but this can be higher if using compression.
flushMaxWait   |     N     |  int (ms)   | 1000          | max time to buffer before triggering flush

## Cloudwatch

setting          | mandatory | values      | default       | description 
-----------------|-----------|-------------|---------------|------------
key              |     Y     |  string     | N/A           |
profile          |     N     |  string     | N/A           | The Amazon CloudWatch profile to use. For local development needed only. In the cloud, the profile is known.
region           |     Y     |  string     | N/A           | The Amazon Geo region to send metrics into.
prefix           |     N     |  string     | ""            |
sub              |     N     |  string     | ""            |
regex            |     N     |  string     | ""            |
blocking         |     N     |  true/false | false         | if false, full buffer drops data. if true, full buffer puts backpressure on the table, possibly affecting ingestion and other routes.
bufSize          |     N     |  int        | 10M           | buffer size. assume +- 100B per message, so 10M is about 1GB of RAM.
flushMaxSize     |     N     |  int        | 20            | max MetricDatum objects in slice before triggering flush. 20 is currently the CloudWatch max.
flushMaxWait     |     N     |  int (ms)   | 10000         | max time to buffer before triggering flush.
storageResolution|     N     |  int (s)    | 60            | either 1 or 60. 1 = high resolution metrics, 60 = regular metrics.
namespace        |     Y     |  string     | N/A           | The CloudWatch namespace to publish metrics into.
dimensions       |     N     |  []string   | N/A           | CloudWatch Dimensions that are attached to all metrics.

#### Examples

CloudWatch namespace and dimensions are meant to be set on a metric basis. However, plain carbon metrics do
not include this information. Therefore, we fix them in the config for now.

```
[[route]]
key = 'cloudWatch'
type = 'cloudWatch'
profile = 'for-development'
region = 'eu-central-1'
namespace = 'MyNamespace'
dimensions = [
   ['myDimension', 'myDimensionVal'],
]
storageResolution = 1
```

## Imperatives

imperatives can be used in two places:
1) via the TCP command interface
2) in the init.cmds config setting (as single quoted strings). However, **the structured settings as shown above are the better and recommended way**
That said, here are some examples:

```
# a plain carbon route that sends all data to the specified carbon (graphite) server (note the double space separating route options from destination options)
#'addRoute sendAllMatch carbon-default  your-graphite-server:2003 spool=true pickle=false',

# example route for https://grafana.com/cloud/metrics (note the double space separating route options from destination options)
#'addRoute grafanaNet grafanaNet  your-base-url/metrics your-grafana.net-api-key /path/to/storage-schemas.conf',

# ignore hosts that don't set their hostname properly via prefix match
#'addBlack prefix collectd.localhost',

# ignore foo.<anything>.cpu.... via regex match
#'addBlack regex ^foo\..*\.cpu+',

# aggregate timer metrics with sums
#'addAgg sum regex=^stats\.timers\.(app|proxy|static)[0-9]+\.requests\.(.*) stats.timers._sum_$1.requests.$2 10 20 cache=true',

# aggregate timer metrics with averages
#'addAgg avg regex=^stats\.timers\.(app|proxy|static)[0-9]+\.requests\.(.*) sub=requests stats.timers._avg_$1.requests.$2 5 10 dropRaw=false',

# all metrics with '=' in them are metrics2.0, send to carbon-tagger service (note the double space separating route options from destination options)
#'addRoute sendAllMatch carbon-tagger sub==  127.0.0.1:2006',

# send to the first carbon destination that matches the metric (note the double spaces between destinations and route)
#'addRoute sendFirstMatch analytics regex=(Err/s|wait_time|logger)  graphite.prod:2003 prefix=prod. spool=true pickle=true  graphite.staging:2003 prefix=staging. spool=true pickle=true'

# send to the Google PubSub topic named "graphite-ingest" in the "myproject" project
#'addRoute pubsub pubsub  myproject graphite-ingest'
```
