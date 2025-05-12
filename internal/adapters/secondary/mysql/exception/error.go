package exception

import (
	"backend/pkg/response"
	"errors"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"net/http"
)

// Error messages
const (
	ErrMsgDataNotFound    = "data not found"
	ErrInvalidTransaction = "invalid transaction"
	ErrNotImplemented     = "not implemented"
)

func HandleGormError(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return response.NewAppErrorStatusMessage(http.StatusNotFound, errors.New(ErrMsgDataNotFound))
	case errors.Is(err, gorm.ErrInvalidTransaction):
		return response.NewAppErrorStatusMessage(http.StatusConflict, errors.New(ErrInvalidTransaction))
	case errors.Is(err, gorm.ErrNotImplemented):
		return response.NewAppErrorStatusMessage(http.StatusConflict, errors.New(ErrNotImplemented))
	case errors.Is(err, gorm.ErrMissingWhereClause):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("missing WHERE clause"))
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("unsupported relation"))
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("primary key required"))
	case errors.Is(err, gorm.ErrModelValueRequired):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("model value required"))
	case errors.Is(err, gorm.ErrInvalidData):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("invalid data"))
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		return response.NewAppErrorStatusMessage(http.StatusInternalServerError, errors.New("unsupported database driver"))
	case errors.Is(err, gorm.ErrRegistered):
		return response.NewAppErrorStatusMessage(http.StatusInternalServerError, errors.New("already registered"))
	case errors.Is(err, gorm.ErrInvalidField):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("invalid field"))
	case errors.Is(err, gorm.ErrEmptySlice):
		return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("empty slice found"))
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		return response.NewAppErrorStatusMessage(http.StatusInternalServerError, errors.New("dry run mode not supported"))
	case errors.Is(err, gorm.ErrInvalidDB):
		return response.NewAppErrorStatusMessage(http.StatusInternalServerError, errors.New("invalid database"))
	default:
		// Handle other database errors (like connection issues, constraint violations)
		if err != nil {
			// Check for specific database driver errors
			// For MySQL
			if mysqlErr, ok := err.(*mysql.MySQLError); ok {
				switch mysqlErr.Number {
				case 1062: // Error 1062: Duplicate entry
					return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("duplicate entry"))
				case 1452: // Error 1452: Cannot add or update a child row: a foreign key constraint fails
					return response.NewAppErrorStatusMessage(http.StatusBadRequest, errors.New("foreign key constraint violation"))
				}
			}
			// Generic database error
			return response.NewAppErrorStatusMessage(http.StatusInternalServerError, errors.New("Database error: "+err.Error()))
		}
		return nil
	}
}
