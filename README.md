# navi-analyze

this project will try to collect navigation infomation


## navigation timing

- https://www.w3.org/TR/navigation-timing/
- https://developer.mozilla.org/zh-TW/docs/Web/API/Navigation_timing_API


## API

`POST` http://localhost:8080/navi

```json
{
	"start": 1579153941545,
	"navigation_timing": {
		"navigationStart": 1,
		"unloadEventStart": 2,
		"unloadEventEnd": 3
        ...
	}
}
```
