package authfacebook

import (
	"net/http"
)

/*
const (
	nextPageKey     = "next_page" // 세션에 저장되는 next page의 키
	authSecurityKey = "auth_security_key"
)
*/

func init() {
	// 동작안함
	/*
		// gomniauth 정보 세팅
		gomniauth.SetSecurityKey(authSecurityKey)
		gomniauth.WithProviders(
			facebook.New("636296155193-a9abes4mc1p81752l116qkr9do6oev3f.apps.googleusercontent.com", "EVvuy0Agv4jWflml0pvC6-vI", "http://127.0.0.1:3000/auth/callback/google"),
		)
	*/
}

// AuthFacebook 함수는 페이스북 소셜 로그인 작업을 수행합니다.
func AuthFacebook(w http.ResponseWriter, r *http.Request, action string, provider string) {
	// 동작안함
	/*
		s := sessions.GetSession(r)

		switch action {
		case "login":
			// gomniauth.Provider의 login 페이지로 이동
			p, err := gomniauth.Provider(provider)
			if err != nil {
				log.Fatalln(err)
			}
			loginURL, err := p.GetBeginAuthURL(nil, nil)
			if err != nil {
				log.Fatalln(err)
			}
			http.Redirect(w, r, loginURL, http.StatusFound)
		case "callback":
			// gomniauth 콜백 처리
			p, err := gomniauth.Provider(provider)
			if err != nil {
				log.Fatalln(err)
			}
			creds, err := p.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
			if err != nil {
				log.Fatalln(err)
			}

			// 콜백 결과로부터 사용자 정보 확인
			user, err := p.GetUser(creds)
			if err != nil {
				log.Fatalln(err)
			}

			if err != nil {
				log.Fatalln(err)
			}

			u := &serversession.SessionUser{
				UID:       user.Data().Get("id").MustStr(),
				Name:      user.Name(),
				Email:     user.Email(),
				AvatarURL: user.AvatarURL(),
			}

			serversession.SetCurrentUser(r, u)
			http.Redirect(w, r, s.Get(nextPageKey).(string), http.StatusFound)
		default:
			http.Error(w, "Auth action '"+action+"' is not supported", http.StatusNotFound)
		}
	*/
}
