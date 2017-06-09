# retry-go

* Retrying made simple and easy for golang.

### Usage

```go
package main

import (
  "time"

  retry "github.com/hellofresh/retry-go"
)

func main() {
  attempts := 3
  sleepTime := time.Second*2
  if err := retry.Do(func() error {
    if err := work(); err != nil {
      return err
    }
    return nil
  }, attempts, sleepTime); err != nil {
    // Retry failed
  }
}
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request

## Badges

[![Build Status](https://travis-ci.org/hellofresh/retry-go.svg?branch=master)](https://travis-ci.org/hellofresh/retry-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/hellofresh/retry-go)](https://goreportcard.com/report/github.com/hellofresh/retry-go)
[![Go Doc](https://godoc.org/github.com/hellofresh/retry-go?status.svg)](https://godoc.org/github.com/hellofresh/retry-go)

---

> GitHub [@hellofresh](https://github.com/hellofresh) &nbsp;&middot;&nbsp;
> Medium [@engineering.hellofresh](https://engineering.hellofresh.com)
