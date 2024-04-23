import { useEffect, useState } from 'react'
import { type GetTokenDataResponse } from '../types'
import { toast } from 'sonner'
import { useToken } from './useToken'

export function useUserID() {
  const { accessToken, refreshAccessToken, updateAccessToken, updateRefreshToken } = useToken()
  const [userID, setUserID] = useState('')

  useEffect(() => {
    const fetchGetTokenData = async () => {
      const res = await fetch('http://localhost:5000/v1/token/data', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
          Authorization: `Bearer ${accessToken}`
        },
        credentials: 'include'
      })

      if (!res.ok) {
        toast.error('cannot get like data')
        const err = await refreshAccessToken()
        if (err instanceof Error) {
          updateAccessToken('')
          updateRefreshToken('')
        }
        return
      }

      const data: GetTokenDataResponse = await res.json()
      setUserID(data.user_id)
    }

    fetchGetTokenData()
  }, [accessToken])

  return { userID }
}
