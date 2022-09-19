package siwago

import "testing"

func TestAppleKeys(t *testing.T) {

	//multiple keys from apple
	kids := []string{"W6WcOKB", "fh6Bs8C"}

	for _, kid := range kids {
		t.Logf("Testing kid=%s", kid)
		//check for the apple key object
		applekey, err := getApplePublicKey(kid)
		if err != nil {
			t.Error(err)
			continue
		}
		if applekey.Kid != kid {
			t.Errorf("Invalid Key " + kid)
		}
		//check RSA object
		key, err := getApplePublicKeyObject(kid, "RS256")
		if err != nil {
			t.Error(err)
			continue
		}
		if key.N == nil {
			t.Errorf("Invalid Key " + kid)
		}
	}
}
