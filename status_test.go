package obs

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"testing"
)

func toJsonStr(v any) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}

func TestStatusJsonify(t *testing.T) {
	s := New(codes.OK, "Hello")
	assert.Equal(t, ``, toJsonStr(s))
}
