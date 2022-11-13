package core

import (
	login "apirest/internal/jwt"
	"apirest/internal/logger"
	"apirest/internal/models"
	"apirest/internal/password"
	"apirest/internal/ws"
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type handlerCore struct {
	TxID string
}

func (h *handlerCore) GetUserByID(c *fiber.Ctx) error {
	request := RequestUserById{}
	res := Response{Error: true}
	err := c.BodyParser(&request)

	if err != nil || request.IdUser == "" {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición al modelo Request: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "La peticion realizada no es valida"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	template := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:exam="http://lasestrellitas.com/example">
   <soapenv:Header/>
   <soapenv:Body>
      <exam:getUserByIdRequest>
         <exam:id>%s</exam:id>
      </exam:getUserByIdRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	payload := fmt.Sprintf(template, request.IdUser)

	responseWS, code, err := ws.ConsumeSOAP(payload)

	if err != nil || code != 200 {
		logger.Error.Println("No se pudo consultar la persona: %v", err)
		res.Code, res.Type, res.Msg = 22, "error", "No se pudo consultar la persona"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	envelope := &ResponseByID{}

	err = xml.Unmarshal(responseWS, &envelope)

	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: "+err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = User{
		IdUser:   envelope.Body.GetUserByIdResponse.User.ID,
		Name:     envelope.Body.GetUserByIdResponse.User.Names,
		LastName: envelope.Body.GetUserByIdResponse.User.Lastname,
		Email:    envelope.Body.GetUserByIdResponse.User.Email,
		Age:      envelope.Body.GetUserByIdResponse.User.Age,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerCore) GetUserByEmail(c *fiber.Ctx) error {
	request := RequestUserByEmail{}
	res := Response{Error: true}
	err := c.BodyParser(&request)

	if err != nil || request.Email == "" {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición al modelo Request: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "La peticion realizada no es valida"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	template := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:exam="http://lasestrellitas.com/example">
   <soapenv:Header/>
   <soapenv:Body>
      <exam:getUserByEmailRequest>
         <exam:email>%s</exam:email>
      </exam:getUserByEmailRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	payload := fmt.Sprintf(template, request.Email)

	responseWS, code, err := ws.ConsumeSOAP(payload)

	if err != nil || code != 200 {
		logger.Error.Println("No se pudo consultar la persona: %v", err)
		res.Code, res.Type, res.Msg = 22, "error", "No se pudo consultar la persona"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	envelope := &ResponseByEmail{}

	err = xml.Unmarshal(responseWS, &envelope)

	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: "+err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = User{
		IdUser:   envelope.Body.GetUserByEmailResponse.User.ID,
		Name:     envelope.Body.GetUserByEmailResponse.User.Names,
		LastName: envelope.Body.GetUserByEmailResponse.User.Lastname,
		Email:    envelope.Body.GetUserByEmailResponse.User.Email,
		Age:      envelope.Body.GetUserByEmailResponse.User.Age,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerCore) CreateUser(c *fiber.Ctx) error {
	request := RequestCreateUser{}
	res := Response{Error: true}
	err := c.BodyParser(&request)

	if err != nil || request.Name == "" || request.LastName == "" || request.Email == "" || request.Age == "" || request.Gender == "" || request.Password == "" {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición al modelo Request: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "La peticion realizada no es valida"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	template := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:exam="http://lasestrellitas.com/example">
   <soapenv:Header/>
   <soapenv:Body>
      <exam:createUserRequest>
         <exam:id>%s</exam:id>
         <exam:names>%s</exam:names>
         <exam:lastname>%s</exam:lastname>
         <exam:email>%s</exam:email>
         <exam:age>%s</exam:age>
         <exam:gender>%s</exam:gender>
         <exam:password>%s</exam:password>
      </exam:createUserRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	passwordEncrypt := password.Encrypt(request.Password)

	payload := fmt.Sprintf(template, uuid.New().String(), request.Name, request.LastName, request.Email, request.Age, request.Gender, passwordEncrypt)

	responseWS, code, err := ws.ConsumeSOAP(payload)

	if err != nil || code != 200 {
		logger.Error.Println("No se pudo consultar la persona: %v", err)
		res.Code, res.Type, res.Msg = 22, "error", "No se pudo consultar la persona"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	envelope := &ResponseCreateUser{}

	err = xml.Unmarshal(responseWS, &envelope)

	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: "+err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = User{
		IdUser:   envelope.Body.GetUserByIdResponse.User.ID,
		Name:     envelope.Body.GetUserByIdResponse.User.Names,
		LastName: envelope.Body.GetUserByIdResponse.User.Lastname,
		Email:    envelope.Body.GetUserByIdResponse.User.Email,
		Age:      envelope.Body.GetUserByIdResponse.User.Age,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerCore) DeleteUser(c *fiber.Ctx) error {
	request := RequestDeleteUser{}
	res := Response{Error: true}
	err := c.BodyParser(&request)

	if err != nil || request.Email == "" || request.Password == "" {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición al modelo Request: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "La peticion realizada no es valida"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	template := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:exam="http://lasestrellitas.com/example">
   <soapenv:Header/>
   <soapenv:Body>
      <exam:deleteUserRequest>
         <exam:email>%s</exam:email>
         <exam:password>%s</exam:password>
      </exam:deleteUserRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	payload := fmt.Sprintf(template, request.Email, request.Password)

	responseWS, code, err := ws.ConsumeSOAP(payload)

	if err != nil || code != 200 {
		logger.Error.Println("No se pudo consultar la persona: %v", err)
		res.Code, res.Type, res.Msg = 22, "error", "No se pudo consultar la persona"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	envelope := &ResponseDeleteUser{}

	err = xml.Unmarshal(responseWS, &envelope)

	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: "+err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", envelope.Body.DeleteUserResponse.Message
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerCore) Login(c *fiber.Ctx) error {
	request := RequestLogin{}
	res := ResponseLogin{Error: true}
	err := c.BodyParser(&request)
	if err != nil {
		logger.Error.Printf("No se pudo parsear el cuerpo de la petición al modelo Request: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "La peticion realizada no es valida"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	template := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:exam="http://lasestrellitas.com/example">
   <soapenv:Header/>
   <soapenv:Body>
      <exam:getUserByEmailRequest>
         <exam:email>%s</exam:email>
      </exam:getUserByEmailRequest>
   </soapenv:Body>
</soapenv:Envelope>`

	payload := fmt.Sprintf(template, request.Email)

	responseWS, code, err := ws.ConsumeSOAP(payload)

	if err != nil || code != 200 {
		logger.Error.Println("No se pudo consultar la persona: %v", err)
		res.Code, res.Type, res.Msg = 22, "error", "No se pudo consultar la persona"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	envelope := &ResponseByEmail{}

	err = xml.Unmarshal(responseWS, &envelope)
	if err != nil {
		logger.Error.Println("No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "No se pudo parsear el cuerpo de la petición al modelo RequestCorePerson: "+err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}
	if !password.Compare(envelope.Body.GetUserByEmailResponse.User.ID, envelope.Body.GetUserByEmailResponse.User.Password, request.Password) {
		logger.Error.Println("Usuario o contraseña incorrecto")
		res.Code, res.Type, res.Msg = 22, "error", "Usuario o contraseña incorrecto"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	token, cod, err := login.GenerateJWT(models.User{
		IdUser:   envelope.Body.GetUserByEmailResponse.User.ID,
		Name:     envelope.Body.GetUserByEmailResponse.User.Names,
		LastName: envelope.Body.GetUserByEmailResponse.User.Lastname,
		Email:    envelope.Body.GetUserByEmailResponse.User.Email,
		Age:      envelope.Body.GetUserByEmailResponse.User.Age,
	})
	if err != nil {
		logger.Error.Printf("no se pudo el token de autorización: %v", err)
		res.Code, res.Type, res.Msg = cod, "error", "no se pudo el token de autorización"
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = Token{
		AccessToken:  token,
		RefreshToken: token,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}
