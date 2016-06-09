atomx-api-go
============

[![GoDoc](https://godoc.org/github.com/atomx/atomx-api-go?status.svg)](https://godoc.org/github.com/atomx/atomx-api-go)

atomx-api-go is a [Go](http://golang.org/) client for the [Atomx](https://www.atomx.com/) [API](https://wiki.atomx.com/api)

Example
-------

List all your creatives:

```go
client := atomx.New()

if err := client.Login("email", "password"); err != nil {
	panic(err)
}

offset := int64(0)

for {
	creatives := atomx.CreativesList{
		List: atomx.List{
			Sort:   "id.asc",
			Offset: offset,
			Limit:  100,
		},
	}

	if err := client.List(&creatives, &atomx.Options{
		Extra: []string{"quickstats.yesterday.impressions>=100", "id>=1"},
	}); err != nil {
		return nil, err
	}
	
	if len(creatives.Creatives) == 0 {
		break
	}

	for _, creative := range creatives.Creatives {
		println(creative)
	}

	offset += creatives.Limit
}
```
