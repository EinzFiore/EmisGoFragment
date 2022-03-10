package emisAuth

import "time"

type ApiMeta struct {
	Version  int
	Message  string
	Metadata interface{}
	Errors   interface{}
	Results  UserData
}

type UserData struct {
	ID              int         `json:"id"`
	Name            string      `json:"name"`
	Nik             interface{} `json:"nik"`
	Email           string      `json:"email"`
	EmailVerifiedAt interface{} `json:"email_verified_at"`
	AccountStatus   int         `json:"account_status"`
	IssuanceStatus  int         `json:"issuance_status"`
	IssuanceBy      int         `json:"issuance_by"`
	IsRejected      int         `json:"is_rejected"`
	RejectionReason interface{} `json:"rejection_reason"`
	Passcode        interface{} `json:"passcode"`
	PasscodeExpired interface{} `json:"passcode_expired"`
	IsServices      int         `json:"is_services"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Institution     struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		CategoryID        int    `json:"category_id"`
		ProvinceID        string `json:"province_id"`
		CityID            string `json:"city_id"`
		IsActive          int    `json:"is_active"`
		LaravelThroughKey int    `json:"laravel_through_key"`
	} `json:"institution"`
	Role struct {
		ID            int         `json:"id"`
		UserID        int         `json:"user_id"`
		RoleID        int         `json:"role_id"`
		ProvinceID    interface{} `json:"province_id"`
		CityID        interface{} `json:"city_id"`
		DistrictID    interface{} `json:"district_id"`
		SubDistrictID interface{} `json:"sub_district_id"`
		InstitutionID int         `json:"institution_id"`
		MasterRole    struct {
			ID                int    `json:"id"`
			RoleName          string `json:"role_name"`
			Level             int    `json:"level"`
			IdentifiableField string `json:"identifiable_field"`
			RoleLevel         struct {
				ID          int    `json:"id"`
				Label       string `json:"label"`
				Description string `json:"description"`
			} `json:"role_level"`
			RoleScope struct {
				ID          int    `json:"id"`
				Label       string `json:"label"`
				Description string `json:"description"`
			} `json:"role_scope"`
		} `json:"master_role"`
	} `json:"role"`
}
