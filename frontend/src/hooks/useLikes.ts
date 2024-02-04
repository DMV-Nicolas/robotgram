import { toast } from 'sonner'
import { useToken } from './useToken'

export function useLikes() {
  const { accessToken } = useToken()

  const toggleLike = async (targetID: string) => {
    const res = await fetch('http://localhost:5000/v1/likes', {
      method: 'POST',
      body: JSON.stringify({ target_id: targetID }),
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${accessToken.current}`
      },
      credentials: 'include'
    })

    if (!res.ok) {
      const data = await res.json()
      console.log(data)
      console.log(accessToken.current)
      toast.error('cannot ToggleLike')
    }
  }

  return { toggleLike }
}
