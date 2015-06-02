package shippable

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_accounts.json")
		fmt.Fprint(w, string(rawResponse))
	})
	accounts, _, err := client.Accounts.GetAccounts()
	if err != nil {
		t.Errorf("Accounts.GetAccounts returned error %v", err)
	}

	if len(*accounts) != 2 {
		t.Errorf("Accounts.GetAccounts accounts returned; expected %d, actual %d", 2, len(*accounts))
	}

	wantedAccounts := []string{"foo", "bar"}
	if !reflect.DeepEqual(*accounts, wantedAccounts) {
		t.Errorf("Actual = %+v, Expected = %+v", *accounts, wantedAccounts)
	}
}

func TestGetAccount(t *testing.T) {
	setup()
	defer teardown()

	wantedAccountID := "640e74943999391400416qr0"
	mux.HandleFunc("/accounts/"+wantedAccountID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_accounts_640e74943999391400416qr0.json")
		fmt.Fprint(w, string(rawResponse))
	})
	account, _, err := client.Accounts.GetAccount(wantedAccountID)
	if err != nil {
		t.Errorf("Accounts.GetAccount returned error %v", err)
	}

	if *account.ID != wantedAccountID {
		t.Errorf("Accounts.GetAccount expected: %s, actual %s", wantedAccountID, *account.ID)
	}
}

func TestGetAccountIdentities(t *testing.T) {
	setup()
	defer teardown()

	wantedAccountID := "640e74943999391400416qr0"

	mux.HandleFunc("/accounts/"+wantedAccountID+"/identities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		rawResponse, _ := ioutil.ReadFile("./mocks/response_get_accounts_640e74943999391400416qr0_identities.json")
		fmt.Fprint(w, string(rawResponse))
	})
	identities, _, err := client.Accounts.GetAccountIdentities(wantedAccountID)
	if err != nil {
		t.Errorf("Accounts.GetAccountIdentities returned error %v", err)
	}

	if len(*identities) != 2 {
		t.Errorf("Accounts.GetAccountIdentities identities returned; expected %d, actual %d", 2, len(*identities))
	}

	wantedIdentities := []string{"foo", "bar"}
	if !reflect.DeepEqual(*identities, wantedIdentities) {
		t.Errorf("Actual = %+v, Expected = %+v", *identities, wantedIdentities)
	}
}

func TestDeleteAccount(t *testing.T) {
	setup()
	defer teardown()

	wantedAccountID := "640e74943999391400416qr0"

	mux.HandleFunc("/accounts/"+wantedAccountID, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, "")
	})

	ok, _, err := client.Accounts.DeleteAccount(wantedAccountID)
	if err != nil {
		t.Errorf("Accounts.DeleteAccount returned error %v", err)
	}
	if !ok {
		t.Errorf("Accounts.DeleteAccount should return ok")
	}
}
