document.getElementsByTagName('form').item(0).addEventListener('submit', sendRequest());

loading = false
hideLoading();

function sendRequest() {
    return async function (e) {
        e.preventDefault();
        let links = document.getElementById('link').value.trim().split(" ");
        if (links.length > 0 && links[0] !== '' && !loading) {
            showLoading();
            // development:
            // let urlString = "http://localhost:8080/api?";
            // production:
            let urlString = "https://letterboxd-scraper.herokuapp.com/api?";

            links.forEach((link) => {
                urlString += "src=" + link + "&";
            })
            if (getType() === "intersection") {
                urlString += "i=true";
            }

            console.log("url: " + urlString)
            const response = await fetch(urlString);
            hideLoading();

            let movie = document.getElementById("movie-container");
            if (response.status === 200) {
                movie.innerHTML = ``;

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
                fileName = links.join().replace(/\//gm, "_") + ".csv";
                let fileToSave = new Blob([result], {
                    type: "csv",
                    name: fileName
                });
                saveAs(fileToSave, fileName);
            } else {
                movie.innerHTML = `<p id="missingMovie">Sorry that list does not exist.</p>`;
            }
        }
    }
}

function showLoading() {
    loading = true
    document.getElementById('submitButton').innerHTML = '<span id="spinner" class="spinner-border text-light spinner-border-sm" role="status" aria-hidden="true"></span>Loading...';
}

function hideLoading() {
    loading = false
    document.getElementById('submitButton').innerHTML = 'SUBMIT';
}

function getType() {
    if (document.querySelector('input[id="intersection"]:checked') !== null) {
        return "intersection";
    }
    return "union";
}