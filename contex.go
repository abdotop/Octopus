package octopus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net"
	"net/http"
	"strings"
)

type Ctx struct {
	// sync.RWMutex
	handlers []HandlerFunc
	index    int
	Values   *value
	Context  context.Context
	// sse      *sse
}

func NewCtx() *Ctx {
	return &Ctx{
		handlers: []HandlerFunc{},
		index:    0,
		Values:   new(value),
		Context:  context.Background(),
	}
}

type Map = map[string]interface{}

func (ctx *Ctx) AppStore() (*value, error) {
	app_value, ok := ctx.Values.Get("app")
	if !ok {
		return nil, fmt.Errorf("failed to get App from context")
	}
	app, ok := app_value.(*App)
	if !ok {
		return nil, fmt.Errorf("failed to get App from context")
	}
	return app.Store, nil
}

func (c *Ctx) BasicAuth() (string, string, bool) {
	r, ok := c.Values.Get("request")
	if !ok {
		return "", "", false
	}
	request, ok := r.(*http.Request)
	if !ok {
		return "", "", false
	}

	return request.BasicAuth()
}

func (ctx *Ctx) BodyParser(out interface{}) error {
	// c.RLock()
	// defer c.RUnlock()
	r, ok := ctx.Values.Get("request")
	if ok {
		r := r.(*http.Request)
		return json.NewDecoder(r.Body).Decode(&out)
	}
	return errors.New("request not found in context values")
}

// Get returns the value of the key in the context header
func (ctx *Ctx) Get(key string) string {
	// c.RLock()
	// defer c.RUnlock()
	r, ok := ctx.Values.Get("request")
	if ok {
		r := r.(*http.Request)
		return r.Header.Get(key)
	}
	return ""
}

// func (c *Ctx) GetSse() (*sse, error) {
// 	app_value, ok := c.Values.Get("app")
// 	if !ok {
// 		return nil, fmt.Errorf("failed to get App from context")
// 	}
// 	app, ok := app_value.(*App)
// 	if !ok {
// 		return nil, fmt.Errorf("failed to get App from context")
// 	}
// 	return app.sse_service, nil
// }

func (ctx *Ctx) JSON(data interface{}) error {
	// c.Lock()
	// defer c.Unlock()
	r := ctx.Response()
	r.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(r).Encode(data)
}

func (ctx *Ctx) Next() {
	if ctx.index < len(ctx.handlers) {
		handler := ctx.handlers[ctx.index]
		ctx.index++
		handler(ctx)
	}
}

func (ctx *Ctx) Query(key string) string {
	// c.RLock()
	// defer c.RUnlock()
	r, ok := ctx.Values.Get("request")
	if ok {
		r := r.(*http.Request)
		return r.URL.Query().Get(key)
	}
	return ""
}

func (ctx *Ctx) Render(path string, data interface{}) error {
	// c.Lock()
	// defer c.Unlock()
	r := ctx.Response()
	tp, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	return tp.Execute(r, data)
}

func (ctx *Ctx) SendString(code StatusCode, s string) error {
	// c.Lock()
	// defer c.Unlock()
	r := ctx.Response()
	ctx.Status(code)
	_, err := r.Write([]byte(s))
	return err
}

func (ctx *Ctx) Status(code StatusCode) *Ctx {
	// c.RLock()
	// defer c.RUnlock()
	r := ctx.Response()
	a, appExist := ctx.Values.Get("app")
	r.WriteHeader(int(code))
	if appExist {
		a := a.(*App)
		a.handleError(code, ctx)
	} else {
		a := New()
		a.handleError(code, ctx)
	}

	return ctx
}

func (ctx *Ctx) RemoteIP() (string, error) {
	r, ok := ctx.Values.Get("request")
	if !ok {
		return "", errors.New("request not found in context")
	}

	req := r.(*http.Request)
	ips := extractValidIPsFromHeader(req, "X-Forwarded-For")
	if len(ips) > 0 {
		return ips[0], nil // retourne la première IP valide
	}

	// Fallback sur l'adresse IP directe
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	return ip, nil
}

func (ctx *Ctx) Response() http.ResponseWriter {
	r, ok := ctx.Values.Get("response")
	if !ok {
		// Comme nous ne pouvons pas retourner une erreur ici, nous allons logger l'erreur
		// et retourner un ResponseWriter vide
		return &emptyResponseWriter{}
	}

	resp, ok := r.(http.ResponseWriter)
	if !ok {
		return &emptyResponseWriter{}
	}

	return resp
}

// emptyResponseWriter est un ResponseWriter vide qui ne fait rien
type emptyResponseWriter struct{}

func (e *emptyResponseWriter) Header() http.Header       { return http.Header{} }
func (e *emptyResponseWriter) Write([]byte) (int, error) { return 0, nil }
func (e *emptyResponseWriter) WriteHeader(int)           {}

// extractValidIPsFromHeader extrait et valide les adresses IP à partir d'un en-tête HTTP spécifié.
func extractValidIPsFromHeader(r *http.Request, headerName string) []string {
	headerValue := r.Header.Get(headerName)
	if headerValue == "" {
		return nil
	}

	ips := strings.Split(headerValue, ",")
	validIPs := make([]string, 0, len(ips))

	for _, ip := range ips {
		trimmedIP := strings.TrimSpace(ip)
		if isValidIP(trimmedIP) {
			validIPs = append(validIPs, trimmedIP)
		}
	}

	return validIPs
}

// isValidIP vérifie si une chaîne est une adresse IP valide.
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// func (c *Ctx) GetSseConnection(id string) (*sseConn, error) {
// 	app_value, ok := c.Values.Get("app")

// 	if !ok {
// 		return nil, fmt.Errorf("failed to retrieve app from context")
// 	}
// 	app, ok := app_value.(*App)
// 	if !ok {
// 		return nil, fmt.Errorf("failed to retrieve app from context")
// 	}
// 	conn, ok := app.getConnection(id)
// 	if !ok {
// 		return nil, fmt.Errorf("no connection found with ID %s", id)
// 	}

// 	return conn, nil
// }

func (ctx *Ctx) WriteString(s string) error {
	r := ctx.Response()
	_, err := r.Write([]byte(s))
	return err
}
