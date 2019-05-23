## Overview

httpdig - allows to make DNS queries using Google's HTTPS DNS service. This repo was forked from [https://github.com/firstrow/httpdig](github.com/firstrow/httpdig) because that project appears abandoned.

## Install

``` bash
go get -u github.com/jgoodall/httpdig
```

## Usage

``` go
import (
       "fmt"
       "github.com/jgoodall/httpdig"
)

resp, _ := httpdig.Query("google.com", "NS")
fmt.Print(resp.Answer)
```

## Links
[RR Types](https://en.wikipedia.org/wiki/List_of_DNS_record_types)
