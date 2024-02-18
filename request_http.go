package mel

import (
	"io"

	"net/http"
)

type Response struct {
	Res        []byte
	StatusCode int
}

type Headers struct {
	Key   string
	Value string
}

func Get(url string) (Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{
		Res:        b,
		StatusCode: res.StatusCode,
	}, nil
}

func GetWithHeaders(url string, headers []Headers) (Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Response{}, err
	}

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{Res: b, StatusCode: res.StatusCode}, nil
}

func validBody(body []byte) io.Reader {
	var b io.Reader
	if body != nil {
		b = validBody(body)
	} else {
		b = nil
	}

	return b
}

func Post(url string, body []byte) (Response, error) {
	res, err := http.Post(url, "application/json", validBody(body))
	if err != nil {
		return Response{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{
		Res:        b,
		StatusCode: res.StatusCode,
	}, nil
}

func PostWithHeaders(url string, body []byte, headers []Headers) (Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, validBody(body))
	if err != nil {
		return Response{}, err
	}

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{Res: b, StatusCode: res.StatusCode}, nil
}

func PutWithHeaders(url string, body []byte, headers []Headers) (Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, validBody(body))
	if err != nil {
		return Response{}, err
	}

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{Res: b, StatusCode: res.StatusCode}, nil
}

func PatchWithHeaders(url string, body []byte, headers []Headers) (Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, validBody(body))
	if err != nil {
		return Response{}, err
	}

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{Res: b, StatusCode: res.StatusCode}, nil
}

func DeleteWithHeaders(url string, body []byte, headers []Headers) (Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, validBody(body))
	if err != nil {
		return Response{}, err
	}

	for _, v := range headers {
		req.Header.Add(v.Key, v.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{
			StatusCode: res.StatusCode,
		}, nil
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{
			StatusCode: res.StatusCode,
		}, err
	}

	return Response{Res: b, StatusCode: res.StatusCode}, nil
}
