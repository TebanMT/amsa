package cognito

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/TebanMT/amsa/domain/entities"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type AuthRepository struct {
	client *cognitoidentityprovider.CognitoIdentityProvider
}

func NewCognitoAuthService() *AuthRepository {
	// Crear una sesión de AWS
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	// Crear un cliente de Cognito
	client := cognitoidentityprovider.New(sess)
	return &AuthRepository{
		client: client,
	}
}

func computeSecretHash(clientSecret string, username string, clientId string) string {
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientId))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func (repo *AuthRepository) LogIn(ctx context.Context, username, password string) (string, error) {

	secretHash := computeSecretHash("1cc122nfia7krmoj5pmkr8l9vm2muvirm2nj2snl1193n5fuc8l5", username, "428bs1od6aoqdhicqm94groolg")

	// Parámetros para la solicitud de inicio de sesión
	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: map[string]*string{
			"USERNAME":    aws.String(username),
			"PASSWORD":    aws.String(password),
			"SECRET_HASH": aws.String(secretHash),
		},
		ClientId: aws.String("428bs1od6aoqdhicqm94groolg"),
	}

	// Realizar la solicitud de inicio de sesión
	resp, err := repo.client.InitiateAuth(params)
	fmt.Printf("Cognito response: %s ..... %s \n", resp, resp.AuthenticationResult)
	if err != nil {
		return "", err
	}

	if resp.AuthenticationResult != nil && resp.AuthenticationResult.AccessToken != nil {
		accessToken := *resp.AuthenticationResult.AccessToken
		// Utilizar el token de acceso según sea necesario
		fmt.Printf("Cognito response: %s", accessToken)
	}

	return "resp", nil
}

func (repo *AuthRepository) RegisterUser(ctx context.Context, user entities.User) (*entities.User, error) {
	//secretHash := computeSecretHash("1cc122nfia7krmoj5pmkr8l9vm2muvirm2nj2snl1193n5fuc8l5", user.Username, "428bs1od6aoqdhicqm94groolg")
	signUpInput := &cognitoidentityprovider.AdminCreateUserInput{
		UserPoolId:    aws.String("428bs1od6aoqdhicqm94groolg"), // El ID de cliente de tu User Pool
		Username:      aws.String(user.Username),
		MessageAction: aws.String("SUPPRESS"),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String("email@example.com"),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String("+524151266863"),
			},
		},
	}
	signUpOutput, err := repo.client.AdminCreateUser(signUpInput)
	fmt.Printf("RESULT SINGOUTPUT === ", signUpOutput, err)
	if err != nil {
		// Maneja el error (por ejemplo, usuario ya existe, contraseña no cumple con los requisitos, etc.)
		return nil, err
	}
	return &user, nil
}
