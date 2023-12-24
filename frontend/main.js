const getData = async () => {
    const data = await fetch("http://localhost:5000/users?offset=0&limit=100")
    const users = await data.json()
    let $users = document.getElementById("users")

    for (let i = 0; i < users.length; i++) {
        $users.innerHTML += users[i].username + "<br>"
    }
}
getData()