package regression

import (

	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"github.com/jb843051627/mangrove-flux/internal/handler"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"github.com/jb843051627/mangrove-flux/internal/store"

)


func TestBug02_MissingBatchIsNotAnInternalServerError(t *testing.T) {
	db, err := store.Open(filepath.Join(t.TempDir(), "case.db")); if err != nil { t.Fatal(err) }
	app := service.NewLab(db); defer app.Close(); defer db.Close()
	req := httptest.NewRequest(http.MethodGet, "/batches/missing-batch", nil); rec := httptest.NewRecorder()
	handler.New(app).ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound { t.Fatalf("missing batch status = %d, body=%s", rec.Code, rec.Body.String()) }
}

