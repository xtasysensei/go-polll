package routes

import (
	"flag"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/docgen"
	"github.com/xtasysensei/go-poll/internal/mymiddleware"
	"github.com/xtasysensei/go-poll/pkg/handlers"
	"github.com/xtasysensei/go-poll/pkg/handlers/poll"
	"github.com/xtasysensei/go-poll/pkg/handlers/user"
	"github.com/xtasysensei/go-poll/pkg/handlers/vote"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func RegisterRoutes(apiRouter *chi.Mux) {
	flag.Parse()
	apiRouter.Get("/", handlers.Index)
	apiRouter.Get("/ping", handlers.Health)

	apiRouter.Route("/v1", func(route chi.Router) {
		//user routes
		route.Route("/auth", func(r chi.Router) {
			r.Post("/login", user.HandleLogin)
			r.Post("/register", user.HandleRegister)
		})

		//poll routes
		route.With(mymiddleware.WithUserID).Route("/polls", func(r chi.Router) {
			r.Post("/", poll.HandleCreatePoll)
			r.Get("/", poll.RetrieveAllPolls)
			r.Get("/{pollId}", poll.RetrievePollByID)
			//r.With(paginate).Get("/", ListArticles)
			// r.Get("/search", SearchArticles)

			//vote
			r.Post("/{pollId}/vote", vote.HandleCastVote)
		})

	})

	if *routes {
		fmt.Println(docgen.MarkdownRoutesDoc(apiRouter, docgen.MarkdownOpts{
			ProjectPath: "go-mongo",
			Intro:       "Welcome to the go-mongo generated docs.",
		}))
		return
	}

}
