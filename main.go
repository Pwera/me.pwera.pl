package main

import (
	"embed"
	"fmt"
	"github.com/a-h/templ"
	"io/fs"
	"net/http"
	"regexp"
)

//go:embed static/*
var content embed.FS

func handler() http.Handler {
	fsys := fs.FS(content)
	html, _ := fs.Sub(fsys, "static")
	return http.FileServer(http.FS(html))
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", handler()))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var pg PageData
		pg.ContainersAndKubernetes = []RepoLink{
			{Repo: "https://github.com/Pwera/Docker-Notes", Img: img("docker.png"), Name: "Docker"},
			{Repo: "https://github.com/Pwera/Kubernetes-Notes", Img: img("kube.svg"), Name: "Kubernetes"},
			{Img: img("helm.svg")},
			{Img: img("istio.png")},
			{Repo: "https://github.com/Pwera/Traefik-Notes", Img: img("traefik.png")},
			{Repo: "https://github.com/Pwera/Serverless-Notes", Img: img("openwhisk.png")},
		}
		pg.Devops = []RepoLink{
			{Repo: "https://github.com/Pwera/Fluentd-Notes", Img: img("fluent.png")},
			{Repo: "https://github.com/Pwera/Anchore-Notes", Img: img("anchore.png"), Name: "Anchore"},
			{Repo: "https://github.com/Pwera/Prometheus-Notes", Img: img("prometheus.svg"), Name: "Prometheus"},
			{Repo: "https://github.com/Pwera/Terraform-Notes", Img: img("terraform.png")},
			{Img: img("appdynamics.png")},
			{Img: img("splunk.svg")},
		}

		pg.Books = []RepoLink{
			oreilly2("https://learning.oreilly.com/library/view/learning-go/9781492077206/"),
			oreilly2("https://www.oreilly.com/library/view/http2-in-action/9781617295164/"),
			oreilly2("https://www.oreilly.com/library/view/certified-kubernetes-application/9781492083726/"),
			oreilly2("https://www.oreilly.com/library/view/100-go-mistakes/9781617299599/"),
			oreilly2("https://www.oreilly.com/library/view/go-for-devops/9781801818896/"),
			oreilly2("https://learning.oreilly.com/library/view/container-security/9781492056690/"),
			oreilly2("https://www.oreilly.com/library/view/learn-kubernetes-security/9781839216503/"),
		}

		pg.Books2 = []RepoLink{
			oreilly2("https://www.oreilly.com/library/view/certified-kubernetes-administrator/9781098107215/"),
			oreilly2("https://www.oreilly.com/library/view/the-kubernetes-operator/9781803232850/"),
			oreilly2("https://www.oreilly.com/library/view/cloud-native-devops/9781492040750/"),
			oreilly2("https://learning.oreilly.com/library/view/cloud-native-observability-with/9781801077705/"),
			oreilly2("https://learning.oreilly.com/library/view/networking-and-kubernetes/9781492081647/"),
			oreilly2("https://learning.oreilly.com/library/view/certified-kubernetes-security/9781098132965/"),
			oreilly2("https://learning.oreilly.com/library/view/kubernetes-security-and/9781098107093/"),
			oreilly2("https://learning.oreilly.com/library/view/kubernetes-best-practices/9781098142155/"),
			{Repo: "https://www.amazon.pl/DevOps-Containers-Security-Monitoring-English/dp/9389423538", Img: "https://images-na.ssl-images-amazon.com/images/I/41IVS-yNotL._SX403_BO1,204,203,200_.jpg"},
		}

		pg.WantToRead = []RepoLink{
			oreilly2("https://learning.oreilly.com/library/view/communication-patterns/9781098140533/"),
			oreilly2("https://learning.oreilly.com/library/view/network-programmability-and/9781098110826/"),
			oreilly2("https://learning.oreilly.com/library/view/prometheus-up/9781098131135/"),
			oreilly2("https://learning.oreilly.com/library/view/learning-ebpf/9781098135119/"),
			oreilly2("https://learning.oreilly.com/library/view/practical-opentelemetry-adopting/9781484290750/"),
			oreilly2("https://learning.oreilly.com/library/view/podman-in-action/9781633439689/"),
			oreilly2("https://learning.oreilly.com/library/view/kubernetes-secrets-management/9781617298912/"),
			oreilly2("https://learning.oreilly.com/library/view/microservice-apis/9781617298417/"),
			oreilly2("https://learning.oreilly.com/library/view/application-security-program/9781633439818/"),
			oreilly2("https://learning.oreilly.com/library/view/gitops-cookbook/9781492097464/"),
			oreilly2("https://learning.oreilly.com/library/view/kubernetes-programming-with/9781484290262/"),
			oreilly2("https://learning.oreilly.com/library/view/practical-gitops-infrastructure/9781484286739/"),
			oreilly2("https://learning.oreilly.com/library/view/serverless-beyond-the/9781484287613/"),
			oreilly2("https://learning.oreilly.com/library/view/overcoming-it-complexity/9781492098485/"),
			oreilly2("https://learning.oreilly.com/library/view/efficient-go/9781098105709/"),
			oreilly2("https://learning.oreilly.com/library/view/making-sense-of/9781617298004/"),
			oreilly2("https://learning.oreilly.com/library/view/a-developers-essential/9781803234366/"),
			oreilly2("https://learning.oreilly.com/library/view/managing-kubernetes-resources/9781803242897/"),
			oreilly2("https://learning.oreilly.com/library/view/practical-cloud-native/9781098118563/"),
			oreilly2("https://learning.oreilly.com/library/view/practical-linux-devops/9781484283189/"),
			oreilly2("https://learning.oreilly.com/library/view/learn-wireshark/9781803231679/"),
			oreilly2("https://learning.oreilly.com/library/view/lean-devops-a/9780133853643/"),
			oreilly2("https://learning.oreilly.com/library/view/kubernetes-application-developer/9781484280324/"),
			oreilly2("https://learning.oreilly.com/library/view/skills-of-a/9781617299704/"),
			oreilly2("https://learning.oreilly.com/library/view/rancher-deep-dive/9781803246093/"),
			oreilly2("https://learning.oreilly.com/library/view/hacking-apis/9781098130244/"),
			oreilly2("https://learning.oreilly.com/library/view/getting-started-with/9781484281277/"),
			oreilly2("https://learning.oreilly.com/library/view/rust-web-development/9781800561304/"),
			oreilly2("https://learning.oreilly.com/library/view/core-kubernetes/9781617297557/"),
			oreilly2("https://learning.oreilly.com/library/view/machine-learning-on/9781803241807/"),
			oreilly2("https://learning.oreilly.com/library/view/network-programming-with/9781484280959/"),
			oreilly2("https://learning.oreilly.com/library/view/designing-apis-with/9781617296284/"),
			oreilly2("https://learning.oreilly.com/library/view/consul-up-and/9781098106133/"),
			oreilly2("https://learning.oreilly.com/library/view/elasticsearch-8x-cookbook/9781801079815/"),
			oreilly2("https://learning.oreilly.com/library/view/software-architecture-metrics/9781098112226/"),
			oreilly2("https://learning.oreilly.com/library/view/software-mistakes-and/9781617299209/"),
			oreilly2("https://learning.oreilly.com/library/view/observability-engineering/9781492076438/"),
			oreilly2("https://learning.oreilly.com/library/view/continuous-delivery-with/9781803237480/"),
			oreilly2("https://learning.oreilly.com/library/view/podman-for-devops/9781803248233/"),
			oreilly2("https://learning.oreilly.com/library/view/a-complete-guide/9781484281178/"),
			oreilly2("https://learning.oreilly.com/library/view/hugo-in-action/9781617297007/"),
			oreilly2("https://learning.oreilly.com/library/view/learning-modern-linux/9781098108939/"),
			oreilly2("https://learning.oreilly.com/library/view/logging-in-action/9781617298356/"),
			oreilly2("https://learning.oreilly.com/library/view/istio-in-action/9781617295829/"),
			oreilly2("https://learning.oreilly.com/library/view/learning-devops/9781801818964/"),
			oreilly2("https://learning.oreilly.com/library/view/getting-started-with/9781800569492/"),
			oreilly2("https://learning.oreilly.com/library/view/kubernetes-native-development/9781484279427/"),
			oreilly2("https://learning.oreilly.com/library/view/wireshark-fundamentals-a/9781484280027/"),
			oreilly2("https://learning.oreilly.com/library/view/the-kubernetes-bible/9781838827694/"),
			oreilly2("https://learning.oreilly.com/library/view/efficient-linux-at/9781098113391/"),
			oreilly2("https://learning.oreilly.com/library/view/kafka-in-action/9781617295232/"),
			oreilly2("https://learning.oreilly.com/library/view/cassandra-the-definitive/9781492097136/"),
			oreilly2("https://learning.oreilly.com/library/view/street-coder/9781617298370/"),
			oreilly2("https://learning.oreilly.com/library/view/command-line-rust/9781098109424/"),
			oreilly2("https://learning.oreilly.com/library/view/practical-go/9781119773818/"),
			oreilly2("https://learning.oreilly.com/library/view/rust-for-rustaceans/9781098129828/"),
			oreilly2("https://learning.oreilly.com/library/view/apache-pulsar-in/9781617296888/"),
			oreilly2("https://learning.oreilly.com/library/view/powerful-command-line-applications/9781680509311/"),
			oreilly2("https://learning.oreilly.com/library/view/mastering-apache-pulsar/9781492084891/"),
			oreilly2("https://learning.oreilly.com/library/view/simplifying-service-management/9781800202627/"),
			oreilly2("https://learning.oreilly.com/library/view/designing-microservices-platforms/9781801072212/"),
			oreilly2("https://learning.oreilly.com/library/view/linux-for-networking/9781800202399/"),
			oreilly2("https://learning.oreilly.com/library/view/kafka-the-definitive/9781492043072/"),
			oreilly2("https://learning.oreilly.com/library/view/software-architecture-the/9781492086888/"),
			oreilly2("https://learning.oreilly.com/library/view/effortless-cloud-native-app/9781801077118/"),
			oreilly2("https://learning.oreilly.com/library/view/hacking-kubernetes/9781492081722/"),
			oreilly2("https://learning.oreilly.com/library/view/five-lines-of/9781617298318/"),
			oreilly2("https://learning.oreilly.com/library/view/learning-test-driven-development/9781098106461/"),
			oreilly2("https://learning.oreilly.com/library/view/building-ci-cd-systems/9781801078214/"),
			oreilly2("https://learning.oreilly.com/library/view/97-things-every/9781098101381/"),
			oreilly2("https://learning.oreilly.com/library/view/modern-devops-practices/9781800562387/"),
			oreilly2("https://learning.oreilly.com/library/view/the-programmers-brain/9781617298677/"),
			oreilly2("https://learning.oreilly.com/library/view/good-code-bad/9781617298936/"),
			oreilly2("https://learning.oreilly.com/library/view/building-microservices-2nd/9781492034018/"),
			oreilly2("https://learning.oreilly.com/library/view/software-telemetry/9781617298141/"),
			oreilly2("https://learning.oreilly.com/library/view/linux-cookbook-2nd/9781492087151/"),
			oreilly2("https://learning.oreilly.com/library/view/building-distributed-applications/9781801074858/"),
			oreilly2("https://learning.oreilly.com/library/view/api-design-patterns/9781617295850/"),
			oreilly2("https://learning.oreilly.com/library/view/devops-adoption-strategies/9781801076326/"),
			oreilly2("https://learning.oreilly.com/library/view/introducing-distributed-application/9781484269985/"),
		}
		page(pg).Render(request.Context(), writer)
	})

	http.ListenAndServe(":3000", nil)

}

func img(img string) string {
	return fmt.Sprintf("/static/images/%s", img)
}

func oreilly2(b templ.SafeURL) RepoLink {
	re := regexp.MustCompile(`/(\d+)/$`)
	matches := re.FindStringSubmatch(string(b))
	bookNumber := fmt.Sprintf("https://learning.oreilly.com/library/cover/%s/250w.png/", matches[1])
	return RepoLink{Repo: b, Img: bookNumber}
}
