import { useEffect, useState } from 'react'
import './App.css'

function App() {
  const [data, setData] = useState<any[]>([])

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('http://localhost:5000/users?offset=0&limit=1000')
      const jsonData = await response.json()
      setData(jsonData)
    }
    fetchData()

  }, [])

  return (
    <>
      <h1>Robotgram</h1>
      <div>
        <ul>
          {data.map(user => (
            <li key={user.id}>
              {user.username} ({user.gender})
            </li>
          ))}
        </ul>
      </div>
    </>
  )
}

export default App
