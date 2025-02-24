document.getElementById('loginForm').addEventListener('submit', function(event) {
    event.preventDefault();  // Предотвращаем перезагрузку страницы при отправке формы

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    // Отправляем запрос на сервер для получения токена
    fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name: username, password: password })
    })
        .then(response => response.json())
        .then(data => {
            if (data.token) {
                // Сохраняем токен в localStorage
                localStorage.setItem('token', data.token);
                alert('Вы успешно вошли!');

                // После авторизации скрываем форму логина и показываем посты
                document.getElementById('login-form').style.display = 'none';
                document.getElementById('posts-section').style.display = 'block';
                getPosts();  // Загрузка постов

                // Показываем кнопку Admin
                document.getElementById('adminButton').style.display = 'inline-block';  // Показываем кнопку сразу

                // Показываем кнопку Logout
                document.getElementById('logoutButton').style.display = 'inline-block';  // Показываем кнопку logout
            } else {
                alert('Ошибка авторизации');
            }
        })
        .catch(error => console.error('Ошибка:', error));
});

// Функция для выхода из системы
function logout() {
    // Удаляем токен из localStorage
    localStorage.removeItem('token');

    // Скрываем кнопки Admin и Logout
    document.getElementById('adminButton').style.display = 'none';
    document.getElementById('logoutButton').style.display = 'none';

    // Показываем форму логина
    document.getElementById('login-form').style.display = 'block';

    // Скрываем раздел с постами
    document.getElementById('posts-section').style.display = 'none';
}
