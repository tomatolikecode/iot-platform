package define

import "github.com/golang-jwt/jwt/v4"

var MysqlDSN = "root:123456@tcp(43.139.116.74:3306)" //os.Getenv("MYSQLDSN")

type M map[string]interface{}

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "iot-platform"
)
