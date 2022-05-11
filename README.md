# athome-scraper
A tool to scrape and visualise apartments from athome.lu

Related article can be found [here](https://mattrighetti.com/2022/04/05/i-need-to-find-an-appartment.html)

## Dependencies
For this project you will need a few binaries like
1. [just](https://github.com/casey/just)
2. Jupyter notebooks

## Usage

1. Clone the project
```sh
$ git clone git@github.com:mattrighetti/athome-scraper.git
```

2. Create a `homes.txt` file in the project folder
```sh
$ touch homes.txt
```

3. Paste links to `homes.txt`, each link on a new line
```sh
https://www.athome.lu/en/rent/apartment/luxembourg/id-7486621.html
https://www.athome.lu/en/rent/house/luxembourg/id-7511396.html
...
https://www.athome.lu/en/rent/house/strassen/id-7512843.html
```

4. Run script with just

```sh
$ just show
```
