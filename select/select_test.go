package select_

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compare speeds of servers", func(t *testing.T) {
		slowServer := makeDeplayedServer(20 * time.Millisecond)
		fastServer := makeDeplayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Errorf("didn't expected err but got one %s", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return error if server is non-responding", func(t *testing.T) {
		slowServer := makeDeplayedServer(20 * time.Millisecond)
		fastServer := makeDeplayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)
		if err != nil {
			t.Errorf("didn't expected err but got one %s", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func makeDeplayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
