function loginAccess() {
	let button = document.getElementById('submit')
	let login = document.getElementById('Username')
	let password = document.getElementById('Password')
	if ((login = 14299) && (password = '1')) {
		let notif = document.createElement('div')
		document.body.prepend(notif)
		notif.className = 'active'
		notif.textContent = 'Авторизация прошла успешно'
		setTimeout(window.location = '../Home_page/home.html', 70000)
	}
}