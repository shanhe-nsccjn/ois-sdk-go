# GET Bucket External Mirror

## 代码片段

使用您的 AccessKeyID 和 SecretAccessKey 初始化 Qingstor 对象。

```go
import (
	"github.com/shanhe-nsccjn/ois-sdk-go/v4/config"
	"github.com/shanhe-nsccjn/ois-sdk-go/v4/service"
)

var conf, _ = config.New("YOUR-ACCESS-KEY-ID", "YOUR--SECRET-ACCESS-KEY")
var oIs, _ = service.Init(conf)
```

然后根据要操作的 bucket 信息（zone, bucket name）来初始化 Bucket。

```go
	bucketName := "your-bucket-name"
	zoneName := "pek3b"
	bucketService, _ := oIs.Bucket(bucketName, zoneName)
```

然后您可以 GET Bucket External Mirror

```go
	if output, err := bucketService.GetExternalMirror(); err != nil {
		fmt.Printf("Get external mirror of bucket(name: %s) failed with given error: %s\n", bucketName, err)
	} else {
		b, _ := json.Marshal(output.SourceSite)
		fmt.Println("The external mirror of this bucket: ", string(b))
	}
```