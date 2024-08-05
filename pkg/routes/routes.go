package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/xtasysensei/go-poll/internal/mymiddleware"
	"github.com/xtasysensei/go-poll/pkg/handlers"
	"github.com/xtasysensei/go-poll/pkg/handlers/poll"
	"github.com/xtasysensei/go-poll/pkg/handlers/user"
	"github.com/xtasysensei/go-poll/pkg/handlers/vote"
)

func RegisterRoutes(apiRouter *chi.Mux) {

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
			//r.With(paginate).Get("/", ListArticles)
			r.Post("/", poll.HandleCreatePoll)
			r.Get("/{pollId}", poll.RetrievePoll)
			// r.Get("/search", SearchArticles)

			//vote
			r.Post("/{pollId}/vote", vote.HandleCastVote)
		})

		// vote routes
		route.With(mymiddleware.WithUserID)

	})

}
