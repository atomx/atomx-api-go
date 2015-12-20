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
		Id: 5,
	}

	if err := client.Get(&site, nil); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)

	site.DomainId += 1

	if err := client.Put(&site); err != nil {
		t.Fatal(err)
	}

	expected := site.DomainId

	if err := client.Get(&site, nil); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)

	if site.DomainId != expected {
		t.Fatalf("got %d expected %d", site.DomainId, expected)
	}
}

func TestList(t *testing.T) {
	domains := Domains{
		List: List{
			Offset: 0,
			Limit:  10,
		},
	}

	if err := client.List(&domains, nil); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", domains)

	if len(domains.Domains) != 10 {
		t.Fatal("expected 10 domains")
	}

	d := domains.Domains[0]

	domains.Offset = 10
	domains.Limit = 2

	if err := client.List(&domains, nil); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", domains)

	if len(domains.Domains) != 2 {
		t.Fatal("expected 2 domains")
	}

	if domains.Domains[0].Id == d.Id {
		t.Fatal("d changed")
	}
}

func TestDomain(t *testing.T) {
	domain := Domain{
		Id: 777835,
	}

	if err := client.Get(&domain, &Options{
		Expand: []string{"attributes"},
	}); err != nil {
		t.Fatal(err)
	} else if !domain.Attributes.Has(7) {
		domain.Attributes = append(domain.Attributes, DomainAttribute{
			Id: 7,
		})

		if err := client.Put(&domain); err != nil {
			t.Fatal(err)
		}
	}
}
