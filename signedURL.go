package agozon

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"net/url"
	"strings"
)

/* Type SignedURL */
// SignedURL is important to make call AWSEcommerce Rest
type SignedURL string

// Create new Signed URL.
// This is such a constructor.
func (s *SignedURL) Create(theURL string, secretKey string) {
	// parse the url, fetch the information
	u, _ := url.Parse(theURL)

	// Sort the Query
	sortedQuery := s.SortedQuery([]byte(u.RawQuery))
	// Sign the url with the information
	sign := s.Sign(u.Host, u.Path, sortedQuery, secretKey)

	// Construct new query
	v := u.Query()
	// Add signature
	v.Add("Signature", sign)
	// Encode again
	u.RawQuery = v.Encode()

	// Convert string to signedURL.
	*s = SignedURL(u.String())
}

// Sort the query.
// This step is required
func (s *SignedURL) SortedQuery(rawquery []byte) string {
	// Split byte
	bb := bytes.Split(rawquery, []byte("&"))
	// Split query into array string
	sq := []string{}
	for _, b := range bb {
		sq = append(sq, string(b))
	}
	// join string with &
	return strings.Join(sq, "&")
}

// Sign the string with host, path, canonical and secretKey.
// This step is also required
func (s *SignedURL) Sign(host string, path string, canonical string, secretKey string) string {
	// create string to sign
	stringToSign := "GET" + "\n"
	stringToSign += host + "\n"
	stringToSign += path + "\n"
	stringToSign += canonical

	// create hash
	h := hmac.New(func() hash.Hash {
		return sha256.New()
	}, []byte(secretKey))
	// write string to sign to hash
	h.Write([]byte(stringToSign))

	// final signed string
	signed := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signed
}

// convert signedURL to string
func (s *SignedURL) String() string {
	return string(*s)
}
