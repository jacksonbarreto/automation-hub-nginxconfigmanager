package autoconfig

const configTemplate = `
location /{{.Name}} {
    allow 127.0.0.1;
    deny all;
    auth_request /auth-verify;
    proxy_pass http://{{.Host}}:{{.Port}};
}
`
