package main

type api_key struct {
	id           uint
	size_message uint
	message      []byte
}

type User struct {
	ID string
	Username string
	Password string
}

type who interface {
	login(user *User)  // определяет кто залогинился возвращает ID

} 
// Сначала приходит ключ на то что надо залогиниться
// Если залогинлися -> переходим к функциям. Если нет создаем юзера
// Есть 4 функций у юзера -> Создание, просмотр, изменение, удаление. 
//  Есть 6 функций у админа -> Создать юзера, удалить юзера, и 4 функции обычного юзера
//  в том числе управлять его листом (todo)
 
func main() {
	var (
		user_admin uint = 1;
		user_user  uint= 2;
		user_new   uint= 3;
    )

	var (
		key_create_user uint = 1
		key_delete_user uint = 2	
	)
	var (
		key_create_todo uint = 11
		key_view_todo uint = 12
		key_edit_todo uint = 13
		key_delete_todo uint = 14	
	)
	var user User

	switch api_key.id {
	case who.login(&user):
		//дальше надо как то ограничить кол-во функций (? ) мб выдаваемыми Api ключами
	case key_create_user:
		//
	case 

	//api key
}
