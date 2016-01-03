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
		ID: 5,
	}

	if err := client.Get(&site, nil); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)

	site.DomainID += 1

	if err := client.Put(&site, nil); err != nil {
		t.Fatal(err)
	}

	expected := site.DomainID

	if err := client.Get(&site, nil); err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", site)

	if site.DomainID != expected {
		t.Fatalf("got %d expected %d", site.DomainID, expected)
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

	if domains.Domains[0].ID == d.ID {
		t.Fatal("d changed")
	}
}

func TestDomain(t *testing.T) {
	domain := Domain{
		ID: 777835,
	}

	if err := client.Get(&domain, &Options{
		Expand: []string{"attributes"},
	}); err != nil {
		t.Fatal(err)
	} else if !domain.Attributes.Has(7) {
		domain.Attributes = append(domain.Attributes, DomainAttribute{
			ID: 7,
		})

		if err := client.Put(&domain, nil); err != nil {
			t.Fatal(err)
		}
	}
}
