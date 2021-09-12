package response

type RequestError struct {
	Message string `json:"error"`
}

func (r *RequestError) Error() string {
	return r.Message
}

// func SendResponse(c *gin.Context, response RequestError) {
// 	c.JSON(response.StatusCode, map[string]interface{}{
// 		"message":     response.Message,
// 		"discription": response.Err.Error(),
// 	})
// }
