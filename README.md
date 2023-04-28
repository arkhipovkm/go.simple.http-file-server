# Go Simple HTTP file server
Simple static content HTTP server. While it is generic, it was with web-app deployments in mind. 

Lightweight (~5MB uncompressed docker image), straightforward and has default handler for unknown paths (that are otherwise handled by the web app's router)

## Usage:
```
Simple HTTP file server. Serves contents of a directory, defaults to index.html if path not found.

Options:
  -a string
        tcp binding addr (default 0.0.0.0)
  -d string
        path to directory to serve statically (default ./)
  -p int
        tcp binding port (default 3000)
```

