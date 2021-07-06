# Service Initialization

First, we need to initialize a OIS service to call the services provided by OIS.

```go
import (
	"github.com/shanhe-nsccjn/ois-sdk-go/v4/config"
	"github.com/shanhe-nsccjn/ois-sdk-go/v4/service"
)

const accessKeyId = "YOUR-ACCESS-KEY-ID"
const secretAccessKey = "YOUR--SECRET-ACCESS-KEY"

var conf, _ = config.New(accessKeyId, secretAccessKey)
var oIs, _ = service.Init(conf)
var bucketService, _ = oIs.Bucket("your-bucket-name", "zone-name")
```

The object that appears in the above code:
- The `conf` object carries the user's authentication information and configuration.
- The `oIs` object is used to operate the OIS object storage service, which is used to call all Service level APIs or to create a specified Bucket object to call Bucket and Object level APIs.
- The `bucketService` object is bound to the specified bucket and provides a series of object storage operations for the bucket.