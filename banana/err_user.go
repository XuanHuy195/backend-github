package banana

import "errors"

var (
	UserConflict = errors.New("Nguời dùng đã tồn tại")
	SignUpFail   = errors.New("Đăng ký thất bại")
	UserNotFound = errors.New("Người dùng không tồn tại")
	UserNotUpdated = errors.New("Cập nhật thất bại")
)
