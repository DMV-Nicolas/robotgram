import { useState, useEffect } from 'react'
import { type UserType, type UserResponse } from '../types'
import { DEFAULT_USER } from '../constants'
import { toast } from 'sonner'

export const useUserByID = ({ userID }: { userID: string }) => {
  const [user, setUser] = useState(DEFAULT_USER)

  useEffect(() => {
    const fetchGetUser = async () => {
      const res = await fetch(`http://localhost:5000/v1/users/${userID}`)
      if (!res.ok) {
        toast.error('cannot get user data')
      }
      const data: UserResponse = await res.json()
      const user: UserType = {
        id: data.id,
        username: data.username,
        fullName: data.full_name,
        email: data.email,
        avatar: data.avatar,
        description: data.description,
        gender: data.gender,
        createdAt: data.created_at
      }
      setUser(user)
    }

    if (userID.length !== 24) {
      return
    }

    fetchGetUser()
  }, [userID])

  return { user }
}
