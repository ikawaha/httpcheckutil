package calcapi

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	gen "calc/gen/calc"
	"calc/gen/http/calc/server"
	"github.com/ikawaha/httpcheck"
	"github.com/ikawaha/httpcheckutil/goa/mounter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMounter(t *testing.T) {
	m := mounter.New()
	m.Mount(mounter.EndpointModular{
		Builder:  server.NewMultiplyHandler,
		Mounter:  server.MountMultiplyHandler,
		Endpoint: gen.NewMultiplyEndpoint(NewCalc(log.New(os.Stdout, "", log.LstdFlags))),
	})
	httpcheck.New(m).Test(t, "GET", "/multiply/3/5").
		Check().
		MustHasStatus(http.StatusOK).
		Cb(func(r *http.Response) {
			b, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			i, err := strconv.Atoi(strings.TrimSpace(string(b)))
			assert.NoError(t, err, string(b))
			assert.Equal(t, 3*5, i)
		})
}
