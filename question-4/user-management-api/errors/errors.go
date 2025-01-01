package errors

// Error Messages
const (
	ErrNotFound     = "kayıt bulunamadı"
	ErrInvalidInput = "Geçersiz giriş"
	ErrInternal     = "Dahili bir hata oluştu"
	ErrCreateUser   = "kullanıcı oluşturma hatası"
	ErrLastInsertId = "son eklenen ID alınamadı"
	ErrGetUsers     = "kullanıcıları getirme hatası"
	ErrReadUsers    = "kullanıcı verisi okuma hatası"
	ErrUpdateUsers  = "kullanıcı güncelleme hatası"
	ErrRowsAffected = "etkilenen satır sayısı alınamadı"
	ErrDeleteUser   = "kullanıcı silme hatası"
)
