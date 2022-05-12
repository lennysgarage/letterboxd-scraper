document.getElementsByTagName('form').item(0).addEventListener('submit', sendRequest());

hideLoading();

function sendRequest() {
    return async function(e) {
        e.preventDefault();
        let link = document.getElementById('link').value;
        if (link !== '') {
            showLoading();
            // development: http://localhost:8080/api?src=${link}
            // production: https://letterboxd-scraper.herokuapp.com/api?src=${link}
            const response = await fetch(`https://letterboxd-scraper.herokuapp.com/api?src=${link}`);
            hideLoading();
            let movie = document.getElementById("movie-container");
            if (response.status === 200) {
                const arr = await response.json();
                arr.unshift(["Title", "LetterboxdURI"])

                // Source: https://embed.plnkr.co/aO2DwSpmZCQ0GHSmPv7m/
                function ConvertToCSV(objArray) {
                    var array = typeof objArray != 'object' ? JSON.parse(objArray) : objArray;
                    var str = '';
        
                    for (var i = 0; i < array.length; i++) {
                        var line = '';
                        for (var index in array[i]) {
                            if (line != '') line += ','
        
                            line += array[i][index];
                        }
        
                        str += line + '\r\n';
                    }
        
                    return str;
                }

                result = ConvertToCSV(arr);
                let fileToSave = new Blob([result], {
                    type: "csv",
                    name: link.replace(/\//gm, "_") + ".csv"
                });
                saveAs(fileToSave, link.replace(/\//gm, ":") + ".csv");
            } else {
                movie.innerHTML = `<p id="missingMovie">Sorry that list does not exist.</p>`;
            }
        }
    }
}

function showLoading() {
    document.getElementById('submitButton').innerHTML = '<span id="spinner" class="spinner-border text-light spinner-border-sm" role="status" aria-hidden="true"></span>Loading...';
}

function hideLoading() {
    document.getElementById('submitButton').innerHTML = 'SUBMIT';
}