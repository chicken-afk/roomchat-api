package auths

type AuthService interface {
	Login(request LoginRequest) (interface{}, error)
	Register(request RegisterRequest) (interface{}, error)
}

type aService struct {
}

func NewAuthService() AuthService {
	return &aService{}
}

func (a *aService) Login(request LoginRequest) (interface{}, error) {
	/**Call Repository**/

	var response LoginResponse
	response.Token = "ewafaofjds920r3jsda=29"
	response.Uuid = "24802-4534-klasd-3j329"
	response.Email = request.Email

	return response, nil
}

func (a *aService) Register(request RegisterRequest) (interface{}, error) {
	/**Call Repository**/

	var response RegisterResponse
	response.Uuid = "24802-4534-klasd-3j329"
	response.Email = request.Email
	response.Status = "Active"
	response.Token = "ewafaofjds920r3jsda=29"

	return response, nil
}
