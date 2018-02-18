package main

import (
	"net/http"
	"log"
	"net/url"
)

type RedirectHandler struct {
	Opts Options
}

func NewRedirectHandler(opts Options) RedirectHandler {
	return RedirectHandler{
		Opts: opts,
	}
}

func (h RedirectHandler) buildRedirectURL(req http.Request) (string, error) {
	log.Printf("before: %s", req.URL.String())

	u, err := url.Parse(h.Opts.RedirectURL)
	if err != nil {
		return "", err
	}
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host

	log.Printf("after: %s", req.URL.String())

	return req.URL.String(), nil
}

func (h RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, err := h.buildRedirectURL(*r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	http.Redirect(w, r, u, http.StatusMovedPermanently)
}