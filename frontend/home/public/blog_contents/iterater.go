package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ArticleMetadata struct {
	Title string `json:"title"`
	Uri   string `json:"uri"`
	Date  string `json:"date"`
}

const (
	EXT_HTML          string = ".html"
	EXT_MD            string = ".md"
	DirPrefix         string = "../../public/blog_contents"
	PathBlogComponent string = "../../src/components/Blog.vue"
	PathVueConfigJs   string = "../../vue.config.js"
)

var articles []string
var prerenderingArticles = []string{
	"/",
	"/blog/",
	"/daily/",
}

func execRmOldHtml(path string) {
	if (len(path) > len(EXT_HTML)) &&
		path[len(path)-len(EXT_HTML):] == EXT_HTML {

		cmd := exec.Command("rm", path)
		// cmd := exec.Command("pwd")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("pandoc output: %q\n", out.String())
		log.Printf("deleted html document: %q\n", path)
	}
}

func execPandocMdToHtml(path string) {
	// Length of a file have to be logner than extension
	if (len(path) > len(EXT_MD)) &&
		path[len(path)-len(EXT_MD):] == EXT_MD {

		result := path[0:len(path)-len(EXT_MD)] + EXT_HTML

		cmd := exec.Command(
			"pandoc",
			path,
			"-o",
			result,
			"--to",
			"html5",
			//"--standalone",
			"--no-highlight",
		)
		// cmd := exec.Command("pwd")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		//log.Printf("pandoc output: %q\n", out.String())

		articles = append(articles, path[0:len(path)-len(EXT_MD)]+EXT_HTML)

		log.Printf("generated document: %q\n", result)
	}
}

func iterate(path string, dirInfo []os.FileInfo) {
	for _, info := range dirInfo {
		if info.Name()[0:2] == "aA" {
			continue
		}

		absolutePathName := path + "/" + info.Name()

		if info.IsDir() {
			subDirInfo, err := ioutil.ReadDir(absolutePathName)
			if err != nil {
				log.Println("(iterate) ", err)
				log.Fatal(err)
			}
			iterate(absolutePathName, subDirInfo)
		} else {
			execRmOldHtml(absolutePathName)
			execPandocMdToHtml(absolutePathName)
		}
		fmt.Println(info.Name())
	}
}

func GetArticleMetadata(htmlPaths []string) []ArticleMetadata {

	r := regexp.MustCompile("(19|20)[0-9][0-9][- /.](0[1-9]|1[012])[- /.](0[1-9]|[12][0-9]|3[01])")

	var metadata []ArticleMetadata

	for _, htmlPath := range htmlPaths {
		// Get title
		htmlSource, err := ioutil.ReadFile(htmlPath)
		if err != nil {
			log.Fatal(err.Error())
		}

		from := bytes.Index(htmlSource, []byte(">"))
		from += len(">")
		to := bytes.Index(htmlSource, []byte("</h1>"))

		// Get Uri
		uri := strings.Replace(htmlPath, "../../public/blog_contents/", "/blog/", 1)
		uri = strings.Replace(uri, ".html", "/", 1)

		// Get date
		date := r.FindString(htmlPath)
		date = strings.Replace(date, "/", ".", 2)
		date += "."

		metadata = append(metadata, ArticleMetadata{
			Title: string(htmlSource[from:to]),
			Uri:   uri,
			Date:  date,
		})
	}

	return metadata
}

func insertJsonToBlogComponent(json string) {
	componentSourceBytes, err := ioutil.ReadFile(PathBlogComponent)
	if err != nil {
		log.Fatal(err)
	}
	componentSource := string(componentSourceBytes)

	i := strings.Index(componentSource, "__INSERTION_POSITION__")

	from := strings.Index(componentSource[i:], "[")
	from += i

	k := strings.Index(componentSource, "__INSERTION_POSITION_END__")
	to := strings.LastIndex(componentSource[i:k+2], "]")
	to += i

	fmt.Println(" -- old json -- ")
	fmt.Println(string(componentSource[from : to+1]))
	fmt.Println("-- old json end -- ")

	componentSource = componentSource[0:from] + json + componentSource[to+1:]
	fmt.Println(componentSource)

	ioutil.WriteFile(PathBlogComponent, []byte(componentSource), 0644)
}

func insertJsonToVueConfigJs(json string) {
	json = strings.Replace(json, "]", ",", 1)

	vueConfigJsSourceBytes, err := ioutil.ReadFile(PathVueConfigJs)
	if err != nil {
		log.Fatal(err)
	}

	vueConfigJsSource := string(vueConfigJsSourceBytes)

	i := strings.Index(vueConfigJsSource, "__INSERTION_POSITION__")

	from := strings.Index(vueConfigJsSource[i:], "[")
	from += i

	k := strings.Index(vueConfigJsSource, "__INSERTION_POSITION_END__")
	to := strings.LastIndex(vueConfigJsSource[i:k+2], "\"/insight/\"")
	to += i

	fmt.Println(" -- old json -- ")
	fmt.Println(string(vueConfigJsSource[from : to+1]))
	fmt.Println("-- old json end -- ")

	vueConfigJsSource = vueConfigJsSource[0:from] + json + vueConfigJsSource[to:]
	fmt.Println(vueConfigJsSource)

	ioutil.WriteFile(PathVueConfigJs, []byte(vueConfigJsSource), 0644)
}

func main() {
	dirInfo, err := ioutil.ReadDir(DirPrefix)
	if err != nil {
		log.Println("(main) ", err)
		log.Fatal(err)
	}
	iterate(DirPrefix, dirInfo)

	fmt.Println("\n\n articles \n\n")
	for _, article := range articles {
		fmt.Println(article)
	}

	// Reverse articles' order
	i := 0
	k := len(articles) - 1
	for {
		articles[i], articles[k] = articles[k], articles[i]

		i++
		k--
		if i >= k {
			break
		}
	}

	metadata := GetArticleMetadata(articles)

	b, err := json.Marshal(metadata)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n\n json \n\n")
	fmt.Println(string(b))

	insertJsonToBlogComponent(string(b))

	//prerendering
	for _, articlePath := range articles {
		uri := strings.Replace(articlePath, "../../public/blog_contents/", "/blog/", 1)
		uri = strings.Replace(uri, ".html", "/", 1)
		prerenderingArticles = append(prerenderingArticles, uri)
	}
	u, err := json.Marshal(prerenderingArticles)
	if err != nil {
		log.Fatal(err)
	}
	insertJsonToVueConfigJs(string(u))
}
