materialize
===========

Materialize go package provide simple API to create UI using component builder in source code

Quick start
-----------

You may understand base concept of using materialize builder by next example:

```
package main

import "log"

import "github.com/vit1251/materialize"

func main() {

	doc, _ := materialize.NewDocument()
	body, _ := doc.Pack();

	log.Printf("body = %s", body)
}
```
