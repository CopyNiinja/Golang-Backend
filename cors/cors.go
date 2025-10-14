package cors

import (
	"strings"
	"time"

	"github.com/copyniinja/movie-streaming-app-go-react-langchaingo/server/flimflixServer/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)


func CorsSetup(cfg *config.Config) gin.HandlerFunc {
	origins := strings.Split(cfg.ALLOW_ORIGINS, ",")
	return cors.New(cors.Config{
				AllowOrigins:       origins,
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
				AllowCredentials: true,
		    MaxAge:           12 * time.Hour,

			})
	}
