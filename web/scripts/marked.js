
function getPosts(page = 1) {
    fetch(`http://localhost:8080/?page=${page}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    })
        .then(response => response.json())
        .then(data => {
            const postsList = document.getElementById('posts-list');
            const postsSection = document.getElementById('posts-section');
            const pagination = document.getElementById('pagination');

            // Показываем блок с постами
            postsSection.style.display = 'block';

            // Очищаем предыдущие посты
            postsList.innerHTML = '';

            // Отображаем посты
            data.posts.forEach(post => {
                const li = document.createElement('li');
                li.classList.add('post');

                // Преобразуем текст поста из Markdown в HTML с помощью библиотеки marked.js
                const postContent = marked.parse(post.text);  // Преобразуем Markdown в HTML
                const postTitleMarked = marked.parse(post.title);

                li.innerHTML = `
                    <strong style ="font-size: 18px; display:flex; justify-content: center;">${postTitleMarked}</strong><br/>
                    ${postContent}<br/>
                `;
                postsList.appendChild(li);
            });

            // Обновляем пагинацию
            pagination.innerHTML = '';

            if (data.hasPrev) {
                const prevButton = document.createElement('button');
                prevButton.innerText = 'Previous';
                prevButton.onclick = () => getPosts(page - 1); // Переход к предыдущей странице
                pagination.appendChild(prevButton);
            }

            if (data.hasNext) {
                const nextButton = document.createElement('button');
                nextButton.innerText = 'Next';
                nextButton.onclick = () => getPosts(page + 1); // Переход к следующей странице
                pagination.appendChild(nextButton);
            }
        })
        .catch(error => console.error('Error:', error));
}