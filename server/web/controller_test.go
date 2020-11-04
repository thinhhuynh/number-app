package web

import (
	"net/http"
	"net/http/httptest"
	"number-app/model"
	"testing"
)

type MockDb struct {
	tech []*model.Technology
	err  error
}

func (m *MockDb) GetTechnologies() ([]*model.Technology, error) {
	return m.tech, m.err
}

func TestApp_FindHighestPrimeNumber(t *testing.T) {
	app := App{d: &MockDb{
		tech: nil,
	}}

	r, _ := http.NewRequest("GET", "/api/find-highest-prime/?number=15", nil)
	w := httptest.NewRecorder()

	app.FindHighestPrimeNumber(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}

	want := `13` + "\n"
	got := w.Body.String()
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func TestApp_FindHighestPrimeNumber_WithNotFoundValue(t *testing.T) {
	app := App{d: &MockDb{
		tech: nil,
	}}

	r, _ := http.NewRequest("GET", "/api/find-highest-prime/?number=2", nil)
	w := httptest.NewRecorder()

	app.FindHighestPrimeNumber(w, r)

	want := `-1` + "\n"
	got := w.Body.String()

	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
}

func TestApp_FindHighestPrimeNumber_WithInvalidInputNumberError(t *testing.T) {
	app := App{d: &MockDb{
		tech: nil,
	}}

	r, _ := http.NewRequest("GET", "/api/find-highest-prime/?number=a", nil)
	w := httptest.NewRecorder()

	app.FindHighestPrimeNumber(w, r)

	want := `{"error":"INVALID_INPUT_NUMBER"}` + "\n"
	got := w.Body.String()
	
	if got != want {
		t.Errorf("handler returned unexpected body: got %v want %v", got, want)
	}
	if w.Code != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", w.Code, http.StatusOK)
	}
}
