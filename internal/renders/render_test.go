package renders

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexedwards/scs/v2"
	"github.com/sleepiinuts/webapp-plain/configs"
	"github.com/sleepiinuts/webapp-plain/pkg/models"
	"github.com/sleepiinuts/webapp-plain/test"
)

var (
	r  *Renderer
	ap *configs.AppProperties
	sm *scs.SessionManager
)

func TestRenderFromMap(t *testing.T) {
	w := httptest.NewRecorder()
	rq := makeReqWithSession()

	if err := r.RenderTemplateFromMap(w, rq, "index.tmpl", &models.Template{}); err != nil {
		t.Fail()
		t.Logf("unexpected error, got %v\n", err)
	}
}

func init() {
	ap, sm = test.GetDependencies()
	r = New(ap, sm, "../../web/templates/")
}

func makeReqWithSession() *http.Request {

	// doesnt require actual METHOD & URL
	rq := httptest.NewRequest("GET", "/", nil)

	// following https://gist.github.com/alexedwards/cc6190195acfa466bf27f05aa5023f50
	// when to modify req with context

	ctx, err := sm.Load(rq.Context(), rq.Header.Get("X-Session"))
	if err != nil {
		panic("cant make request with session context")
	}

	return rq.WithContext(ctx)
}
