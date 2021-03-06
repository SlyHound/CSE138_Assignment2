package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRequest(r *gin.Engine, dict map[string]string) {

	r.DELETE("/key-value-store/:key", func(c *gin.Context) {
		key := c.Param("key")

		// if the key-value pair exists, then delete it //
		if _, exists := dict[key]; exists {
			c.JSON(http.StatusOK, gin.H{"doesExist": true, "message": "Deleted successfully"})
			delete(dict, key)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"doesExist": false, "error": "Key does not exist", "message": "Error in DELETE"})
		}
	})
}
