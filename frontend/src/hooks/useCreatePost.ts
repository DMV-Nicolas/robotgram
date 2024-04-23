import { toast } from 'sonner'
import { useToken } from './useToken'

export function useCreatePost() {
  const { accessToken, refreshAccessToken, updateAccessToken, updateRefreshToken } = useToken()

  const createPost = async ({ description, images }: { description: string, images: string[] }) => {
    const res = await fetch('http://localhost:5000/v1/posts', {
      method: 'POST',
      body: JSON.stringify({ description, images }),
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${accessToken}`
      },
      credentials: 'include'
    })

    if (!res.ok) {
      toast.error('cannot create post')
      const err = await refreshAccessToken()
      if (err instanceof Error) {
        updateAccessToken('')
        updateRefreshToken('')
      }
      return
    }

    toast.success('post created successfully')
  }

  return { createPost }
}
