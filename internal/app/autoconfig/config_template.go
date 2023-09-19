package autoconfig

const configTemplate = `
location /{{.Name}} {
    auth_request /auth-verify;
    proxy_pass http://{{.Host}}:{{.Port}};
}
`
