module github.com/ki365/ki365

go 1.23.1

require (
	github.com/go-chi/chi/v5 v5.0.14
	github.com/ki365/ki365/abatement v0.0.0
	github.com/ki365/ki365/kicad v0.0.0
	github.com/ki365/ki365/optimization v0.0.0
)

replace github.com/ki365/ki365/kicad => ./server/internal/kicad

replace github.com/ki365/ki365/optimization => ./server/internal/optimization

replace github.com/ki365/ki365/abatement => ./server/abatement

replace github.com/ki365/ki365/pkg => ./server/pkg

require (
	github.com/bramvdbogaerde/go-scp v1.5.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
)
