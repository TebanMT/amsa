module github.com/TebanMT/amsa

go 1.21.5

require (
	github.com/TebanMT/amsa/infrastructure/cognito v0.0.0
	github.com/TebanMT/amsa/infrastructure/db v0.0.0
	github.com/TebanMT/amsa/infrastructure/authentication v0.0.0
	github.com/aws/aws-cdk-go/awscdk/v2 v2.118.0
	github.com/aws/constructs-go/constructs/v10 v10.3.0
	github.com/aws/jsii-runtime-go v1.93.0
	github.com/golang-jwt/jwt/v4 v4.5.0
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
)

require (
	github.com/Masterminds/semver/v3 v3.2.1 // indirect
	github.com/aws/aws-sdk-go v1.49.20 // indirect
	github.com/cdklabs/awscdk-asset-awscli-go/awscliv1/v2 v2.2.201 // indirect
	github.com/cdklabs/awscdk-asset-kubectl-go/kubectlv20/v2 v2.1.2 // indirect
	github.com/cdklabs/awscdk-asset-node-proxy-agent-go/nodeproxyagentv6/v2 v2.0.1 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/yuin/goldmark v1.4.13 // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
	gorm.io/gorm v1.25.5 // indirect
)

replace github.com/TebanMT/amsa => ./

replace github.com/TebanMT/amsa/infrastructure/db => ./infrastructure/db
replace github.com/TebanMT/amsa/infrastructure/authentication => ./infrastructure/authentication

replace github.com/TebanMT/amsa/infrastructure/cognito => ./infrastructure/cognito
