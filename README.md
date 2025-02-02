# Web Page Analyzer

Is a simple web application built in Go that allows users to analyze web pages/URLs and retrieve details about the structure and contents.
it performs below tasks,
- Detects the HTML version of the page.
- Extracts the title of the page.
- Counts the number of headings (`h1`, `h2`, `h3`, etc.) on the page.
- Analyzes internal and external links, reporting inaccessible links as well.
- Checks if the page contains a login form.

The application allows users to input a URL through a web interface and returns an analysis of the page.


## Setup Instructions


### Prerequisites
To run this project, you'll need:
- **Go**: A working Go installation. You can download Go from the [official website](https://golang.org/dl/).
- **Git**: If you want to clone the repository and manage version control.

### Installing Dependencies
To manage project dependencies, you will need to initialize Go modules if you haven’t done so already.

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/web-page-analyzer.git
   cd web-page-analyzer
   ```

2. Install Dependencies:
   ```
   go mod tidy
   ```

External Dependencies,
- github.com/PuerkitoBio/goquery
- github.com/chromedp/chromedp
- github.com/stretchr/testify

To Install indiviually,
`go get <dependency>`

### Build & Run

Run below command.

To Build,
```
go build .
```

To Run,
```
go run main.go
```

Open `http://localhost:8080/` in browser to confirm.

### Test

`go test -v ./...` to run unit tests.
`go test -cover ./...` for coverage.

### Usage

For a given URL `https://www.example.com`, Results:

```
Analysis Results

URL: https://www.example.com
HTML Version: HTML5
Title: Example Page
Headings:
  h1: 1
  h2: 3
  h3: 2
Internal Links: 15
External Links: 10
Inaccessible Links: 2
Login Form Present: No
```


### Run Test

```
go test
```

###

## Make Integrated.
`make help` will show available commands.

## Possible Improvements

- Use Goroutines to parellelize data extraction from html document like title, links etc.
- Pooling workers for accessible link check.
- Data Persistance/Caching for history & Performance. 

- Streaming API response. some fetch are easy compare check broken links. to improve user experience, we can stream data from API.
- FrontEnd Improvements. e.g: React or Styling.
- Include Page Performance metrics. e.g: latency, load time.
- CI/CD for this Project. e.g: Automated build, test(+static-sonar) and deploy.
- Load balancing and distribute requests.
