package atomx

import (
	"os"
	"testing"
)

var (
	client = New()
)

func init() {
	client.ApiURL = "https://sandbox-api.atomx.com/v2/"

	if os.Getenv("EMAIL") == "" || os.Getenv("PASSWORD") == "" {
		panic("EMAIL and PASSWORD are required")
	}
}

func TestClientLogin(t *testing.T) {
	if err := client.Login(os.Getenv("EMAIL"), os.Getenv("PASSWORD")); err != nil {
		t.Fatal(err)
	}
}

func TestPut(t *testing.T) {
	site := Site{
		Id: 8299,
	}

	if err := client.Get(&site); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)

	site.DomainId += 1

	if err := client.Put(&site); err != nil {
		t.Fatal(err)
	}

	expected := site.DomainId

	if err := client.Get(&site); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)

	if site.DomainId != expected {
		t.Fatalf("got %d expected %d", site.DomainId, expected)
	}
}
