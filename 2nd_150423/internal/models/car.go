package models

// CRUD - CREATE/READ/UPDATE/DELETE
type (
	CarView struct {
		ID    uint8  `json:"id"`
		Model string `db:"model"`
		Typ   string `db:"type"` // внедорожник/легковая/грузовая ...
	}

	CarCreate struct {
		ID    uint8  `json:"id"`
		Brand string `db:"brand"`
		Model string `db:"model"`
		Typ   string `db:"type"` // внедорожник/легковая/грузовая ...
	}

	CarDelete struct {
		ID uint8
	}

	// CarObject - полное описание объекта так как он описан в базе данных
	CarObject struct {
		ID uint8 `db:"id"`

		Brand   string `db:"brand"`
		Model   string `db:"model"`
		Typ     string `db:"type"` // внедорожник/легковая/грузовая ...
		Field01 string `db:"field_01"`
		Field02 string `db:"field_02"`
		Field03 string `db:"field_03"`
		Field04 string `db:"field_04"`
		Field05 string `db:"field_05"`

		IsCertificated bool `db:"cert"` // имеет ли авто сертификацию
	}
)
