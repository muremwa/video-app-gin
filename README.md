## SIMPLE CRUD APP IN Go _(GIN)_

Save and retrieve videos using an api

1. `GET http://localhost:8000/videos` -> a list of videos
2. ```http request
    POST http://localhost:8000/videos
    Accept: application/json
    
    {
        "title": "Video Title",
        "description": "video description",
        "url": "https://some-url.com/video"
    }
```

