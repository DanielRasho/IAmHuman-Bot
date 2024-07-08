const APISettings = {
  Headers: new Headers({
    Accept: 'application/json'
  }),
  baseURL: import.meta.env.VITE_SERVER_URI || 'http://localhost:8080/'
}

export { APISettings }
