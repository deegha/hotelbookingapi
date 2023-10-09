package utils
import (
  "github.com/golang-jwt/jwt"
    "time"
)
const SecretKey = "secret"
func ValidateCookie(cookie string) (*jwt.StandardClaims, error){
   token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error) {
     return []byte(SecretKey), nil})
	if err != nil {
		return nil, err
	}

  claims := token.Claims.(*jwt.StandardClaims)


	return claims, nil
}


func CalculateNumberOfNights(checkInTime, checkOutTime time.Time) (int, error) {

    // Calculate the duration between check-in and check-out times
    duration := checkOutTime.Sub(checkInTime)

    // Calculate the number of nights
    nights := int(duration.Hours() / 24)

    // If check-out is on the same day, subtract 1 night
    if checkInTime.Year() == checkOutTime.Year() &&
        checkInTime.Month() == checkOutTime.Month() &&
        checkInTime.Day() == checkOutTime.Day() {
        nights--
    }

    return nights, nil
}

