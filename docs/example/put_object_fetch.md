# PUT Object - Fetch

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

Then set the input parameters used by the PutObject method (core parameter: XQSFetchSource).

```go
	// Fetch source looks like this: "protocol://host[:port]/[path]"
	sourceLink := "https://www.shanhe.com/static/assets/images/icons/common/footer_logo.svg"
	input := &service.PutObjectInput{
		XQSFetchSource: &sourceLink,
	}
```

Please note that not all fields in PutObjectInput required to be set. For details, please refer to [Official API Documentation](https://docsv3.shanhe.com/ois/api/object/fetch).

Then call the PutObject method to fetch object. objectKey Sets the filepath after put (in the current bucket).

```go
	objectKey := "file-fetched/the_file_fetched.svg"
	if output, err := bucketService.PutObject(objectKey, input); err != nil {
		fmt.Printf("Fetch object from source link(%s) to target path(%s) failed with given error: %s\n", sourceLink, objectKey, err)
	} else {
		fmt.Printf("The status code expected: 201(actually: %d)\n", *output.StatusCode)
	}
```
