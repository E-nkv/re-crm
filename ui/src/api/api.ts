export function FetchLogin(nick: string, pass: string) {
  return fetch("http://localhost:8080/api/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ nick, pass }),
    credentials: "include",
  })
}

export function FetchMe() {
  return fetch("http://localhost:8080/api/me", {
    credentials: "include",
  })
}
