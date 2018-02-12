# scramble-service

Scramble service accepts a text in query string and returns a scrambled version of it in JSON.

```
  http://localhost:8000/scramble?text=Hello,%20world!
```

and returns:

```
  {"text":"Hlelo, wrlod!"}
```

See [MRC Cognition and Brain Sciences Unit](https://www.mrc-cbu.cam.ac.uk/people/matt.davis/cmabridge/) for how brain can decode scrambled words in a sentence.