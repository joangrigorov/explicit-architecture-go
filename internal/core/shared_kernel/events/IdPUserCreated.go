package events

import "app/internal/core/shared_kernel/domain"

type IdPUserCreated struct {
	userID  domain.UserID
	idpUser domain.IdPUserID
}

func (i IdPUserCreated) ID() EventID {
	return makeEventID(i)
}

func (i IdPUserCreated) UserID() string {
	return i.userID.String()
}

func NewIdPUserCreated(userID domain.UserID, idpUserID domain.IdPUserID) IdPUserCreated {
	return IdPUserCreated{
		userID:  userID,
		idpUser: idpUserID,
	}
}
