const http = require('http')
async function fetch(url) {
    return new Promise((resolve, reject) => {
        http.get(url, res => {
            let data = ''
        
            res.on('data', (chunk) => {
                data += chunk
            })
        
            res.on('end', () => {
                resolve(JSON.parse(data))
            })
        }).on('error', err => {
            reject(err)
        })
    })
}
http.createServer(async function(req, res){
    const movies = await fetch('http://movies-api')
    const content = movies
        .map(({ universe, title, year }) => 
            `${universe} (${title}) - Released on ${year}`)
        .reduce((list, str) => `${list}<li>${str}</li>`, '')
    res.writeHead(200, { 'Content-Type': 'text/html' })
    res.write(`<html><body><h1>Movies</h1><ul>${content}</ul></body></html>`)
    res.end()
}).listen(80)