package db

// htpasswd -nb -B tal@whatexit.org 'THE PASSWORD'

func (d *dataAccess) GetPasswordHash(user string) string {
	var hash string
	err := d.db.Get(&hash, "SELECT pwhash FROM users WHERE email = ?", user)
	if err != nil {
		return ""
	}
	return hash
}
