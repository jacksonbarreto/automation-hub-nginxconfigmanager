package autoconfig

const configTemplate = `
location /{{.URLPath}} {
    auth_request /auth-verify;
    proxy_pass http://{{.Host}}:{{.Port}};
}
`
