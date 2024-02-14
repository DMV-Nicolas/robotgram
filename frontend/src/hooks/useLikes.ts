import { toast } from 'sonner'
import { useToken } from './useToken'
import { useEffect, useState } from 'react'
import { type LikesCountResponse, type IsLikedResponse } from '../types'

export function useLikes({ targetID }: { targetID: string }) {
  const { accessToken, refreshAccessToken } = useToken()
  const [likes, setLikes] = useState(0)
  const [liked, setLiked] = useState(false)

  const toggleLike = async () => {
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
      toast.error('cannot toggle like')
      refreshAccessToken()
      return
    }
    setLikes(liked ? likes - 1 : likes + 1)
    setLiked(!liked)
  }

  useEffect(() => {
    const fetchCountLikes = async () => {
      const res = await fetch(`http://localhost:5000/v1/likes/${targetID}/count`)

      if (!res.ok) {
        toast.error('cannot count likes')
        return
      }

      const data: LikesCountResponse = await res.json()
      setLikes(data)
    }

    fetchCountLikes()
  }, [targetID])

  useEffect(() => {
    const fetchIsLiked = async () => {
      const res = await fetch(`http://localhost:5000/v1/likes/${targetID}/liked`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          Authorization: `Bearer ${accessToken.current}`
        },
        credentials: 'include'
      })

      if (!res.ok) {
        toast.error('cannot get like data')
        refreshAccessToken()
        return
      }

      const data: IsLikedResponse = await res.json()
      setLiked(data)
    }

    fetchIsLiked()
  }, [])

  return { toggleLike, liked, likes }
}
