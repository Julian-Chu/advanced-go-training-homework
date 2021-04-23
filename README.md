# advanced-go-training-homework

## week2
Q: When getting sql.ErrNoRows, should we wrap it and return?
my answer is no, please see comment below.
```go
err := s.db.QueryRow(`select user_id from user where email = $1`, email).Scan(&user.ID)
if err !=nil{
    // sql.ErrorNoRows means "user not found" in business logic level.
    // don't return sql.ErrorNoRows or it wrapped error because it's not a real database error
    // recommend to use custom error for business logic level
    if err == sql.ErrNoRows{
        return User{}, ErrUserNotFound
    }
    return User{}, err
}
return user, nil
```