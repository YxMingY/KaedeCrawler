package modules

import (
	"io"
	"log"
	"net/http"
	"time"
)

var tick = time.Tick(time.Millisecond * 100)

func Download(url string) string {
	resp, err := http.Get(url)
	if nil != err {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return ""
	}
	return string(content)
}

func AntiAntiCrawlerDownload(url string) string {
	<-tick
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if nil != err {
		log.Println(err)
		return ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36 Edg/114.0.1823.82")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cookie", "BIDUPSID=626524610FD7539A57FBB1B2A93801B6; PSTM=1686454104; BAIDUID=626524610FD7539AA7B6AB4CF1E740AD:FG=1; BDUSS=ZxbFVSczFQYUJCSEVnZEdmS1pUQjFxeHlDMkdZNURVSTk0SDlnU0lBQS1XY2RrRVFBQUFBJCQAAAAAAAAAAAEAAADaW187WXhNaW5nWQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD7Mn2Q-zJ9kSj; BDUSS_BFESS=ZxbFVSczFQYUJCSEVnZEdmS1pUQjFxeHlDMkdZNURVSTk0SDlnU0lBQS1XY2RrRVFBQUFBJCQAAAAAAAAAAAEAAADaW187WXhNaW5nWQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD7Mn2Q-zJ9kSj; MCITY=-289%3A231%3A; H_PS_PSSID=36556_38643_39026_39022_38943_38876_38984_39014_39040_38808_38828_38987_39085_26350_39093_39100_39043; BA_HECTOR=8k0la40lag0481ahag0k010a1ibfhs21p; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; ZFY=b3hEFml4Wb9hPbtnTzqFQTlp2zthd3Z4MMqymb:ALV2M:C; BAIDUID_BFESS=626524610FD7539AA7B6AB4CF1E740AD:FG=1; BDRCVFR[feWj1Vr5u3D]=I67x6TjHwwYf0; delPer=0; BDPASSGATE=IlPT2AEptyoA_yiU4SDm3lIN8eDEUsCD34OtViBi3ECGh67BmhGq4qJTSlrLZW3sCU0MdI3Jm1ldjijsQmFuirMenAkAe6NufHi85t727aLUOt-8-rce0rXoGFgNsA8PbRhL-3MEF3V4VFYKbRT9hNgAeOqsxRZIecrR5EDHiMfs2keRBWGFzYKIPoxkPG4fPNu5cPrlnygdPk_cWe8oTi_FgS1iVp1L7qaKiOMmPOD5rkoXGur_QhwlIYvPFXR8_Bjb12q770ys0yU5q-2YSkUtdEiV5sj9IUMMCsDeot2DNv0fJB7AUlPhAKcmjbPbLQdWKQ3zmtsGPTYyynZlJp-j0aLUOmvqNE95RNGGiBjZCXwVqlOMJezY97YzRv4Q4BRYZwNdSikErVaAreCmpSr2HwPgi_lwT0Ug_Sjn9Xgwfn63Gm4K6Hqzc7Rvu7O8ZmGX164p0LbnSIhKuuKyJaXaVKr9E6g2sESWh5DveTCHBfj6RMR7kzgCwnV7TXC3ceKbPjDfXPyKeLZfwrXqvszDu9STzTC3tT8PLOKi4fgZ0oUcmqV4CmjJzGy_eXxAO3bM8mvkE5fTnCdMwtyIjyZnlPA1FtZqZujjJPITwcSIkWwV0yYtpaTHDeAQXJjkufxGIvD0UY8Y1kD9X7lRgrtA2V4vIJSCvqTf1z1g3TWO1C8fG0xP55VKg7UWS_fgC5bxaNYMIliBp4AsP2qip7diFOEjEkV-y7zJ3-8KO4aBz4lJqUVu8ron4o8TYujezIQLoagqUyKuewfzOKytgJgVi9OIgQx9A0m-Omvk8FKyIBEq-nm4; PSINO=6; kleck=8d3b8521902c25ad72ad9ae0cc058300; BDSVRBFE=Go; BDSVRTM=12; __bsi=12250684182493753871_00_44_R_R_0_0303_c02f_Y")
	client := &http.Client{}
	resp, err := client.Do(req)
	if nil != err {
		log.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if nil != err {
		log.Println(err)
		return ""
	}
	return string(body)
}
