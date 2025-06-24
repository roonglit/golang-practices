package db

import "context"

func (s SQLStore) CreateUserWithAuditLog(ctx context.Context, arg CreateUserParams) (User, error) {
	var user User
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		user, err = q.CreateUser(ctx, arg)
		if err != nil {
			return err
		}

		auditLog := CreateAuditLogParams{
			UserID: user.ID,
			Action: "create",
		}
		_, err = q.CreateAuditLog(ctx, auditLog)
		if err != nil {
			return err
		}
		return nil
	})

	return user, err
}
