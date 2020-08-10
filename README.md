# Go Properties

<img align="right" src="https://img.shields.io/badge/License-MIT-blue.svg">

### Summary
 <a href="https://www.linkedin.com/in/sven-schaper/" style="text-align: right"><img align="right" src="https://www.lime-anchor.com/img/connect.png" height="250"></a><br>

**Go Properties** is a property file loader for golang projects without any additional dependency except golang core components.
At the moment the property loader is able to open configuration files and read key value pairs. More complex structures like arrays or nested objects are not supported yet but will be extended soon. The property loader provides a method to check a named property inside a file. The property file itlsef will be initialized as a singelton.

### Set Up

Download the project
```bash
go get https://github.com/svenschaper/goproperties
```

Import the package into your go File
```golang
import "github.com/svenschaper/goproperties"
```


### How to use?

Add the following to your go files
```golang
var prop *properties.Propertie

func init() {
	properties.SetPropertiePath("path/to/your/config.yml")
    p, err := properties.LoadProperty()
    if err != nil {
        fmt.Errorf(err.Error())
    }
	prop = p
}

func yourFunction(){
    value := prop.GetProperty("yourkey")
}

```




<br>
<br>


### Need custom development?

<a href="https://lime-anchor.com"><img align="right" src="https://www.lime-anchor.com/img/gint.png" height="50"></a>

**Cloud Computing**
* Infrastructure automation based on ansible
* Custom solutions based on AWS
* Custom solutions based on Google Cloud
* Custom solutions based on Azure Cloud

**Integration**
* API Management and Design
* Lightning fast Microservices based on Golang



# License

MIT licensed. In short -> Have fun with it!
