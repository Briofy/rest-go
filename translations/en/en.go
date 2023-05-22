package en

var Messages = map[string]interface{}{

	"success": "success",
	"failed":  "failed",

	"errors": map[string]map[string]interface{}{

		"general": {
			//error messages
			"error": "Error has been detected!",
			"not_found":  "No resource was found!",
			"wrong_parameter": "Wrong parameter has been detected!",
			"invalid_query": "Invalid query has been detected!",
			"invalid_parameters": "Invalid parameters has been detected!",
			"unauthorized": "Unauthorized action has been detected!",
			"forbidden":  "Forbidden action has been detected!",
			"not_acceptable": "Unacceptable action detected!",
			"access_deny": "Access Deny!",
			"not_exist":" The desired item is not available in the system.",
			// return messages
			"entity_removed": "Entity has been removed successfully!",
			"successful_message": "Response Successfully Returned.",
			"successful_create_message": "The data has been successfully created.",
			"successful_update_message": "The data has been successfully updated.",
			// authentication
			"login_successfully": "User logged in successfully.",
			"qr_generated": "Qr Code login image created successfully.",
			// license
			"active_license": "license is activated",
			// google2fa
			"email_not_set": "Email not set, You have to first add your email to account",
			"verify_failed": "Secret or Code is invalid",
			"unauthenticated_message": "The user is not authenticated",
		},
		"success": {
			"insert": "The requested parameter is added successfully!",
			"delete": "The requested parameter is deleted successfully!",
			"update": "The requested parameter is updated successfully!",
		},
		"failed": {
			"insert": "The requested parameter is not added!",
			"delete": "The requested parameter is not deleted!",
			"update": "The requested parameter is not updated!",
		},
		"1001": {
			"message": "Oops... Requested field is not found!",
			"type":    "error",
		},
		"1002": {
			"message": "Oops... Requested User does not exists!",
			"type":    "error",
		},
		"1003": {
			"message": "Oops... Client type is not entered!",
			"type":    "error",
		},
		"1004": {
			"message": "Failed because of duplicate",
			"type":    "error",
		},
		"1005": {
			"message": "Failed because of dablicated user role",
			"type":    "error",
		},
		"3001": {
			"message": "You are not logged on",
			"type":    "error",
			"cat":     "auth",
			"short":   "not-logged-on",
		},
		"3002": {
			"message": "Application token did not generated successfully",
			"type":    "error",
			"cat":     "auth",
		},
		"3003": {
			"message": "User token did not generated successfully",
			"type":    "error",
			"cat":     "auth",
		},
		"3005": {
			"message": "Request token did not contains user information",
			"type":    "error",
			"cat":     "auth",
		},
		"3006": {
			"message": "Did not set request token",
			"type":    "error",
			"cat":     "auth",
		},
		"3007": {
			"message": "can not decode the token",
			"type":    "error",
			"cat":     "auth",
		},
		"3008": {
			"message": "can not generate token for authentication",
			"type":    "error",
			"cat":     "auth",
		},
		"3009": {
			"message": "can not create token",
			"type":    "error",
			"cat":     "auth",
		},
		"3010": {
			"message": "Token expired!",
			"type":    "error",
			"cat":     "auth",
		},
		"3011": {
			"message": "Token is invalid!",
			"type":    "error",
			"cat":     "auth",
		},
		"3012": {
			"message": "Token Blacklisted",
			"type":    "error",
			"cat":     "auth",
		},
		"3013": {
			"message": "Payload invalid!",
			"type":    "error",
			"cat":     "auth",
		},
		"3014": {
			"message": "Claim Invalid",
			"type":    "error",
			"cat":     "auth",
		},
		"3015": {
			"message": "An error occurred on token validation",
			"type":    "error",
			"cat":     "auth",
		},
		"5401": {
			"message": "Authentication unauthorized...",
			"type":    "error",
		},
		"5404": {
			"message": "Oops... The requested page not found!",
			"type":    "error",
		},
		"5405": {
			"message": "Oops... The method you requested is not allowed!",
			"type":    "error",
		},
		"5406": {
			"message": "Oops... The parameters you entered are wrong!",
			"type":    "error",
		},
		"5420": {
			"message": "Validation Error",
			"type":    "error",
		},
		"5422": {
			"message": "Token is not valid",
			"type":    "error",
		},
		"5445": {
			"message": "Oops... Database connection refused",
			"type":    "error",
		},
		"5447": {
			"message": "Oops... Delete action was not successfully executed",
			"type":    "error",
		},
		"5448": {
			"message": "Oops... Insert action was not successfully executed",
			"type":    "error",
		},
		"5449": {
			"message": "Oops... Update action was not successfully executed",
			"type":    "error",
		},
	},
}
