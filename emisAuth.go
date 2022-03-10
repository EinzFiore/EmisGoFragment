package emisAuth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	conf "gitlab.com/EinzFiore/emis-modules/configs"
	helpers "gitlab.com/EinzFiore/emis-modules/src/Helpers"
)

func EmisAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authUrl := fmt.Sprintf("%s/me", conf.GetConfig().AccountServiceUrl)
		response := helpers.UnauthorizedRes("Unauthorized", http.StatusUnauthorized, nil)

		// set cache key
		cacheKey := fmt.Sprintf("%s", c.GetHeader("Authorization"))

		cacheData, found := helpers.GetCache(cacheKey)
		if found == false {
			authRes, err := helpers.RequestWithAuth(http.MethodPost, authUrl, c)
			if err != nil {
				response.Message = err.Error()
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			// set cache
			var userData UserData
			err = json.Unmarshal([]byte(string(authRes)), &userData)

			if err != nil {
				response.Message = err.Error()
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			helpers.SetCache(cacheKey, userData, 10*time.Minute)
			c.Set("_userData", userData)
		}

		if cacheData != nil {
			var userData UserData
			err := json.Unmarshal([]byte(string(cacheData)), &userData)

			if err != nil {
				response.Message = err.Error()
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			c.Set("_userData", userData)
		}
	}
}

func EmisAuthIris(c iris.Context) {
	authUrl := fmt.Sprintf("%s/me", conf.GetConfig().AccountServiceUrl)
	response := helpers.UnauthorizedRes("Unauthorized", http.StatusUnauthorized, nil)

	// set cache key
	cacheKey := fmt.Sprintf("%s", c.GetHeader("Authorization"))

	cacheData, found := helpers.GetCache(cacheKey)
	if found == false {
		authRes, err := helpers.IrisRequestHTTP(http.MethodPost, authUrl, c)
		if err != nil {
			response.Message = err.Error()
			c.StatusCode(http.StatusUnauthorized)
			c.JSON(response)
			return
		}

		// set cache
		var userData UserData
		err = json.Unmarshal([]byte(string(authRes)), &userData)

		if err != nil {
			response.Message = err.Error()
			c.StatusCode(http.StatusUnauthorized)
			c.JSON(response)
			return
		}

		helpers.SetCache(cacheKey, userData, 10*time.Minute)
		c.Values().Set("_userData", userData)
	}

	if cacheData != nil {
		var userData UserData
		err := json.Unmarshal([]byte(string(cacheData)), &userData)

		if err != nil {
			response.Message = err.Error()
			c.StatusCode(http.StatusUnauthorized)
			c.JSON(response)
			return
		}
		c.Values().Set("_userData", userData)
	}
}
