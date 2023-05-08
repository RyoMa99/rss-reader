package constant

import "time"

type Site struct {
	Name           string
	URL            string
	DateTimeFormat string
}

var SITE_LIST = []Site{
	{
		Name:           "メルカリ",
		URL:            "https://engineering.mercari.com/blog/feed.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "メドレー",
		URL:            "https://developer.medley.jp/rss.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "FiNC",
		URL:            "https://medium.com/feed/finc-engineering",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "みてね",
		URL:            "https://team-blog.mitene.us/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "pinterest",
		URL:            "https://medium.com/feed/@Pinterest_Engineering",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "netflix",
		URL:            "https://netflixtechblog.com/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "kawasima",
		URL:            "https://scrapbox.io/api/feed/kawasima",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "azukiazusa のテックブログ2",
		URL:            "https://azukiazusa.dev/rss.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "hiroppy",
		URL:            "https://hiroppy.me/blog/rss.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "ASnoKaze blog",
		URL:            "https://asnokaze.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "SIOS セキュリティブログ",
		URL:            "https://security.sios.jp/feed/",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "千里霧中",
		URL:            "https://goyoki.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "ぷらすのブログ",
		URL:            "https://blog.p1ass.com/index.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "valid,invalid",
		URL:            "https://ohbarye.hatenablog.jp/rss",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Voicy Tech Blog",
		URL:            "https://tech-blog.voicy.jp/rss",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "SST バックヤード",
		URL:            "https://techblog.securesky-tech.com/rss",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "RyotaK's Blog",
		URL:            "https://blog.ryotak.net/index.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "JSer.info",
		URL:            "https://jser.info/rss/",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn人気記事",
		URL:            "https://zenn.dev/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics AWS",
		URL:            "https://zenn.dev/topics/aws/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics Linux",
		URL:            "https://zenn.dev/topics/linux/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics Docker",
		URL:            "https://zenn.dev/topics/docker/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics k8s",
		URL:            "https://zenn.dev/topics/kubernetes/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics Security",
		URL:            "https://zenn.dev/topics/security/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics Go",
		URL:            "https://zenn.dev/topics/go/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Golang News",
		URL:            "https://golangnews.com/index.xml",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "Zenn Topics FrontEnd",
		URL:            "https://zenn.dev/topics/frontend/feed",
		DateTimeFormat: time.RFC1123,
	},
	{
		Name:           "iCARE",
		URL:            "https://dev.icare.jpn.com/dev_cat/feed/",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "クラスメソッド",
		URL:            "https://dev.classmethod.jp/feed/",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "ラック",
		URL:            "https://devblog.lac.co.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "CyberAgent",
		URL:            "https://developers.cyberagent.co.jp/blog/feed/",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "DeNA",
		URL:            "https://engineering.dena.com/blog/index.xml",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "スタディプラス",
		URL:            "https://tech.studyplus.co.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "BASE",
		URL:            "https://devblog.thebase.in/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "CADDi",
		URL:            "https://caddi.tech/feed",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "エウレカ",
		URL:            "https://medium.com/feed/eureka-engineering",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "dely",
		URL:            "https://tech.dely.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "アスクル",
		URL:            "https://tech.askul.co.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "STORES",
		URL:            "https://product.st.inc/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "クラウドワークス",
		URL:            "https://engineer.crowdworks.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "オイシックス",
		URL:            "https://creators.oisix.co.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "エムスリー",
		URL:            "https://www.m3tech.blog/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "Github",
		URL:            "https://github.blog/feed/",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "継続は力なり",
		URL:            "https://sadayoshi-tada.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "Carpe Diem",
		URL:            "https://christina04.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "Please Sleep",
		URL:            "https://please-sleep.cou929.nu/rss.xml",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "yohgaki's blog",
		URL:            "https://blog.ohgaki.net/feed",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "pospomeのプログラミング日記",
		URL:            "https://www.pospome.work/feed",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "まっちゃだいふくの日記",
		URL:            "https://ripjyr.hatenablog.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "piyolog",
		URL:            "https://piyolog.hatenadiary.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "かとじゅんの技術日誌",
		URL:            "https://blog.j5ik2o.me/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "年中アイス",
		URL:            "https://reiki4040.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "理系学生日記",
		URL:            "https://kiririmode.hatenablog.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "Big Sky",
		URL:            "https://mattn.kaoriya.net/index.rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "赤帽エンジニアブログ",
		URL:            "https://rheb.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "地方エンジニアの学習日記",
		URL:            "https://ryuichi1208.hateblo.jp/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "kakakakakku blog",
		URL:            "https://kakakakakku.hatenablog.com/rss",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "Golang Weekly News",
		URL:            "https://cprss.s3.amazonaws.com/golangweekly.com.xml",
		DateTimeFormat: time.RFC1123Z,
	},
	{
		Name:           "フューチャー",
		URL:            "https://future-architect.github.io/atom.xml",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "GMOペパボ",
		URL:            "https://tech.pepabo.com/feed.xml",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "ラクス",
		URL:            "https://tech-blog.rakus.co.jp/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "DeNA SWET",
		URL:            "https://swet.dena.com/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "サイボウズ",
		URL:            "https://blog.cybozu.io/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "miki_beneのセキュリティ情報のタイムライン - 大チェッカー",
		URL:            "https://daichkr.hatelabo.jp/antenna/2261401675296895784/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "徳丸浩の日記",
		URL:            "https://blog.tokumaru.org/feeds/posts/default",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "Web Scratch",
		URL:            "https://efcl.info/feed/",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "JVN",
		URL:            "https://jvn.jp/rss/jvn.rdf",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "Security NEXT",
		URL:            "https://www.security-next.com/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "窓の社",
		URL:            "https://forest.watch.impress.co.jp/data/rss/1.0/wf/feed.rdf",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "技術評論社",
		URL:            "https://gihyo.jp/dev/feed/atom",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "Qiita人気記事",
		URL:            "https://qiita.com/popular-items/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "ジンジャー研究所",
		URL:            "https://jinjor-labo.hatenablog.com/feed",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "ZDNet Japan 最新情報 総合",
		URL:            "https://www.watch.impress.co.jp/data/rss/1.0/ipw/feed.rdf",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "Impress Watch",
		URL:            "http://feeds.japan.zdnet.com/rss/zdnet/all.rdf",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "Neo's world",
		URL:            "https://neos21.net/feeds.xml",
		DateTimeFormat: time.RFC3339,
	},
	{
		Name:           "帰ってきた「しっぽのさきっちょ」",
		URL:            "https://text.baldanders.info/index.xml",
		DateTimeFormat: time.RFC3339,
	},
}

// rss未対応
// "https://www.kimullaa.com/"
