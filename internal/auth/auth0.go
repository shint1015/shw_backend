package auth

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const defaultJWKSCacheTTL = 24 * time.Hour

type Auth0Validator struct {
	issuer   string
	audience string
	jwksURL  string
	client   *http.Client

	mu          sync.RWMutex
	keys        map[string]*rsa.PublicKey
	cacheExpiry time.Time
}

type jwks struct {
	Keys []jwk `json:"keys"`
}

type jwk struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func NewAuth0ValidatorFromEnv() (*Auth0Validator, error) {
	issuer := strings.TrimSpace(os.Getenv("AUTH0_ISSUER_BASE_URL"))
	if issuer == "" {
		return nil, errors.New("AUTH0_ISSUER_BASE_URL is required")
	}
	audience := strings.TrimSpace(os.Getenv("AUTH0_AUDIENCE"))
	if audience == "" {
		return nil, errors.New("AUTH0_AUDIENCE is required")
	}
	jwksURL := strings.TrimRight(issuer, "/") + "/.well-known/jwks.json"
	return NewAuth0Validator(issuer, audience, jwksURL), nil
}

func NewAuth0Validator(issuer, audience, jwksURL string) *Auth0Validator {
	return &Auth0Validator{
		issuer:   strings.TrimRight(issuer, "/") + "/",
		audience: audience,
		jwksURL:  jwksURL,
		client:   &http.Client{Timeout: 10 * time.Second},
		keys:     map[string]*rsa.PublicKey{},
	}
}

func (v *Auth0Validator) Validate(ctx context.Context, tokenStr string) (*jwt.Token, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, v.keyFunc(ctx), jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Alg()}))
	if err != nil {
		return nil, err
	}
	if !claims.VerifyIssuer(v.issuer, true) {
		return nil, fmt.Errorf("invalid issuer")
	}
	if !claims.VerifyAudience(v.audience, true) {
		return nil, fmt.Errorf("invalid audience")
	}
	return token, nil
}

func (v *Auth0Validator) keyFunc(ctx context.Context) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		kid, ok := token.Header["kid"].(string)
		if !ok || kid == "" {
			return nil, fmt.Errorf("missing kid")
		}
		key, err := v.getKey(ctx, kid)
		if err != nil {
			return nil, err
		}
		return key, nil
	}
}

func (v *Auth0Validator) getKey(ctx context.Context, kid string) (*rsa.PublicKey, error) {
	v.mu.RLock()
	if v.cacheExpiry.After(time.Now()) {
		if key, ok := v.keys[kid]; ok {
			v.mu.RUnlock()
			return key, nil
		}
	}
	v.mu.RUnlock()

	if err := v.refreshKeys(ctx); err != nil {
		return nil, err
	}

	v.mu.RLock()
	defer v.mu.RUnlock()
	key, ok := v.keys[kid]
	if !ok {
		return nil, fmt.Errorf("unknown kid")
	}
	return key, nil
}

func (v *Auth0Validator) refreshKeys(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, v.jwksURL, nil)
	if err != nil {
		return err
	}
	resp, err := v.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("jwks fetch failed: %s", resp.Status)
	}

	var jwksPayload jwks
	if err := json.NewDecoder(resp.Body).Decode(&jwksPayload); err != nil {
		return err
	}

	keys := map[string]*rsa.PublicKey{}
	for _, k := range jwksPayload.Keys {
		if k.Kty != "RSA" || k.Kid == "" || k.N == "" || k.E == "" {
			continue
		}
		pubKey, err := parseRSAPublicKey(k.N, k.E)
		if err != nil {
			continue
		}
		keys[k.Kid] = pubKey
	}
	if len(keys) == 0 {
		return fmt.Errorf("no valid jwk keys found")
	}

	v.mu.Lock()
	defer v.mu.Unlock()
	v.keys = keys
	v.cacheExpiry = time.Now().Add(defaultJWKSCacheTTL)
	return nil
}

func parseRSAPublicKey(n, e string) (*rsa.PublicKey, error) {
	nb, err := base64.RawURLEncoding.DecodeString(n)
	if err != nil {
		return nil, err
	}
	eb, err := base64.RawURLEncoding.DecodeString(e)
	if err != nil {
		return nil, err
	}
	eInt := 0
	for _, b := range eb {
		eInt = eInt<<8 + int(b)
	}
	if eInt == 0 {
		return nil, fmt.Errorf("invalid exponent")
	}
	return &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: eInt,
	}, nil
}
