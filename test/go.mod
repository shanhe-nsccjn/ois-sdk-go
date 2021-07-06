module github.com/shanhe-nsccjn/ois-sdk-go/test

go 1.12

require (
	github.com/DATA-DOG/godog v0.7.13
	github.com/shanhe-nsccjn/log v0.0.0-20200804082313-615256cccabc
	github.com/shanhe-nsccjn/ois-sdk-go/v4 v4.0.0
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/shanhe-nsccjn/ois-sdk-go/v4 => ../
