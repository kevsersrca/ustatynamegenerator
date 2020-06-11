# USTATY NAME GENERATOR

A random name generator for ustaty server name suggestions. 

## Installation

```shell script
git clone https://github.com/ustaty/ustatynamegenerator.git
```

## Usage
```shell script
package main

import (
	"fmt"
	"math/rand"
	"time"

	gen "github.com/ustaty/ustatynamegenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(gen.GetRandomName(0))
}

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[APACHE](LICENSE)