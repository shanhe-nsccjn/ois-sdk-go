# 初始化服务

首先我们需要初始化一个 OIS Service 来调用 OIS 提供的各项服务。

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

上面代码中出现的对象：
- `conf` 对象承载了用户的认证信息及配置。
- `oIs` 对象用于操作 OIS 对象存储服务，用于调用所有 Service 级别的 API 或创建指定的 Bucket 对象来调用 Bucket 和 Object 级别的 API。
- `bucketService` 对象绑定了指定 bucket，提供一系列针对该 bucket 的对象存储操作。