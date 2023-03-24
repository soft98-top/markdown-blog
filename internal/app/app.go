package app

import (
	"bufio"
	"github.com/robfig/cron"
	"html/template"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/gaowei-space/markdown-blog/internal/api"
	"github.com/gaowei-space/markdown-blog/internal/bindata/assets"
	"github.com/gaowei-space/markdown-blog/internal/bindata/views"
	"github.com/gaowei-space/markdown-blog/internal/types"
	"github.com/gaowei-space/markdown-blog/internal/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"github.com/urfave/cli/v2"
)

var (
	MdDir      string
	Env        string
	Title      string
	HeadTitle  string
	Index      string
	LayoutFile = "layouts/layout.html"
	LogsDir    = "cache/logs/"
	TocPrefix  = "[toc]"
	IgnoreFile = []string{`favicon.ico`, `.DS_Store`, `.gitignore`, `README.md`}
	IgnorePath = []string{`.git`}
	Cache      time.Duration
	Analyzer   types.Analyzer
	Gitalk     types.Gitalk
	Proxy      types.Proxy
	Sitemap    types.Sitemap
)

type urlNode struct {
	url  string `json:"url"`
	time string `json:"time"`
	name string `json:"name"`
}

// web服务器默认端口
const DefaultPort = 5006

func RunWeb(ctx *cli.Context) error {
	initParams(ctx)

	app := iris.New()

	setLog(app)

	tmpl := iris.HTML(views.AssetFile(), ".html").Reload(true)
	app.RegisterView(tmpl)
	app.OnErrorCode(iris.StatusNotFound, api.NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, api.InternalServerError)

	setIndexAuto := false
	if Index == "" {
		setIndexAuto = true
	}

	app.Use(func(ctx iris.Context) {
		activeNav := getActiveNav(ctx)

		navs, firstNav := getNavs(activeNav)

		firstLink := strings.TrimPrefix(firstNav.Link, "/")
		if setIndexAuto && Index != firstLink {
			Index = firstLink
		}

		// 设置 Gitalk ID
		Gitalk.Id = utils.MD5(activeNav)

		ctx.ViewData("Gitalk", Gitalk)
		ctx.ViewData("Analyzer", Analyzer)
		ctx.ViewData("Proxy", Proxy)
		ctx.ViewData("Title", Title)
		ctx.ViewData("HeadTitle", Title)
		ctx.ViewData("Nav", navs)
		ctx.ViewData("ActiveNav", activeNav)
		ctx.ViewLayout(LayoutFile)

		ctx.Next()
	})

	app.Favicon("./favicon.ico")
	app.HandleDir("/static", assets.AssetFile())
	if Sitemap.Domain != "" || Sitemap.ExistFile != "" {
		err := genSitemap()
		if err == nil {
			path := "/sitemap.xml"
			if Sitemap.Path != "" {
				path = Sitemap.Path
			}
			app.Get(path, iris.Cache(Cache), sitemapHandler)
			genRobots(path)
			app.Get("/robots.txt", iris.Cache(Cache), robotsHandler)
			app.Get("/rss.xml", iris.Cache(Cache), rssHandler)
			if Sitemap.Cron != "" && Sitemap.ExistFile == "" {
				c := cron.New()
				c.AddFunc(Sitemap.Cron, func() {
					app.Logger().Debugf("Run genSitemap...")
					genSitemap()
				})
				c.Start()
			}
		}
	}
	app.Get("/{f:path}", iris.Cache(Cache), articleHandler)

	app.Run(iris.Addr(":" + strconv.Itoa(parsePort(ctx))))

	return nil
}

func initParams(ctx *cli.Context) {
	MdDir = ctx.String("dir")
	if strings.TrimSpace(MdDir) == "" {
		log.Panic("Markdown files folder cannot be empty")
	}
	MdDir, _ = filepath.Abs(MdDir)

	Env = ctx.String("env")
	Title = ctx.String("title")
	Index = ctx.String("index")
	cache := int64(ctx.Int("cache"))
	Cache = time.Minute * time.Duration(cache)
	if Env == "dev" {
		Cache = time.Minute * 0
	}

	// 设置分析器
	Analyzer.SetAnalyzer(ctx.String("analyzer-baidu"), ctx.String("analyzer-google"), ctx.String("analyzer-google-ad"))

	// 设置Gitalk
	Gitalk.SetGitalk(ctx.String("gitalk.client-id"), ctx.String("gitalk.client-secret"), ctx.String("gitalk.repo"), ctx.String("gitalk.owner"), ctx.StringSlice("gitalk.admin"), ctx.StringSlice("gitalk.labels"))

	// 设置Proxy
	Proxy.SetProxy(ctx.String("proxy.github-api"), ctx.String("proxy.github-cors"), ctx.String("proxy.google-ad"), ctx.String("proxy.google-ay"))

	//设置Sitemap
	Sitemap.SetSitemap(ctx.String("sitemap.domain"), ctx.String("sitemap.tls"), ctx.String("sitemap.cron"), ctx.String("sitemap.path"), ctx.String("sitemap.exist-file"))

	// 忽略文件
	IgnoreFile = append(IgnoreFile, ctx.StringSlice("ignore-file")...)
	IgnorePath = append(IgnorePath, ctx.StringSlice("ignore-path")...)
}

func setLog(app *iris.Application) {
	os.MkdirAll(LogsDir, 0777)
	f, _ := os.OpenFile(LogsDir+"access-"+time.Now().Format("20060102")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)

	if Env == "prod" {
		app.Logger().SetOutput(f)
	} else {
		app.Logger().SetLevel("debug")
		app.Logger().Debugf(`Log level set to "debug"`)
	}

	// Close the file on shutdown.
	app.ConfigureHost(func(su *iris.Supervisor) {
		su.RegisterOnShutdown(func() {
			f.Close()
		})
	})

	ac := accesslog.New(f)
	ac.AddOutput(app.Logger().Printer)
	app.UseRouter(ac.Handler)
	app.Logger().Debugf("Using <%s> to log requests", f.Name())
}

func parsePort(ctx *cli.Context) int {
	port := DefaultPort
	if ctx.IsSet("port") {
		port = ctx.Int("port")
	}
	if port <= 0 || port >= 65535 {
		port = DefaultPort
	}

	return port
}

func getNavs(activeNav string) ([]map[string]interface{}, utils.Node) {
	var option utils.Option
	option.RootPath = []string{MdDir}
	option.SubFlag = true
	option.IgnorePath = IgnorePath
	option.IgnoreFile = IgnoreFile
	tree, _ := utils.Explorer(option)

	navs := make([]map[string]interface{}, 0)
	for _, v := range tree.Children {
		for _, item := range v.Children {
			searchActiveNav(item, activeNav)
			navs = append(navs, structs.Map(item))
		}
	}

	firstNav := getFirstNav(*tree.Children[0])

	return navs, firstNav
}

func searchActiveNav(node *utils.Node, activeNav string) {
	if !node.IsDir && node.Link == "/"+activeNav {
		node.Active = "active"
		return
	}
	if len(node.Children) > 0 {
		for _, v := range node.Children {
			searchActiveNav(v, activeNav)
		}
	}
}

func getFirstNav(node utils.Node) utils.Node {
	if !node.IsDir {
		return node
	}
	return getFirstNav(*node.Children[0])
}

func getActiveNav(ctx iris.Context) string {
	f := ctx.Params().Get("f")
	if f == "" {
		f = Index
	}
	return f
}

func articleHandler(ctx iris.Context) {
	f := getActiveNav(ctx)

	if utils.IsInSlice(IgnoreFile, f) {
		return
	}
	// 防止目录穿越
	if strings.Contains(f, "..") || strings.Contains(f, "./") {
		return
	}
	// 文件类型
	ftype := "md"

	fpath := MdDir + "/" + f
	// 判断文件后缀是否是.png
	if strings.HasSuffix(fpath, ".png") {
		ftype = "img"
	} else {
		fpath = fpath + ".md"
	}

	_, err := os.Stat(fpath)
	if err != nil {
		ctx.StatusCode(404)
		ctx.Application().Logger().Errorf("Not Found '%s', Path is %s", fpath, ctx.Path())
		return
	}

	bytes, err := os.ReadFile(fpath)
	if err != nil {
		ctx.StatusCode(500)
		ctx.Application().Logger().Errorf("ReadFile Error '%s', Path is %s", fpath, ctx.Path())
		return
	}

	if ftype == "img" {
		ctx.Header("Content-Type", "image/png")
		ctx.Write(bytes)
	} else {
		tmp := strings.Split(f, "/")
		title := tmp[len(tmp)-1]
		if strings.Contains(title, "@") {
			tmp = strings.Split(title, "@")
			title = tmp[len(tmp)-1]
		}
		ctx.ViewData("HeadTitle", title+" - "+Title)
		ctx.ViewData("Article", mdToHtml(bytes))
		ctx.View("index.html")
	}
}

func mdToHtml(content []byte) template.HTML {
	strs := string(content)

	var htmlFlags blackfriday.HTMLFlags

	if strings.HasPrefix(strs, TocPrefix) {
		htmlFlags |= blackfriday.TOC
		strs = strings.Replace(strs, TocPrefix, "<br/><br/>", 1)
	}

	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: htmlFlags,
	})

	unsafe := blackfriday.Run([]byte(strs), blackfriday.WithRenderer(renderer), blackfriday.WithExtensions(blackfriday.CommonExtensions))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	return template.HTML(string(html))
}

func genRSS(urls []urlNode, domain string) {
	var rssTemplate = `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">

<channel>
  <title>${site_title}</title>
  <link>${domain}</link>
  <description>${site_desc}</description>
  ${items}
</channel>

</rss>`
	var itemTemlate = `<item>
    <title>${title}</title>
    <link>${url}</link>
    <description>${desc}</description>
  </item>`
	var articleArray []string
	for _, url := range urls {
		articleStr := os.Expand(itemTemlate, func(s string) string {
			switch s {
			case "url":
				return url.url
			case "title":
				return url.name
			case "desc":
				return url.name
			}
			return ""
		})
		articleArray = append(articleArray, articleStr)
	}
	itemsStr := strings.Join(articleArray, "\n  ")
	rssStr := os.Expand(rssTemplate, func(s string) string {
		switch s {
		case "site_title":
			return Title
		case "domain":
			return domain
		case "site_desc":
			return Title
		case "items":
			return itemsStr
		}
		return ""
	})
	filePath := "./rss.xml"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(rssStr)
	write.Flush()
}

func genRobots(path string) {
	sitemapUrl := Sitemap.Domain + path
	if strings.ToLower(Sitemap.Tls) == "true" {
		sitemapUrl = "https://" + sitemapUrl
	} else {
		sitemapUrl = "http://" + sitemapUrl
	}
	robotsStr := `User-agent: *
Allow: /

Sitemap: ` + sitemapUrl
	filePath := "./robots.txt"
	file, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(robotsStr)
	write.Flush()
}

func genSitemap() error {
	var sitemapTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9 http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
  ${urls}
</urlset>`
	var urlTemplate = `<url>
	<loc>${url}</loc>
	<lastmod>${time}</lastmod>
  </url>`
	var domain string
	if strings.ToLower(Sitemap.Tls) == "true" {
		domain = "https://" + Sitemap.Domain + "/"
	} else {
		domain = "http://" + Sitemap.Domain + "/"
	}
	var option utils.Option
	option.RootPath = []string{MdDir}
	option.SubFlag = true
	option.IgnorePath = IgnorePath
	option.IgnoreFile = IgnoreFile
	tree, _ := utils.Explorer(option)
	var urls []urlNode
	for _, v := range tree.Children {
		getUrls(&urls, domain, v)
	}
	var urlsArray []string
	for _, url := range urls {
		urlStr := os.Expand(urlTemplate, func(s string) string {
			switch s {
			case "url":
				return url.url
			case "time":
				return url.time
			}
			return ""
		})
		urlsArray = append(urlsArray, urlStr)
	}
	urlsStr := strings.Join(urlsArray, "\n  ")
	sitemapStr := os.Expand(sitemapTemplate, func(s string) string {
		switch s {
		case "urls":
			return urlsStr
		}
		return ""
	})
	filePath := "./sitemap.xml"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(sitemapStr)
	write.Flush()
	go genRSS(urls, domain)
	return err
}

func getUrls(urls *[]urlNode, prefix string, node *utils.Node) {
	for _, item := range node.Children {
		if !item.IsDir && strings.HasSuffix(item.Name, ".md") {
			info, _ := os.Stat(item.Path)
			names := strings.Split(item.Name, "@")
			var name string
			if len(names) > 1 {
				name = strings.TrimSuffix(names[1], ".md")
			} else {
				name = strings.TrimSuffix(item.Name, ".md")
			}
			url := urlNode{
				url:  prefix + url.PathEscape(strings.TrimSuffix(item.Name, ".md")),
				time: info.ModTime().Format("2006-01-02 15:04:05"),
				name: name,
			}
			*urls = append(*urls, url)
		} else {
			getUrls(urls, prefix+url.PathEscape(item.Name)+"/", item)
		}
	}
}

func sitemapHandler(ctx iris.Context) {
	fpath := "./sitemap.xml"
	if Sitemap.ExistFile != "" {
		fpath = Sitemap.ExistFile
	}
	ctx.SendFile(fpath, "sitemap.xml")
}
func robotsHandler(ctx iris.Context) {
	fpath := "./robots.txt"
	ctx.SendFile(fpath, "robots.txt")
}
func rssHandler(ctx iris.Context) {
	fpath := "./rss.xml"
	ctx.SendFile(fpath, "rss.xml")
}
