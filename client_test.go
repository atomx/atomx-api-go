package atomx

import (
	"os"
	"testing"
)

var (
	client = New()

	site *Site
)

func init() {
	client.ApiUrl = "https://sandbox-api.atomx.com/v2/"

	if os.Getenv("EMAIL") == "" || os.Getenv("PASSWORD") == "" {
		panic("EMAIL and PASSWORD are required")
	}
}

func TestClientLogin(t *testing.T) {
	if err := client.Login(os.Getenv("EMAIL"), os.Getenv("PASSWORD")); err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	var err error
	site, err = client.Site(8299)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)
}

func TestPut(t *testing.T) {
	if site == nil {
		t.SkipNow()
	}

	site.DomainId += 1

	if err := client.Put(site); err != nil {
		t.Fatal(err)
	}
}
