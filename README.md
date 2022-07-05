# letterboxd-scraper

Download a csv file of a letterboxd watchlist or public list 

## Usage
- Clone repo
    ```
    git clone https://github.com/lennysgarage/letterboxd-scraper.git
    ```
- Get packages
    ```
    cd letterboxd-scraper/backend
    go get
    ```
- Run command line tool
    ```
    go build
    ./letterboxd-scraper <list-link> <link-link-2> ...
    ```
- Run server tool
    ```
    cd letterboxd-scraper/backend/server
    go build
    ./server
    ```

## Example
    ./letterboxd-scraper https://letterboxd.com/lennysgarage/watchlist/
    ./letterboxd-scraper https://letterboxd.com/lennysgarage/list/movies-from-high-school/
