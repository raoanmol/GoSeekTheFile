document.getElementById('search-form').addEventListener('submit', function(e) {
    e.preventDefault();

    var query = document.getElementById('search-input').value;
    var resultsElement = document.getElementById('results');

    fetch('/search', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query: query }),
    })
        .then(function(response) {
            return response.json();
        })
        .then(function(data) {
            resultsElement.innerHTML = '';

            data.files.forEach(function(file) {
                var li = document.createElement('li');
                li.textContent = file;
                resultsElement.appendChild(li);
            });
        });
});

let themeCheckbox = document.querySelector('input[type="checkbox"]');
themeCheckbox.addEventListener('change', function (event) {
    if (event.target.checked) {
        document.getElementById('theme-style').href = './dark.css';
    } else {
        document.getElementById('theme-style').href = './light.css';
    }
});