# GET Bucket External Mirror

## Code Snippet

Initialize the Qingstor object with your AccessKeyID and SecretAccessKey.

```go
import (
	"github.com/shanhe-nsccjn/ois-sdk-go/v4/config"
	"github.com/shanhe-nsccjn/ois-sdk-go/v4/service"
)

var conf, _ = config.New("YOUR-ACCESS-KEY-ID", "YOUR--SECRET-ACCESS-KEY")
var oIs, _ = service.Init(conf)
```

Initialize a Bucket object according to the bucket name you set for subsequent creation:

```go
bucketName := "your-bucket-name"
zoneName := "pek3b"
bucketService, _ := oIs.Bucket(bucketName, zoneName)
```

then you can GET Bucket External Mirror

```go
	if output, err := bucketService.GetExternalMirror(); err != nil {
		fmt.Printf("Get external mirror of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		b, _ := json.Marshal(output.SourceSite)
		fmt.Println("The external mirror of this bucket: ", string(b))
	}
```